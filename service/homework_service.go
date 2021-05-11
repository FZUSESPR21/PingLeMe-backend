package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"time"
)

const PAGESIZE = 6

type HomeworkService struct {
	model.HomeworkRepositoryInterface
	ClassID      uint          `json:"class_id" binding:"required"`
	Type         uint8         `json:"type" binding:"required"`
	Title        string        `json:"title" binding:"required"`
	Content      string        `json:"content" binding:"required"`
	StartTime    time.Time     `json:"start_time"`
	EndTime      time.Time     `json:"end_time"`
	ScoringItems []ScoringItem `json:"scoring_items"`
}

type HomeworkListService struct {
	model.HomeworkRepositoryInterface
	model.ClassRepositoryInterface
	ClassID uint `json:"class_id" binding:"required"`
	Page    int  `json:"page" binding:"required"`
}

type HomeworkDetailService struct {
	model.HomeworkRepositoryInterface
	HomeworkID uint `json:"homework_id"`
}

// ScoringItem 评分项模型
type ScoringItem struct {
	Description   string        `json:"description" binding:"required"`
	Score         int           `json:"score" binding:"required"`
	Option        uint8         `json:"option" binding:"required"`
	Note          string        `json:"note" binding:"required"`
	ChildrenItems []ScoringItem `json:"children_items"`
}

// CreateHomework 创建作业
func (service *HomeworkService) CreateHomework() serializer.Response {
	homework := model.Homework{
		ClassID:   service.ClassID,
		Type:      service.Type,
		Title:     service.Title,
		Content:   service.Content,
		StartTime: service.StartTime,
		EndTime:   service.EndTime,
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

// ViewHomeworkList 查看作业列表
func (service *HomeworkListService) ViewHomeworkList() serializer.Response {
	sum := service.CountHomework(service.ClassID)
	homework, err := service.GetAllHomeworkByPage(service.ClassID, service.Page, PAGESIZE)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.Response{
		Code: 0,
		Data: serializer.BuildHomeworkList(homework, (sum/PAGESIZE)+1, service.Page),
	}
}

// ViewHomework 查看作业详情
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

// GetChildScoringItems 递归获取子项
func GetChildScoringItems(target []ScoringItem, level int) []model.ScoringItem {
	items := make([]model.ScoringItem, 0)
	for _, item := range target {
		items = append(items, model.ScoringItem{
			Description: item.Description,
			Score:       item.Score,
			Option:      item.Option,
			Note:        item.Note,
			Level:       level,
		})
		if item.ChildrenItems != nil {
			items = append(items, GetChildScoringItems(item.ChildrenItems, level+1)...)
		}
	}
	return items
}

// AssignedToAssistantService 将评分项一级项分配给助教
func (scoringItem *ScoringItem) AssignedToAssistantService(assistantID uint) []model.ScoringItem {
	var aim []ScoringItem
	aim = append(aim, *scoringItem)
	result := GetChildScoringItems(aim, 1)
	for i := range result {
		result[i].AssistantID = assistantID
	}
	return result
}

// GetHomeworkList 获取作业列表函数
func (service *HomeworkListService) GetHomeworkList() serializer.Response {
	class, err := service.GetClassByID(service.ClassID)
	if err != nil {
		return serializer.ParamErr("该班级不存在", err)
	}

	homeworks, err1 := class.GetAllHomework()
	if err1 != nil {
		return serializer.ParamErr("获取作业列表失败", err)
	}

	pages := len(homeworks) / 5
	homeworks = homeworks[(service.Page-1)*5 : (service.Page-1)*5+5]

	return serializer.Response{
		Code: 0,
		Data: serializer.BuildHomeworkList(homeworks, pages, service.Page),
	}
}
