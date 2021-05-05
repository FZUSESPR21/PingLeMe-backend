//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEvaluationItemGrade(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetEvaluationItemScore", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "score_item_id", "team_id", "uid", "grade"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, 1, 90)

		tRepo.mock.ExpectQuery("SELECT (.+) FROM `evaluation_item_scores` WHERE `evaluation_item_scores`.`id` = \\? AND `evaluation_item_scores`.`deleted_at` IS NULL ORDER BY `evaluation_item_scores`.`id` LIMIT 1").
			WithArgs(1).WillReturnRows(rows)

		itemScore, err := tRepo.repo.GetEvaluationItemScore(1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, itemScore.Grade, 90)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("GetEvaluationItemScores", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "scoring_item_id", "team_id", "uid", "grade"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, "1", 90).
			AddRow(2, time.Now(), time.Now(), time.Now(), 1, 1, "2", 80)

		tRepo.mock.ExpectQuery("SELECT (.+) FROM `evaluation_item_scores`").
			WithArgs(1, 1).WillReturnRows(rows)

		itemScores, err := tRepo.repo.GetEvaluationItemScores(1, 1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, len(itemScores), 2)
			assert.Equal(t, itemScores[0].ScoringItemID, 1)
			assert.Equal(t, itemScores[0].TeamID, 1)
			assert.Equal(t, itemScores[0].UID, "1")
			assert.Equal(t, itemScores[0].Grade, 90)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
