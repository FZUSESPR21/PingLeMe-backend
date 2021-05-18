package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

//type AnyTime struct {
//}
//
//// Match satisfies sqlmock.Argument interface
//func (a AnyTime) Match(v driver.Value) bool {
//	_, ok := v.(time.Time)
//	return ok
//}

func TestRBAC(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("SetUserRole", func(t *testing.T) {
		tRepo.mock.ExpectQuery("SELECT (.+) FROM `roles` WHERE type = \\? AND .*").
			WithArgs(1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "type"}).
					AddRow(1, time.Now(), time.Now(), time.Now(), 1),
			)
		tRepo.mock.ExpectExec("INSERT IGNORE INTO user_role").
			WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		users := []User{{
			Model: gorm.Model{
				ID: 1,
			},
			UID:            "1",
			PasswordDigest: "1",
			UserName:       "1",
			Role:           1,
		}}

		err := tRepo.repo.SetUsersRole(1, users)
		if err != nil {
			t.Error(err)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("GetUserRoles", func(t *testing.T) {
		tRepo.mock.ExpectQuery("SELECT (.+) FROM `users` WHERE id = \\? AND .*").
			WithArgs(1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "uid", "password_digest", "nickname", "role"}).
					AddRow(1, time.Now(), time.Now(), time.Now(), "test1", "password", "nickname", 1),
			)

		tRepo.mock.ExpectQuery("SELECT (.+) FROM `user_role`.*").
			WillReturnRows(
				sqlmock.NewRows([]string{"user_id", "role_id"}).
					AddRow(1, 1),
			)

		tRepo.mock.ExpectQuery("SELECT (.+) FROM `roles`.*").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "type"}).
					AddRow(1, time.Now(), time.Now(), time.Now(), 1),
			)

		roles, err := tRepo.repo.GetUserRoles(1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, uint8(1), roles[0].Type)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("GetRolePermission", func(t *testing.T) {
		tRepo.mock.ExpectQuery("SELECT (.+) FROM `roles`.*").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "type"}).
					AddRow(1, time.Now(), time.Now(), time.Now(), 1),
			)

		tRepo.mock.ExpectQuery("SELECT (.+) FROM `role_permission` WHERE `role_permission`.`role_id` = \\?").
			WithArgs(1).
			WillReturnRows(
				sqlmock.NewRows([]string{"role_id", "permission_id"}).
					AddRow(1, 1),
			)

		tRepo.mock.ExpectQuery("SELECT (.+) FROM `permissions` WHERE `permissions`.`id` .*").
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "type"}).
					AddRow(1, time.Now(), time.Now(), time.Now(), 1),
			)

		permissions, err := tRepo.repo.GetRolePermissions(uint8(1))

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, uint(1), permissions[0].Type)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
