//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
)

// CreateAssistant 创建助教接口
func CreateAssistant(c *gin.Context) {
	var service service.CreateAssistantService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.CreateAssistant()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// DeleteAssistant 删除助教接口
func DeleteAssistant(c *gin.Context) {
	var service service.DeleteAssistantService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.DeleteAssistant()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

//RemoveAssistant 移除助教接口
func RemoveAssistant(c *gin.Context) {
	var service service.RemoveAssistantService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		service.ClassRepositoryInterface = &model.Repo
		res := service.RemoveAssistant()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
