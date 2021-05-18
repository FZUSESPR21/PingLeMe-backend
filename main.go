//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package main

import (
	"PingLeMe-Backend/conf"
	"PingLeMe-Backend/router"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	//s := service.EvaluationTableService{
	//	EvaluationTableRepositoryInterface: &model.Repo,
	//	TableName:                          "1231",
	//	HomeworkID:                         1,
	//	TeamID:                             1,
	//	TableItems: []service.EvaluationTableItem{
	//		{
	//			Content: "Table Col 1 Row 1",
	//			Score:   10.0,
	//			ChildrenItems: []service.EvaluationTableItem{
	//				{
	//					Content:       "Table Col 1 Row 2",
	//					Score:         5.0,
	//					ChildrenItems: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//s.CreateEvaluationTable()

	// 装载路由
	r := router.NewRouter()
	_ = r.Run(":3000")
}
