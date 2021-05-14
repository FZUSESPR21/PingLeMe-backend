//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPair(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetPair", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "student1_id", "student2_id"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 2)

		tRepo.mock.ExpectQuery("SELECT (.+) FROM `pairs` WHERE `pairs`.`id` = \\? AND `pairs`.`deleted_at` IS NULL ORDER BY `pairs`.`id` LIMIT 1").
			WithArgs(1).WillReturnRows(rows)

		pair, err := tRepo.repo.GetPair(1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, uint(2), pair.Student2ID)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("GetPair", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "student1_id", "student2_id"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 2)

		tRepo.mock.ExpectQuery("SELECT (.+) FROM `pairs`").
			WithArgs(2, 2).WillReturnRows(rows)

		studentID, err := tRepo.repo.GetPairByStudentID(2)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, uint(1), studentID)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
