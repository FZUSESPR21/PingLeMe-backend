package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckLoadedPersonalZeroScore 判断成绩是否已预先存零
func CheckLoadedPersonalZeroScore(c *gin.Context) {
	var checkZeroService service.CheckLoadedBlogService
	if err := c.ShouldBind(&checkZeroService); err == nil {
		checkZeroService.BlogScoreRepositoryInterface = &model.Repo
		res := checkZeroService.CheckLoadedPersonalBlog()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// CheckLoadedTeamZeroScore 判断成绩是否已预先存零
func CheckLoadedTeamZeroScore(c *gin.Context) {
	var checkZeroService service.CheckLoadedBlogService
	if err := c.ShouldBind(&checkZeroService); err == nil {
		checkZeroService.BlogScoreRepositoryInterface = &model.Repo
		res := checkZeroService.CheckLoadedTeamBlog()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UpdatePersonalBlogScore 更新(即评分后载入)个人博客成绩(即助教点击保存评分后)
func UpdatePersonalBlogScore(c *gin.Context) {
	var personalBlogScoreService service.PersonalBlogScoreService
	if err := c.ShouldBind(&personalBlogScoreService); err == nil {
		personalBlogScoreService.BlogScoreRepositoryInterface = &model.Repo
		res := personalBlogScoreService.CountPersonalBlogScore()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UpdatePersonalBlogScore 更新(即评分后载入)团队博客成绩(即助教点击保存评分后)
func UpdateTeamBlogScore(c *gin.Context) {
	var teamBlogScoreService service.TeamBlogScoreService
	if err := c.ShouldBind(&teamBlogScoreService); err == nil {
		teamBlogScoreService.BlogScoreRepositoryInterface = &model.Repo
		res := teamBlogScoreService.CountTeamBlogScore()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}