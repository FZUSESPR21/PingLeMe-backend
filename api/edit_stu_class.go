//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EditStuClass 修改学生班级
func EditStuClass(c *gin.Context) {
	var service service.EditStudentClassService
	err := service.EditStudentClass()
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code: 0,
			Msg:  "修改学生班级成功！",
		})
	}
}