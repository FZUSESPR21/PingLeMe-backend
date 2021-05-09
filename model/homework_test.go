//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"testing"
)

func TestRepository_GetHomeworkByID(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	//t.Run("GetHomeworkByID", func(t *testing.T) {
	//	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "class_id",
	//		"type", "title", "content", "start_time", "end_time"}).
	//		AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, "title_test", "content_test", time.Now(), time.Now())
	//
	//	tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(rows)
	//
	//	homework, err := tRepo.repo.GetHomeworkByID(1)
	//
	//	items := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "homework_id",
	//		"description", "score", "option", "note", "assistant_id", "level", "index"}).
	//		AddRow(1, time.Now(), time.Now(), time.Now(), 1, "description_test",
	//			 20, 5, "note_test", 1, 1, 1)
	//
	//	tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(items)
	//
	//
	//	if err != nil {
	//		t.Error(err)
	//	} else {
	//		assert.Equal(t, homework.Content, "content_test")
	//		assert.Equal(t, homework.Title, "title_test")
	//	}
	//
	//	if err := tRepo.mock.ExpectationsWereMet(); err != nil {
	//		t.Error(err)
	//	}
	//})
}
