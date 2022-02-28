package models_test

import (
	"github.com/stretchr/testify/assert"
	"implementacao/entitity"
	. "implementacao/repository/mysql/models"
	"testing"
)

func TestNewUserInsert(t *testing.T) {
	user := NewUserInsert(entitity.User{Name: "Test", Email: "Test@bool.com.br"})

	assert.NotNil(t, user)
	assert.Equal(t, user.Name, "Test")
	assert.Equal(t, user.Email, "Test@bool.com.br")
}
