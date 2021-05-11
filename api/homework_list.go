//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHomeworkList 获取作业列表的接口
func GetHomeworkList(c *gin.Context) {
	var service service.HomeworkListService
	if err := c.ShouldBind(&service); err == nil {
		service.HomeworkRepositoryInterface = &model.Repo
		service.ClassRepositoryInterface = &model.Repo
		res := service.GetHomeworkList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
