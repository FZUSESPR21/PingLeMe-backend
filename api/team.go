package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
