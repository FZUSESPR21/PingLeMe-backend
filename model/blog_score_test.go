package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestBlogScore(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("SetPersonalBlogScoreByID", func(t *testing.T) {
		//personal_blog_score := PersonalBlogScore{
		//	ScoringItemID: 	1,
		//	ScorekeeperID: 	1,
		//	Grade:         	70,
		//}
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "scoring_item_id", "scorekeeper_id", "grade"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, 70)

		//tRepo.mock.ExpectExec("UPDATE Personal_Blog_Grade").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))

		personalBlogScore, err := tRepo.repo.SetPersonalBlogScoreByID(1,80)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, personalBlogScore , 1)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("SetTeamBlogScoreByID", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "scoring_item_id", "scorekeeper_id", "grade"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, 70)

		//tRepo.mock.ExpectExec("UPDATE Team_Blog_Grade").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))

		teamBlogScore, err := tRepo.repo.SetTeamBlogScoreByID(1,80)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, teamBlogScore , 1)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}