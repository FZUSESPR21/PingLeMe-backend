//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestRepository_GetClassByID(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetClassByID", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), "software21class")

		tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(rows)

		class, err := tRepo.repo.GetClassByID(1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, class.Name, "software21class")
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
