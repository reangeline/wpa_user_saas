package entity

import (
	"errors"

	pkg_entity "github.com/reangeline/wpa_user_saas/pkg/entity"
)

type User struct {
	ID    string `json:"id" dynamodbav:"id"`
	Name  string `json:"name" dynamodbav:"name"`
	Email string `json:"email" dynamodbav:"email"`
	Phone string `json:"phone" dynamodbav:"phone"`
}

func NewUser(name, email, phone string) (*User, error) {
	user := &User{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	user.AddId()

	err := user.IsValid()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) AddId() {
	u.ID = pkg_entity.NewID().String()
}

func (u *User) IsValid() error {

	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.Phone == "" {
		return errors.New("phone is required")
	}

	// phone_numberRegex := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	// if !phone_numberRegex.MatchString(u.PhoneNumber) {
	// 	return errors.New("phone_number is not a valid phone_number number")

	// }

	return nil
}
