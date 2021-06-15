//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ClassAssisList 查看班级助教列表
func ClassAssisList(c *gin.Context) {
	var service service.ClassAssisList
	classID, err1 := strconv.Atoi(c.Param("class_id"))
	if err1 != nil {
		c.JSON(http.StatusOK, ErrorResponse(err1))
	}
	if err := c.ShouldBind(&service); err == nil {
		service.ClassRepositoryInterface = &model.Repo
		res := service.AssistantListOfClass(classID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
