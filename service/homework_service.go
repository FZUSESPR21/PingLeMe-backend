package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"time"
)

type HomeworkService struct {
	model.HomeworkRepositoryInterface
	ClassID      uint      `json:"class_id" binding:"required"`
	Type         uint8     `json:"type" binding:"required"`
	Title        string    `json:"title" binding:"required"`
	Content      string    `json:"content" binding:"required"`
	StartTime    time.Time `json:"start_time" binding:"required"`
	EndTime      time.Time `json:"end_time" binding:"required"`
	ScoringItems []ScoringItem `json:"scoring_items"`
}

type HomeworkDetailService struct {
	model.HomeworkRepositoryInterface
}

// ScoringItem 评分项模型
type ScoringItem struct {
	Description string `json:"description" binding:"required"`
	Score       int    `json:"score" binding:"required"`
	Option      uint8  `json:"option" binding:"required"`
	Note        string `json:"note" binding:"required"`
	ChildrenItems []ScoringItem `json:"children_items"`
}

// CreateHomework 创建作业
func (service *HomeworkService) CreateHomework() serializer.Response {
	homework := model.Homework{
		ClassID:  service.ClassID,
		Type: service.Type,
		Title:     service.Title,
		Content: service.Content,
		StartTime: service.StartTime,
		EndTime: service.EndTime,
	}

	homework.ScoringItems = GetChildScoringItems(service.ScoringItems, 1)

	err := service.SetHomework(homework)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

// ViewHomework 查看作业
func (service *HomeworkDetailService) ViewHomework(ID uint) serializer.Response {
	homework, err := service.GetHomeworkByID(ID)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.Response{
		Code: 0,
		Data: serializer.BuildHomework(homework),
	}
}

// GetChildrenItems 递归获取子项
func GetChildScoringItems(target []ScoringItem, level int) []model.ScoringItem {
	items := make([]model.ScoringItem, 0)
	for _, item := range target {
		items = append(items, model.ScoringItem{
			Description: item.Description,
			Score:   item.Score,
			Option: item.Option,
			Note: item.Note,
			Level:   level,
		})
		if item.ChildrenItems != nil {
			items = append(items, GetChildScoringItems(item.ChildrenItems, level+1)...)
		}
	}
	return items
}
