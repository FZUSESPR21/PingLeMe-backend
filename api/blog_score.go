package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
)

// SetPersonalBlogScore 保存个人博客得分项的接口
func SetPersonalBlogScore(c *gin.Context) {
	var personalBlogScoreService service.BlogScoreService
	if err := c.ShouldBind(&personalBlogScoreService); err == nil {
		personalBlogScoreService.BlogScoreRepositoryInterface = &model.Repo
		res := personalBlogScoreService.StorePersonalBlogScore()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// SetTeamBlogScore 保存团队博客得分项的接口
func SetTeamBlogScore(c *gin.Context) {
	var teamBlogScoreService service.BlogScoreService
	if err := c.ShouldBind(&teamBlogScoreService); err == nil {
		teamBlogScoreService.BlogScoreRepositoryInterface = &model.Repo
		res := teamBlogScoreService.StoreTeamBlogScore()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
