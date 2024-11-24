package memory

import (
	"errors"
	"fmt"
	"sync"

	"foodkart/internal/domain/restaurant"
)

type RestaurantRepository struct {
	restaurants map[string]*restaurant.Restaurant
	mu          sync.RWMutex
}

func NewRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{
		restaurants: make(map[string]*restaurant.Restaurant),
	}
}

func (r *RestaurantRepository) Save(res *restaurant.Restaurant) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.restaurants[res.ID] = res
	return nil
}

func (r *RestaurantRepository) FindByName(name string) (*restaurant.Restaurant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, res := range r.restaurants {
		if res.Name == name {
			return res, nil
		}
	}
	return nil, errors.New("restaurant not found")
}

func (r *RestaurantRepository) FindByPincode(pincode string) ([]*restaurant.Restaurant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*restaurant.Restaurant
	for _, res := range r.restaurants {
		for _, p := range res.ServiceablePincodes {
			if p == pincode {
				result = append(result, res)
				break
			}
		}
	}
	if len(result) == 0 {
		return nil, errors.New("no restaurants found for pincode")
	}
	return result, nil
}

func (r *RestaurantRepository) FindByID(id string) (*restaurant.Restaurant, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, restaurant := range r.restaurants {
		if restaurant.ID == id {
			return restaurant, nil
		}
	}
	return nil, fmt.Errorf("restaurant not found")
}
