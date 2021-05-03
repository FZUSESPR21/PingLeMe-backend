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

	t.Run("GetWorkSubmissionByID", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "submitter_id", "homework_id", "submit_status"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, 0)

		tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(rows)

		workSubmission, err := tRepo.repo.GetWorkSubmissionByID(1)

		if err != nil {
			t.Error(err)
		} else {
			var x uint8 = 0
			assert.Equal(t, workSubmission.SubmitStatus, x)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
