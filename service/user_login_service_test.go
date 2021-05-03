package service

import (
	"PingLeMe-Backend/model"
	"github.com/stretchr/testify/mock"
	"testing"
)

type UserLoginServiceMock struct {
	mock.Mock
}

func (mock *UserLoginServiceMock) GetUserByUID(UID string) (model.User, error) {
	_ = mock.Called(UID)
	return model.User{}, nil
}

func TestUserLoginService(t *testing.T) {

}
