package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestWorkSubmission(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("SetSubmitStatusByID", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "submitter_id", "homework_id", "submit_status"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, 1)

		tRepo.mock.ExpectExec("UPDATE Job_Submission").WithArgs(0, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		x, err := tRepo.repo.SetSubmitStatusByID(1,1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, x , 1)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}