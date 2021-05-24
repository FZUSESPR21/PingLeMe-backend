//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ClassStuList 查看班级学生列表
func ClassStuList(c *gin.Context) {
	var service service.ClassStuList
	classID, err1 := strconv.Atoi(c.Param("class_id"))
	if err1 != nil {
		c.JSON(http.StatusOK, ErrorResponse(err1))
	}
	if err := c.ShouldBind(&service); err == nil {
		service.ClassRepositoryInterface = &model.Repo
		stus, err := service.StuListOfClass(classID)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
		} else {
			c.JSON(http.StatusOK, stus)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
