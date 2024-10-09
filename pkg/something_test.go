package example

import (
	"testing"

	"github.com/zeebo/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDatabase(ctrl)

	expectedUser := User{
		ID:   "123",
		Name: "John Doe",
	}

	mockDB.EXPECT().GetUser("123").Return(expectedUser, nil)

	user, err := mockDB.GetUser("123")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}
