package order

type Repository interface {
	Save(order *Order) error
	FindByUserID(userID string) ([]*Order, error)
}
