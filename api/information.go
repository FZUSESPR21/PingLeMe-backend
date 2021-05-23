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
		res, err1 := service.EditPairInformation()
		if err1 != nil {
			c.JSON(http.StatusOK, ErrorResponse(err1))
		} else {
			if res == 2 {
				c.JSON(http.StatusOK, serializer.ParamErr("对方已和别人结对，修改结对信息失败", nil))
			} else {
				c.JSON(http.StatusOK, serializer.Response{
					Code: 0,
					Msg:  "修改成功",
				})
			}
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
