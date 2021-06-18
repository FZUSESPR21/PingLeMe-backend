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

// CreateTeam 创建团队
func CreateTeam(c *gin.Context) {
	var service service.CreateTeamService
	if err := c.ShouldBind(&service); err == nil {
		service.TeamRepositoryInterface = &model.Repo
		res := service.CreateTeam()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// AddTeammate 增加队员
func AddTeammate(c *gin.Context) {
	var service service.TeammateAddService
	if err := c.ShouldBind(&service); err == nil {
		service.TeamRepositoryInterface = &model.Repo
		service.UserRepositoryInterface = &model.Repo
		res := service.AddTeammate()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteTeammate(c *gin.Context) {
	var service service.TeammateAddService
	if err := c.ShouldBind(&service); err == nil {
		service.TeamRepositoryInterface = &model.Repo
		service.UserRepositoryInterface = &model.Repo
		res := service.DeleteTeammate()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func GetTeamList(c *gin.Context) {
	var service service.TeamListService
	if err := c.ShouldBind(&service); err == nil {
		service.TeamRepositoryInterface = &model.Repo
		res := service.GetTeamList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func GetTeamDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.ParamErr("", err))
		c.Abort()
		return
	}
	var service service.TeamDetailService
	service.UserRepositoryInterface = &model.Repo
	service.TeamRepositoryInterface = &model.Repo
	res := service.GetTeamDetail(uint(id))
	c.JSON(http.StatusOK, res)
}
