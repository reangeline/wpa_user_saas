package entity_test

import (
	"testing"

	"github.com/reangeline/wpa_user_saas/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	name := "John"
	email := "john.doe@example.com"
	phone := "1234567890"

	user, err := entity.NewUser(name, email, phone)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, user.ID)

}
