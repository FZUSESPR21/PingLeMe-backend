package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
)

func CreateTeam(c *gin.Context) {
	var service service.CreateTeamService
	if err := c.ShouldBind(&service); err == nil {
		service.TeamRepositoryInterface = &model.Repo
		res := service.CreateTeam()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func AddTeammate(c *gin.Context) {
	var service service.TeammateSetService
	if err := c.ShouldBind(&service); err == nil {
		service.TeamRepositoryInterface = &model.Repo
		service.UserRepositoryInterface = &model.Repo
		res := service.AddTeammate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func DeleteTeammate(c *gin.Context) {
	var service service.TeammateSetService
	if err := c.ShouldBind(&service); err == nil {
		service.TeamRepositoryInterface = &model.Repo
		service.UserRepositoryInterface = &model.Repo
		res := service.DeleteTeammate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
