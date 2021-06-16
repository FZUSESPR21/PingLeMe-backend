//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// WorkSubmissionService 作业导入的服务
type WorkSubmissionService struct {
	model.WorkSubmissionRepositoryInterface
	model.UserRepositoryInterface
	model.TeamRepositoryInterface
	model.HomeworkRepositoryInterface
}

func (service *WorkSubmissionService) SubmitWork(submitterName string, filepath string, homeworkID uint) serializer.Response {
	homework, err := service.GetHomeworkByID(homeworkID)
	if err != nil {
		return serializer.ParamErr("该作业不存在", err)
	}

	if homework.Type == 0 || homework.Type == 1 {
		user, err := service.GetUserByUserName(submitterName)
		if err != nil {
			return serializer.ParamErr("该用户不存在", err)
		}
		_, err = service.CreateWorkSubmission(user.ID, homeworkID, 1, filepath)
		if err != nil {
			return serializer.DBErr("", err)
		}

	} else if homework.Type == 2 {
		team, _, err := service.GetTeamByName(submitterName)
		if err != nil {
			return serializer.ParamErr("该团队不存在", err)
		}
		_, err = service.CreateWorkSubmission(team.ID, homeworkID, 1, filepath)
		if err != nil {
			return serializer.DBErr("", err)
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
