package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
)

// SetPersonalBlogScore 保存个人博客得分项的接口
func SetPersonalBlogScore(c *gin.Context) {
	var personalBlogScoreService service.PersonalBlogScoreService
	if err := c.ShouldBind(&personalBlogScoreService); err == nil {
		personalBlogScoreService.BlogScoreRepositoryInterface = &model.Repo
		res := personalBlogScoreService.StorePersonalBlogScore()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}