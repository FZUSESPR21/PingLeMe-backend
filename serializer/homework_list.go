//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

import "PingLeMe-Backend/model"

// HomeworkList 作业列表序列化器
type HomeworkList struct {
	TotalPage   int        `json:"total_page"`
	CurrentPage int        `json:"current_page"`
	List        []Homework `json:"list"`
}

// BuildHomeworkList 序列化作业列表
func BuildHomeworkList(homeworks []model.Homework, totalPage int, currentPage int) HomeworkList {
	var homeworkList HomeworkList
	homeworkList.TotalPage = totalPage
	homeworkList.CurrentPage = currentPage
	for _, homework := range homeworks {
		homeworkList.List = append(homeworkList.List, BuildHomework(homework))
	}
	return homeworkList
}
