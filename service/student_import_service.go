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

type ErrorRecord struct {
	RowCnt       int      `json:"row_cnt"`
	RowUID       string   `json:"row_uid"`
	ErrRowUID    bool     `json:"err_row_uid"`
	RowName      string   `json:"row_name"`
	ErrRowName   bool     `json:"err_row_name"`
	RowClass     string   `json:"row_class"`
	ErrRowClass  bool     `json:"err_row_class"`
	RowPasswd    string   `json:"row_passwd"`
	ErrRowPasswd bool     `json:"err_row_passwd"`
	ErrMsg       []string `json:"err_msg"`
}

// Import 导入学生
// Excel 格式
// 学号  姓名  班级  密码
func (service *StudentImportService) Import(filepath string) serializer.Response {
	file, err1 := excelize.OpenFile(filepath)
	if err1 != nil {
		return serializer.ServerInnerErr("excelize error", err1)
	}

	rows, err2 := file.GetRows("Sheet1")
	if err2 != nil {
		return serializer.ServerInnerErr("excelize error", err2)
	}

	tmpMap := make(map[string]model.Class)
	errMsgs := make(map[int]ErrorRecord, 0)

	for index, row := range rows {
		var class model.Class
		if _, ok := tmpMap[row[2]]; !ok {
			c, err := service.GetClassByName(row[2])
			if err != nil {
				if e, ok := errMsgs[index+1]; ok {
					if !e.ErrRowClass {
						e.ErrRowClass = true
						e.ErrMsg = append(e.ErrMsg, err.Error())
					}
				} else {
					e := ErrorRecord{
						RowCnt:       index + 1,
						RowUID:       row[0],
						ErrRowUID:    false,
						RowName:      row[1],
						ErrRowName:   false,
						RowClass:     row[2],
						ErrRowClass:  true,
						RowPasswd:    row[3],
						ErrRowPasswd: false,
						ErrMsg:       make([]string, 0),
					}
					e.ErrMsg = append(e.ErrMsg, err.Error())
					errMsgs[index+1] = e
					continue
				}
			} else {
				tmpMap[row[2]] = c
				class = c
			}
		} else {
			class, _ = tmpMap[row[2]]
		}

		user := model.User{
			UID:      row[0],
			Nickname: row[1],
			Role:     model.RoleStudent,
		}
		if err := user.SetPassword(row[3]); err != nil {
			if e, ok := errMsgs[index+1]; ok {
				if !e.ErrRowPasswd {
					e.ErrRowPasswd = true
					e.ErrMsg = append(e.ErrMsg, err.Error())
				}
			} else {
				e := ErrorRecord{
					RowCnt:       index + 1,
					RowUID:       row[0],
					ErrRowUID:    false,
					RowName:      row[1],
					ErrRowName:   false,
					RowClass:     row[2],
					ErrRowClass:  false,
					RowPasswd:    row[3],
					ErrRowPasswd: true,
					ErrMsg:       make([]string, 0),
				}
				e.ErrMsg = append(e.ErrMsg, err.Error())
				errMsgs[index+1] = e
			}
		}
		if err := service.SetUser(user); err != nil {
			if e, ok := errMsgs[index+1]; ok {
				e.ErrMsg = append(e.ErrMsg, err.Error())
			} else {
				e := ErrorRecord{
					RowCnt:       index + 1,
					RowUID:       row[0],
					ErrRowUID:    false,
					RowName:      row[1],
					ErrRowName:   false,
					RowClass:     row[2],
					ErrRowClass:  false,
					RowPasswd:    row[3],
					ErrRowPasswd: false,
					ErrMsg:       make([]string, 0),
				}
				e.ErrMsg = append(e.ErrMsg, err.Error())
				errMsgs[index+1] = e
			}
		} else {
			if err := service.AddStudent(class, user); err != nil {
				if e, ok := errMsgs[index+1]; ok {
					e.ErrMsg = append(e.ErrMsg, err.Error())
				} else {
					e := ErrorRecord{
						RowCnt:       index + 1,
						RowUID:       row[0],
						ErrRowUID:    false,
						RowName:      row[1],
						ErrRowName:   false,
						RowClass:     row[2],
						ErrRowClass:  false,
						RowPasswd:    row[3],
						ErrRowPasswd: false,
						ErrMsg:       make([]string, 0),
					}
					e.ErrMsg = append(e.ErrMsg, err.Error())
					errMsgs[index+1] = e
				}
			}
		}
	}
	return serializer.Response{
		Code: 0,
		Data: errMsgs,
	}
}
