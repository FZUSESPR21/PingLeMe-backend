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