package dto

type UserInput struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserOutput struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
