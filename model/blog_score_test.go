package model

import (
	"testing"
)

func TestBlogScore(t *testing.T) {
	//var tRepo TestRepository
	//err := tRepo.InitTest()
	//if err != nil {
	//	t.Error(err)
	//}
	//defer tRepo.db.Close()
	//
	//t.Run("GetPersonalBlogScoreByID", func(t *testing.T) {
	//	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "scoring_item_id", "scorekeeper_id", "grade"}).
	//		AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, 80.0)
	//
	//	tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(rows)
	//
	//	personalBlogScore, err := tRepo.repo.GetPersonalBlogScoreByID(1)
	//
	//	if err != nil {
	//		t.Error(err)
	//	} else {
	//		assert.Equal(t, personalBlogScore.Grade, 80.0)
	//	}
	//
	//	if err := tRepo.mock.ExpectationsWereMet(); err != nil {
	//		t.Error(err)
	//	}
	//})
	//
	//t.Run("GetTeamBlogScoreByID", func(t *testing.T) {
	//	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "scoring_item_id", "scorekeeper_id", "grade"}).
	//		AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, 80.0)
	//
	//	tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(rows)
	//
	//	teamBlogScore, err := tRepo.repo.GetTeamBlogScoreByID(1)
	//
	//	if err != nil {
	//		t.Error(err)
	//	} else {
	//		assert.Equal(t, teamBlogScore.Grade, 80.0)
	//	}
	//
	//	if err := tRepo.mock.ExpectationsWereMet(); err != nil {
	//		t.Error(err)
	//	}
	//})
}
