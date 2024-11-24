package main

import (
	"fmt"

	"foodkart/internal/domain/restaurant"
	"foodkart/internal/domain/user"
	"foodkart/internal/foodkart"
	"foodkart/internal/infrastructure/memory"
)

func main() {
	userRepo := memory.NewUserRepository()
	restaurantRepo := memory.NewRestaurantRepository()
	orderRepo := memory.NewOrderRepository()
	application := foodkart.NewApplication(userRepo, restaurantRepo, orderRepo)

	// Register users
	users := []*user.User{
		{ID: "1", Name: "Pralove", Gender: "M", PhoneNumber: "phoneNumber-1", Pincode: "HSR"},
		{ID: "2", Name: "Nitesh", Gender: "M", PhoneNumber: "phoneNumber-2", Pincode: "BTM"},
		{ID: "3", Name: "Vatsal", Gender: "M", PhoneNumber: "phoneNumber-3", Pincode: "BTM"},
		{ID: "4", Name: "Rahul", Gender: "M", PhoneNumber: "phoneNumber-4", Pincode: "HSR"},
	}

	for _, u := range users {
		err := application.RegisterUser(u)
		if err != nil {
			fmt.Println("Error registering user:", err)
		}
	}

	// Login user to register restaurants
	err := application.LoginUser("phoneNumber-1")
	if err != nil {
		fmt.Println("Error logging in user:", err)
	}

	// Register restaurants
	restaurants := []*restaurant.Restaurant{
		{ID: "1", Name: "Food Court-1", ServiceablePincodes: []string{"BTM", "HSR"}, FoodItemName: "NI Thali", FoodItemPrice: 100, Quantity: 5},
		{ID: "2", Name: "Food Court-2", ServiceablePincodes: []string{"BTM"}, FoodItemName: "Burger", FoodItemPrice: 120, Quantity: 3},
		{ID: "3", Name: "Food Court-3", ServiceablePincodes: []string{"HSR"}, FoodItemName: "SI Thali", FoodItemPrice: 150, Quantity: 1},
	}

	for _, r := range restaurants {
		err := application.RegisterRestaurant(r)
		if err != nil {
			fmt.Println("Error registering restaurant:", err)
		}
	}

	// Place orders for user 1
	err = application.PlaceOrder("Food Court-1", 2)
	if err != nil {
		fmt.Println("Error placing order:", err)
	}

	err = application.PlaceOrder("Food Court-3", 1)
	if err != nil {
		fmt.Println("Error placing order:", err)
	}

	// Get order history for user 1
	orderHistory, err := application.GetOrderHistory()
	if err != nil {
		fmt.Println("Error getting order history:", err)
	} else {
		fmt.Println("Order History for user 1:")
		for _, order := range orderHistory {
			fmt.Printf("Restaurant: %s, Food Item: %s, Quantity: %d, Total Price: %.2f\n",
				order.RestaurantName, order.FoodItemName, order.Quantity, order.TotalPrice)
		}
	}

	// Login another user and place orders
	err = application.LoginUser("phoneNumber-2")
	if err != nil {
		fmt.Println("Error logging in user:", err)
	}

	err = application.PlaceOrder("Food Court-2", 1)
	if err != nil {
		fmt.Println("Error placing order:", err)
	}

	// Get order history for user 2
	orderHistory, err = application.GetOrderHistory()
	if err != nil {
		fmt.Println("Error getting order history:", err)
	} else {
		fmt.Println("Order History for user 2:")
		for _, order := range orderHistory {
			fmt.Printf("Restaurant: %s, Food Item: %s, Quantity: %d, Total Price: %.2f\n",
				order.RestaurantName, order.FoodItemName, order.Quantity, order.TotalPrice)
		}
	}

	// Login another user and place orders
	err = application.LoginUser("phoneNumber-3")
	if err != nil {
		fmt.Println("Error logging in user:", err)
	}

	err = application.PlaceOrder("Food Court-1", 1)
	if err != nil {
		fmt.Println("Error placing order:", err)
	}

	// Get order history for user 3
	orderHistory, err = application.GetOrderHistory()
	if err != nil {
		fmt.Println("Error getting order history:", err)
	} else {
		fmt.Println("Order History for user 3:")
		for _, order := range orderHistory {
			fmt.Printf("Restaurant: %s, Food Item: %s, Quantity: %d, Total Price: %.2f\n",
				order.RestaurantName, order.FoodItemName, order.Quantity, order.TotalPrice)
		}
	}
}
