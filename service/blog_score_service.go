package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// PersonalBlogScoreService 个人博客成绩服务
type PersonalBlogScoreItem struct {
	model.BlogScoreRepositoryInterface
	ScoringItemID int `json:"scoring_item_id"`
	ScorekeeperID int `json:"scorekeeper_id"`
	Grade         int `json:"grade"`
}

type PersonalBlogScoreService struct {
	model.BlogScoreRepositoryInterface
	Items	[]PersonalBlogScoreItem  `json:"items"`
}

// StorePersonalBlogScore 存储个人博客成绩
func (service *PersonalBlogScoreService) StorePersonalBlogScore() serializer.Response {
	items := service.Items
	var result []model.PersonalBlogScore
	for _, item := range items {
		temp := model.PersonalBlogScore{
			ScoringItemID: item.ScoringItemID,
			ScorekeeperID: item.ScorekeeperID,
			Grade: item.Grade,
		}
		result = append(result, temp)
	}
	err := service.SetPersonalBlogScore(result)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
