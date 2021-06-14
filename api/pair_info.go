//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PairInfo 结对信息
func PairInfo(c *gin.Context) {
	var service service.PairInfoService
	ID, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		c.JSON(http.StatusOK, ErrorResponse(err1))
	}
	if err := c.ShouldBind(&service); err == nil {
		service.PairRepositoryInterface = &model.Repo
		service.UserRepositoryInterface = &model.Repo
		res := service.PairInformation(uint(ID))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
