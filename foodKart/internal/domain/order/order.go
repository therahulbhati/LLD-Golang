package order

type Order struct {
	ID             string
	UserID         string
	RestaurantName string
	FoodItemName   string
	Quantity       int
	TotalPrice     float64
}
