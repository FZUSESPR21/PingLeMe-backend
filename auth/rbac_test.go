//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package auth

import (
	"PingLeMe-Backend/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UserRepoMock struct {
	mock.Mock
}

func (mock *UserRepoMock) GetUserRoles(ID interface{}) ([]model.Role, error) {
	args := mock.Called(ID)
	roles := []model.Role{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Type: 1,
			Desc: "1",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Type: 2,
			Desc: "2",
		},
	}
	return roles, args.Error(1)
}

func (mock *UserRepoMock) GetUserPermissions(ID interface{}) ([]model.Permission, error) {
	args := mock.Called(ID)
	permissions := []model.Permission{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Type: 1,
			Desc: "1",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Type: 2,
			Desc: "2",
		},
	}
	return permissions, args.Error(1)
}

func TestRBACAuth_CheckUserRole(t *testing.T) {
	repo := UserRepoMock{}
	repo.On("GetUserRoles", uint(1)).Return([]model.Role{}, nil)

	auth := RBACAuth{&repo}

	has, err := auth.CheckUserRole(model.User{
		Model: gorm.Model{
			ID: 1,
		},
		UID:            "1",
		PasswordDigest: "1",
		UserName:       "1",
		Role:           1,
	}, "2")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, true, has)

	repo.AssertExpectations(t)
}

func TestRBACAuth_CheckUserPermission(t *testing.T) {
	repo := UserRepoMock{}
	repo.On("GetUserPermissions", uint(1)).Return([]model.Permission{}, nil)

	auth := RBACAuth{&repo}

	has, err := auth.CheckUserPermission(model.User{
		Model: gorm.Model{
			ID: 1,
		},
		UID:            "1",
		PasswordDigest: "1",
		UserName:       "1",
		Role:           1,
	}, "test")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, false, has)

	repo.AssertExpectations(t)
}
