//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EditStuClass 修改学生班级
func EditStuClass(c *gin.Context) {
	var service service.EditStudentClassService

	if err := c.ShouldBind(&service); err == nil {
		service.ClassRepositoryInterface = &model.Repo
		err1 := service.EditStudentClass()
		if err1 != nil {
			c.JSON(http.StatusOK, ErrorResponse(err1))
		}else {
			c.JSON(http.StatusOK, serializer.Response{
				Code: 0,
				Msg:  "修改学生班级成功！",
			})
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))

	}
}