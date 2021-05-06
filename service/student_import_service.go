//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type StudentImportService struct {
	model.UserRepositoryInterface
	model.ClassRepositoryInterface
}

func (service *StudentImportService) Import(filepath string) serializer.Response {
	file, err1 := excelize.OpenFile(filepath)
	if err1 != nil {
		return serializer.ServerInnerErr("excelize error", err1)
	}

	rows, err2 := file.GetRows("Sheet1")
	if err2 != nil {
		return serializer.ServerInnerErr("excelize error", err2)
	}

	errs := make([]error, 0)
	users := make([]model.User, 0)
	studentClasses := make([]model.StudentClass, 0)
	tmpMap := make(map[string]model.Class)
	for _, row := range rows{
		var class model.Class
		var err error
		if _, ok := tmpMap[row[2]]; !ok {
			class, err = service.GetClassByName(row[2])
			if err {

			}
		}
		user := model.User{
			UID:            row[0],
			Nickname:       row[1],
			Role:           model.RoleStudent,
		}

		users = append(users, user)
	}


}