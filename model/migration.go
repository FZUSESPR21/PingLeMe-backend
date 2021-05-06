//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"PingLeMe-Backend/util"
	"go.uber.org/zap"
)

//执行数据迁移

func migration() {
	// 自动迁移模式
	err := Repo.DB.AutoMigrate(&User{}, &Class{}, &EvaluationItemScore{}, &EvaluationTable{}, &EvaluationTableItem{}, &Homework{},
		&ScoringItem{}, &WorkSubmission{}, &Pair{}, &Performance{}, &PersonalBlogScore{}, &Role{}, &Permission{},
		&Team{}, &TeamBlogScore{})

	if err != nil {
		util.Log().Error("auto migrate error:", zap.Error(err))
	}
}
