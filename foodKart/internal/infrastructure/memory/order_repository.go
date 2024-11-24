package memory

import (
	"errors"
	"sync"

	"foodkart/internal/domain/order"
)

type OrderRepository struct {
	orders map[string][]*order.Order
	mu     sync.RWMutex
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string][]*order.Order),
	}
}

func (r *OrderRepository) Save(o *order.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[o.UserID] = append(r.orders[o.UserID], o)
	return nil
}

func (r *OrderRepository) FindByUserID(userID string) ([]*order.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if orders, exists := r.orders[userID]; exists {
		return orders, nil
	}
	return nil, errors.New("no orders found for user")
}
