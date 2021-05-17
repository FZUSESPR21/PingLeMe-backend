package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
)

// CheckLoadedBlogService 判断成绩是否已预先存零的模型
type CheckLoadedBlogService struct {
	model.BlogScoreRepositoryInterface
	HomeworkID  uint  `json:"homework_id"`
	ScorekeeperID uint `json:"scorekeeper_id"`
}


// PersonalBlogScoreService 个人博客成绩模型
type PersonalBlogScoreService struct {
	model.BlogScoreRepositoryInterface
	FirstLevelItemID uint `json:"first_level_item_id"`
	ScorekeeperID    uint `json:"scorekeeper_id"`
	PersonalBlogScoreItems []PersonalBlogScoreItem  `json:"personal_blog_score_items"`
}

// PersonalBlogScoreItem 个人博客成绩项
type PersonalBlogScoreItem struct {
	ScoringItemID    uint `json:"scoring_item_id"`
	Grade            float32  `json:"grade"`
}

// TeamBlogScoreService 团队博客成绩模型
type TeamBlogScoreService struct {
	model.BlogScoreRepositoryInterface
	FirstLevelItemID uint `json:"first_level_item_id"`
	ScorekeeperID    uint `json:"scorekeeper_id"`
	TeamBlogScoreItems []TeamBlogScoreItem `json:"team_blog_score_items"`
}

// TeamBlogScoreItem 团队博客成绩项
type TeamBlogScoreItem struct {
	ScoringItemID    uint `json:"scoring_item_id"`
	Grade            int  `json:"grade"`
}

// CheckLoadedPersonalBlogService
func (service *CheckLoadedBlogService) CheckLoadedPersonalBlog() serializer.Response{
	model.Repo.LoadPersonalZeroScore(service.HomeworkID, service.ScorekeeperID)
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

// CheckLoadedTeamBlogService
func (service *CheckLoadedBlogService) CheckLoadedTeamBlog() serializer.Response{
	model.Repo.LoadTeamZeroScore(service.HomeworkID, service.ScorekeeperID)
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

// CountFirstLevelScoreItem 通过累加第一级其下所有最子项(叶节点)得到第一级评分项自身的得分
func (service *PersonalBlogScoreService) CountFirstLevelScoreItem() ([]model.PersonalBlogScore, error) {
	var firstLevelItem model.PersonalBlogScore
	firstLevelItem.ID = service.FirstLevelItemID
	firstLevelItem.ScorekeeperID = service.ScorekeeperID
	leaves := make([]model.PersonalBlogScore,0)
	for _, scoreItem := range service.PersonalBlogScoreItems{
		var item model.PersonalBlogScore
		item.ScorekeeperID = service.ScorekeeperID
		item.ScoringItemID = scoreItem.ScoringItemID
		item.Grade = scoreItem.Grade
		firstLevelItem.Grade += item.Grade
		leaves = append(leaves, item)
	}
	leaves = append(leaves, firstLevelItem)
	return leaves, nil
}
