//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetEvaluationTable 获取评审表详情
func GetEvaluationTable(c *gin.Context) {
	var service service.EvaluationTableDetailService
	id := c.Param("id")
	if tableID, err := strconv.Atoi(id); err != nil || tableID < 0 {
		c.JSON(http.StatusOK, serializer.ParamErr("URL错误", err))
	} else {
		service.EvaluationTableRepositoryInterface = &model.Repo
		res := service.ViewEvaluationTable(uint(tableID))
		c.JSON(http.StatusOK, res)
	}
}

// FillEvaluationTable 填写评审表
func FillEvaluationTable(c *gin.Context) {
	var service service.EvaluationTableScoreService
	if err := c.ShouldBind(&service); err == nil {
		service.EvaluationTableRepositoryInterface = &model.Repo
		service.FinalTableScoreRepositoryInterface = &model.Repo
		service.EvaluationItemScoreRepositoryInterface = &model.Repo
		res := service.AddEvaluationTableScore()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.ParamErr("", err))
	}
}

// CreateEvaluationTable 创建评审表
func CreateEvaluationTable(c *gin.Context) {
	var service service.EvaluationTableService
	if err := c.ShouldBind(&service); err == nil {
		service.EvaluationTableRepositoryInterface = &model.Repo
		res := service.CreateEvaluationTable()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.ParamErr("", err))
	}
}

// GetEvaluationTableList 获取评审表列表
func GetEvaluationTableList(c *gin.Context) {
	var service service.EvaluationTableListService
	if err := c.ShouldBind(&service); err == nil {
		service.EvaluationTableRepositoryInterface = &model.Repo
		user := CurrentUser(c)
		res := service.GetTableList(*user)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.ParamErr("", err))
	}
}
