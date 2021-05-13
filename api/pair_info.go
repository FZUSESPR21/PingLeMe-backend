//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PairInfo 结对信息
func PairInfo(c *gin.Context) {
	var service service.PairInfoService
	res, err := service.PairInformation()
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
	} else {
		c.JSON(http.StatusOK,  serializer.Response{
			Code: 0,
			Msg:  fmt.Sprint(res),
		})
	}
}