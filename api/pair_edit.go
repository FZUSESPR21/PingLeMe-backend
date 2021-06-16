//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FillInPairInformation 填写结对信息
func FillInPairInformation(c *gin.Context) {
	var service service.PairEditService
	if err := c.ShouldBind(&service); err == nil {
		service.PairRepositoryInterface = &model.Repo
		service.UserRepositoryInterface = &model.Repo
		res := service.EditPairInformation()
		if res.Data == 2 {
			c.JSON(http.StatusOK, serializer.ParamErr("对方已和别人结对，修改结对信息失败", nil))
		} else if res.Data == 3 {
			c.JSON(http.StatusOK, serializer.ParamErr("保存修改错误，修改结对信息失败", nil))
		} else if res.Data == 4 {
			c.JSON(http.StatusOK, serializer.ParamErr("添加结对失败，修改结对信息失败", nil))
		} else {
			c.JSON(http.StatusOK, serializer.Response{
				Code: 0,
				Msg:  "修改成功",
			})
		}

	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
