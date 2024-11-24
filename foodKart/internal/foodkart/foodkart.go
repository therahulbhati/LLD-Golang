package foodkart

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"foodkart/internal/domain/order"
	"foodkart/internal/domain/restaurant"
	"foodkart/internal/domain/user"
)

type Application struct {
	UserRepo       user.Repository
	RestaurantRepo restaurant.Repository
	OrderRepo      order.Repository
	LoggedInUser   *user.User
}

func NewApplication(userRepo user.Repository, restaurantRepo restaurant.Repository, orderRepo order.Repository) *Application {
	return &Application{
		UserRepo:       userRepo,
		RestaurantRepo: restaurantRepo,
		OrderRepo:      orderRepo,
	}
}

func (app *Application) RegisterUser(u *user.User) error {
	return app.UserRepo.Save(u)
}

func (app *Application) LoginUser(phoneNumber string) error {
	u, err := app.UserRepo.FindByPhoneNumber(phoneNumber)
	if err != nil {
		return err
	}
	app.LoggedInUser = u
	return nil
}

func (app *Application) RegisterRestaurant(r *restaurant.Restaurant) error {
	if app.LoggedInUser == nil {
		return errors.New("no user logged in")
	}
	return app.RestaurantRepo.Save(r)
}

func (app *Application) ShowRestaurant(sortBy string) ([]*restaurant.Restaurant, error) {
	if app.LoggedInUser == nil {
		return nil, errors.New("no user logged in")
	}

	restaurants, err := app.RestaurantRepo.FindByPincode(app.LoggedInUser.Pincode)
	if err != nil {
		return nil, err
	}

	switch sortBy {
	case "price":
		sort.Slice(restaurants, func(i, j int) bool {
			return restaurants[i].FoodItemPrice < restaurants[j].FoodItemPrice
		})
	case "rating":
		sort.Slice(restaurants, func(i, j int) bool {
			return restaurants[i].AverageRating() > restaurants[j].AverageRating()
		})
	default:
		return nil, errors.New("invalid sort option")
	}

	return restaurants, nil
}

func (app *Application) PlaceOrder(restaurantName string, quantity int) error {
	if app.LoggedInUser == nil {
		return errors.New("no user logged in")
	}

	res, err := app.RestaurantRepo.FindByName(restaurantName)
	if err != nil {
		return err
	}

	if res.Quantity < quantity {
		return errors.New("insufficient quantity")
	}

	res.Quantity -= quantity
	order := &order.Order{
		ID:             app.generateOrderID(),
		UserID:         app.LoggedInUser.ID,
		RestaurantName: res.Name,
		FoodItemName:   res.FoodItemName,
		Quantity:       quantity,
		TotalPrice:     float64(quantity) * res.FoodItemPrice,
	}
	return app.OrderRepo.Save(order)
}

func (app *Application) GetOrderHistory() ([]*order.Order, error) {
	if app.LoggedInUser == nil {
		return nil, errors.New("no user logged in")
	}
	return app.OrderRepo.FindByUserID(app.LoggedInUser.ID)
}

func (app *Application) generateOrderID() string {
	return fmt.Sprintf("order-%d", time.Now().UnixNano())
}
