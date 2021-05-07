package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
)

// CreateHomework 创建作业的接口
func CreateHomework(c *gin.Context) {
	var homeworkService service.HomeworkService
	if err := c.ShouldBind(&homeworkService); err == nil {
		homeworkService.HomeworkRepositoryInterface = &model.Repo
		res := homeworkService.CreateHomework()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}