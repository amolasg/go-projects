package model

// User is the representation of user table
type User struct {
	ID           int
	FirstName    string `json:"first_name" validate:"required,gte=10"`
	LastName     string
	EmailID      string
	MobileNumber string `json:"password" validate:"required,gte=10"`
	Password     string
	Errors       map[string]string
}
