package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

const (
	KeyPair = "pair"
	KeyTeam = "team"
)

func CreateClass(c *gin.Context) {
	var service service.CreateClassService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, serializer.ParamErr("", err))
	} else {
		service.UserRepositoryInterface = &model.Repo
		service.ClassRepositoryInterface = &model.Repo
		res := service.CreateClass()
		c.JSON(http.StatusOK, res)
	}
}

func PairStatus(c *gin.Context) {
	var service service.GroupStatusService
	service.ClassRepositoryInterface = &model.Repo
	classID := c.Param("id")
	res := service.Status(classID, KeyPair)
	c.JSON(http.StatusOK, res)
}

func TeamStatus(c *gin.Context) {
	var service service.GroupStatusService
	service.ClassRepositoryInterface = &model.Repo
	classID := c.Param("id")
	res := service.Status(classID, KeyTeam)
	c.JSON(http.StatusOK, res)
}

func TogglePair(c *gin.Context) {
	var service service.ToggleGroupService
	classIDStr := c.Query("class_id")
	classID, err1 := strconv.ParseUint(classIDStr, 10, 64)
	if classIDStr == "" {
		c.JSON(http.StatusOK, serializer.ParamErr("missing class id.", nil))
		return
	} else if err1 != nil || classID <= 0 {
		c.JSON(http.StatusOK, serializer.ParamErr("class id param error.", nil))
		return
	}

	timeStr := c.DefaultQuery("duration", "604800")
	t, err2 := strconv.ParseInt(timeStr, 10, 64)
	if timeStr == "" {
		c.JSON(http.StatusOK, serializer.ParamErr("missing deadline t.", nil))
		return
	} else if err2 != nil || t < 0 {
		c.JSON(http.StatusOK, serializer.ParamErr("t param error.", nil))
		return
	}

	res := service.ToggleGroup(uint(classID), time.Duration(t)*time.Second, KeyPair)
	c.JSON(http.StatusOK, res)
	return
}

func ToggleTeam(c *gin.Context) {
	var service service.ToggleGroupService
	classIDStr := c.Query("class_id")
	classID, err1 := strconv.ParseUint(classIDStr, 10, 64)
	if classIDStr == "" {
		c.JSON(http.StatusOK, serializer.ParamErr("missing class id.", nil))
		return
	} else if err1 != nil || classID <= 0 {
		c.JSON(http.StatusOK, serializer.ParamErr("class id param error.", nil))
		return
	}

	timeStr := c.DefaultQuery("duration", "604800")
	t, err2 := strconv.ParseInt(timeStr, 10, 64)
	if timeStr == "" {
		c.JSON(http.StatusOK, serializer.ParamErr("missing deadline t.", nil))
		return
	} else if err2 != nil || t < 0 {
		c.JSON(http.StatusOK, serializer.ParamErr("t param error.", nil))
		return
	}

	res := service.ToggleGroup(uint(classID), time.Duration(t)*time.Second, KeyTeam)
	c.JSON(http.StatusOK, res)
	return
}

func ClassList(c *gin.Context) {
	var service service.ClassListService
	service.ClassRepositoryInterface = &model.Repo
	res := service.GetClassList()
	c.JSON(http.StatusOK, res)
}
