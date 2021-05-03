package model

import (
	"testing"
)

func TestEvaluationTable(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	//t.Run("GetEvaluationTable", func(t *testing.T) {
	//	tRepo.mock.ExpectQuery("SELECT (.+) FROM `evaluation_tables` WHERE `evaluation_tables`.`id` = \\? AND (.*)").
	//		WithArgs(1).
	//		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "homework_id", "team_id"}).
	//			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1))
	//	tRepo.mock.ExpectQuery("SELECT (.+) FROM `evaluation_table_items` WHERE `evaluation_table_items`.`evaluation_table_id` = \\? AND `evaluation_table_items`.`deleted_at` IS NULL").
	//		WithArgs(1).
	//		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "evaluation_id", "content", "score", "description", "parent_item_id", "sequence"}).
	//			AddRow(1, time.Now(), time.Now(), time.Now(), 1, "item1", 12, "item1desc", 0, 1).
	//			AddRow(2, time.Now(), time.Now(), time.Now(), 2, "item2", 22, "item2desc", 0, 2),
	//		)
	//
	//	table, err := tRepo.repo.GetEvaluationTable(1)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//
	//	assert.Equal(t, "item1", table.TableItems[0].Content)
	//	assert.Equal(t, "item2", table.TableItems[1].Content)
	//
	//	if err := tRepo.mock.ExpectationsWereMet(); err != nil {
	//		t.Error(err)
	//	}
	//})
}
