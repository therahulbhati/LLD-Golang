package restaurant

type Repository interface {
	Save(restaurant *Restaurant) error
	FindByID(id string) (*Restaurant, error)
	FindByName(name string) (*Restaurant, error)
	FindByPincode(pincode string) ([]*Restaurant, error)
}
