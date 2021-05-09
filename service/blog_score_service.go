package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

type BlogScoreItem struct {
	model.BlogScoreRepositoryInterface
	ScoringItemID int `json:"scoring_item_id"`
	ScorekeeperID int `json:"scorekeeper_id"`
	Grade         int `json:"grade"`
}

// BlogScoreService 博客成绩评分录入服务(一整个一级项下所有的最子项)
type BlogScoreService struct {
	model.BlogScoreRepositoryInterface
	Items []BlogScoreItem `json:"items"`
}

// BlogScoreListService 获取作业列表和对应成绩服务
type BlogScoreListService struct {
	model.BlogScoreRepositoryInterface

}


// StorePersonalBlogScore 存储个人博客成绩
func (service *BlogScoreService) StorePersonalBlogScore() serializer.Response {
	items := service.Items
	var result []model.PersonalBlogScore
	for _, item := range items {
		temp := model.PersonalBlogScore{
			ScoringItemID: item.ScoringItemID,
			ScorekeeperID: item.ScorekeeperID,
			Grade:         item.Grade,
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

// StoreTeamBlogScore 存储个人博客成绩
func (service *BlogScoreService) StoreTeamBlogScore() serializer.Response {
	items := service.Items
	var result []model.TeamBlogScore
	for _, item := range items {
		temp := model.TeamBlogScore{
			ScoringItemID: item.ScoringItemID,
			ScorekeeperID: item.ScorekeeperID,
			Grade:         item.Grade,
		}
		result = append(result, temp)
	}
	err := service.SetTeamBlogScore(result)
	if err != nil {
		return serializer.ParamErr("", err)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}
