package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DebugAddUser(c *gin.Context) {
	var service service.DebugUserService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		service.RBACRepositoryInterface = &model.Repo
		res := service.DebugAddUser()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.ParamErr("", err))
	}
}
