//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ClassStuList 修改学生班级
func ClassStuList(c *gin.Context) {
	var service service.ClassStuList
	stus, err := service.StuListOfClass()
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
	} else {
		c.JSON(http.StatusOK, stus)
	}
}