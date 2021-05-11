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
func FillInPairInformation(c *gin.Context, stuUID int) {
	var service service.EditPairIndormationService
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(model.User); ok {
			res, err := service.UpdatePairByStu(int(u.ID), stuUID)
			if err != nil {
				c.JSON(http.StatusOK, ErrorResponse(err))
			}
			if res == 2 {
				c.JSON(http.StatusOK, serializer.ParamErr("对方已和别人结对，修改结对信息失败", nil))
			}
		}
	}
	c.JSON(http.StatusOK, serializer.Response{
		Code: 0,
		Msg:  "修改成功",
	})
}
