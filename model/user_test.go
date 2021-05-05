//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetUser", func(t *testing.T) {
		tRepo.mock.ExpectQuery("SELECT (.+) FROM `users` WHERE `users`.`id` = \\? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1").
			WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "uid", "password_digest", "nickname", "role"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), "test1", "password", "nickname", 1))

		user, err := tRepo.repo.GetUser(1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, user.UID, "test1")
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("GetUserByUID", func(t *testing.T) {
		tRepo.mock.ExpectQuery("SELECT (.+) FROM `users` WHERE uid = \\? AND (.*)").
			WithArgs("1").WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "uid", "password_digest", "nickname", "role"}).
				AddRow(1, time.Now(), time.Now(), time.Now(), "test1", "password", "nickname", 1))

		user, err := tRepo.repo.GetUserByUID("1")

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, user.UID, "test1")
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
