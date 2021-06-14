//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EditStuClass 修改学生班级
func EditStuClass(c *gin.Context) {
	var service service.EditStudentClassService

	if err := c.ShouldBind(&service); err == nil {
		service.ClassRepositoryInterface = &model.Repo
		res := service.EditStudentClass()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))

	}
}
