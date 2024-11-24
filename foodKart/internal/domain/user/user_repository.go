package user

type Repository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindByPhoneNumber(phoneNumber string) (*User, error)
}
