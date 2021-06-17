//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetHomeworkList 获取作业列表的接口
func GetHomeworkList(c *gin.Context) {
	class_id, err := strconv.Atoi(c.Query("class_id"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.ParamErr("", err))
		c.Abort()
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.ParamErr("", err))
	}

	var service service.HomeworkListService
	service.HomeworkRepositoryInterface = &model.Repo
	service.ClassRepositoryInterface = &model.Repo
	res := service.GetHomeworkList(uint(class_id), page)
	c.JSON(http.StatusOK, res)
}
