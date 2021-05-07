package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"strconv"
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

// ViewHomework 查看作业的接口
func ViewHomework(c *gin.Context) {
	var homeworkDetailService service.HomeworkDetailService
	if err := c.ShouldBind(&homeworkDetailService); err == nil {
		homeworkDetailService.HomeworkRepositoryInterface = &model.Repo
		intID, _ := strconv.Atoi(c.Param("homework_id"))
		uintID := uint(intID)
		res := homeworkDetailService.ViewHomework(uintID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ViewHomeworkList 查看作业列表的接口
func ViewHomeworkList(c *gin.Context) {
	var homeworkListService service.HomeworkListService
	if err := c.ShouldBind(&homeworkListService); err == nil {
		homeworkListService.HomeworkRepositoryInterface = &model.Repo
		res := homeworkListService.ViewHomeworkList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
