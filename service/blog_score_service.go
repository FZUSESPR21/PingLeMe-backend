package service

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"fmt"
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
	HomeworkID		 uint `json:"homework_id"`
	AssistantID		 uint `json:"assistant_id"`
	ScorekeeperID    uint `json:"scorekeeper_id"`
	PersonalBlogScoreItems []PersonalBlogScoreItem  `json:"personal_blog_score_items"`
}

// PersonalBlogScoreItem 个人博客成绩项
type PersonalBlogScoreItem struct {
	ScoringItemID    uint `json:"scoring_item_id"`
	Grade            float32  `json:"grade"`
	ChildrenItems	 []PersonalBlogScoreItem
}

// TeamBlogScoreService 团队博客成绩模
type TeamBlogScoreService struct {
	model.BlogScoreRepositoryInterface
	HomeworkID		 uint `json:"homework_id"`
	AssistantID		 uint `json:"assistant_id"`
	ScorekeeperID    uint `json:"scorekeeper_id"`
	TeamBlogScoreItems []TeamBlogScoreItem `json:"team_blog_score_items"`
}

// TeamBlogScoreItem 团队博客成绩项
type TeamBlogScoreItem struct {
	ScoringItemID    uint `json:"scoring_item_id"`
	Grade            float32  `json:"grade"`
	ChildrenItems	 []TeamBlogScoreItem
}

// CheckLoadedPersonalBlogService
func (service *CheckLoadedBlogService) CheckLoadedPersonalBlog() serializer.Response{
	err1, err2, err3 := model.Repo.LoadPersonalZeroScore(service.HomeworkID, service.ScorekeeperID)
	if err1 != nil {
		fmt.Println("ERROR")
	}
	if err2 != nil {
		fmt.Println("ERROR")
	}
	if err3 != nil {
		fmt.Println("ERROR")
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

// CheckLoadedTeamBlogService
func (service *CheckLoadedBlogService) CheckLoadedTeamBlog() serializer.Response{
	err1, err2, err3 := model.Repo.LoadTeamZeroScore(service.HomeworkID, service.ScorekeeperID)
	if err1 != nil {
		fmt.Println("ERROR")
	}
	if err2 != nil {
		fmt.Println("ERROR")
	}
	if err3 != nil {
		fmt.Println("ERROR")
	}
	return serializer.Response{
		Code: 0,
		Msg:  "Success",
	}
}

// CountPersonalBlogScore 将前端传回的某个助教对某人一次作业的评分结果(一级项数组)
func (service *PersonalBlogScoreService) CountPersonalBlogScore() error {
	for _, item := range service.PersonalBlogScoreItems {
		GetPersonalFatherScore(item, service.ScorekeeperID)
	}
	var aim model.AssistantScored
	aim.ScorekeeperID = service.ScorekeeperID
	aim.HomeworkID = service.HomeworkID
	aim.AssistantID = service.AssistantID
	result := model.Repo.AddAssistantScored(aim)
	return result
}

// GetPersonalFatherScore 通过累加一级其下一级所有子项得到本级评分项自身的得分
func GetPersonalFatherScore(fatherItem PersonalBlogScoreItem, scorekeeperID uint) float32 {
	if fatherItem.ChildrenItems != nil {
		fatherItem.Grade = 0
		for _, item := range fatherItem.ChildrenItems {
			fatherItem.Grade += GetPersonalFatherScore(item, scorekeeperID)
		}
	}
	var scoringItem model.ScoringItem
	result1 := model.Repo.DB.Where("id = ?", fatherItem.ScoringItemID).First(&scoringItem)
	if result1 != nil {
		fmt.Println("blog_score_service的GetPersonalFatherScore方法出错(result1)")
	}
	if fatherItem.Grade > float32(scoringItem.Score) {
		fatherItem.Grade = float32(scoringItem.Score)
	}
	result2 := model.Repo.DB.Exec("update personal_blog_score set grade = ? where scorekeeper_id = ? and scoring_item_id = ?",
		fatherItem.Grade, scorekeeperID, fatherItem.ScoringItemID)
	if result2 != nil {
		fmt.Println("blog_score_service的GetPersonalFatherScore方法出错(result2)")
	}
	return fatherItem.Grade
}

// CountTeamBlogScore 将前端传回的某个助教对某团队一次作业的评分结果(一级项数组)
func (service *TeamBlogScoreService) CountTeamBlogScore() error {
	for _, item := range service.TeamBlogScoreItems {
		GetTeamFatherScore(item, service.ScorekeeperID)
	}
	var aim model.AssistantScored
	aim.ScorekeeperID = service.ScorekeeperID
	aim.HomeworkID = service.HomeworkID
	aim.AssistantID = service.AssistantID
	result := model.Repo.AddAssistantScored(aim)
	return result
}

// GetTeamFatherScore 通过累加一级其下一级所有子项得到本级评分项自身的得分
func GetTeamFatherScore(fatherItem TeamBlogScoreItem, scorekeeperID uint) float32 {
	if fatherItem.ChildrenItems != nil {
		fatherItem.Grade = 0
		for _, item := range fatherItem.ChildrenItems {
			fatherItem.Grade += GetTeamFatherScore(item, scorekeeperID)
		}
	}
	var scoringItem model.ScoringItem
	result1 := model.Repo.DB.Where("id = ?", fatherItem.ScoringItemID).First(&scoringItem)
	if result1 != nil {
		fmt.Println("blog_score_service的GetFatherScore方法出错(result1)")
	}
	if fatherItem.Grade > float32(scoringItem.Score) {
		fatherItem.Grade = float32(scoringItem.Score)
	}
	result2 := model.Repo.DB.Exec("update personal_blog_score set grade = ? where scorekeeper_id = ? and scoring_item_id = ?",
		fatherItem.Grade, scorekeeperID, fatherItem.ScoringItemID)
	if result2 != nil {
		fmt.Println("blog_score_service的GetFatherScore方法出错(result2)")
	}
	return fatherItem.Grade
}
