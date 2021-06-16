package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type WorkListService struct {
	model.WorkSubmissionRepositoryInterface
}

type workList struct {
	HomeworkID   uint   `json:"homework_id"`
	HomeworkList []work `json:"homework_list"`
}

type work struct {
	StudentID   uint   `json:"student_id"`
	FileName    string `json:"file_name"`
	IsSubmitted bool   `json:"is_submitted"`
	IsReviewed  bool   `json:"is_reviewed"`
}

func (service *WorkListService) GetWorkList(homeworkID int) serializer.Response {
	works, err := service.GetWorkSubmissionsByHomeworkID(uint(homeworkID))
	if err != nil {
		return serializer.DBErr("", err)
	}
	list := workList{
		HomeworkID:   uint(homeworkID),
		HomeworkList: make([]work, 0),
	}
	for _, w := range works {
		list.HomeworkList = append(list.HomeworkList, work{
			StudentID:   w.ID,
			FileName:    w.Filepath,
			IsSubmitted: w.SubmitStatus != 0,
			IsReviewed:  w.IsReviewed,
		})
	}

	return serializer.Response{
		Code: 0,
		Data: list,
		Msg:  "Success",
	}
}
