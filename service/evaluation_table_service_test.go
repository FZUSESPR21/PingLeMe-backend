package service

import (
	"PingLeMe-Backend/model"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetChildrenItems(t *testing.T) {
	evaluationTableService := EvaluationTableService{
		TableName: "bala bala",
		TableItems: []EvaluationTableItem{
			{
				Content: "1",
				ChildrenItems: []EvaluationTableItem{
					{
						Content: "5",
						ChildrenItems: []EvaluationTableItem{
							{
								Content: "10",
								ChildrenItems: []EvaluationTableItem{
									{
										Content:       "14",
										ChildrenItems: nil,
									},
									{
										Content:       "15",
										ChildrenItems: nil,
									},
								},
							},
							{
								Content:       "11",
								ChildrenItems: nil,
							},
						},
					},
					{
						Content:       "6",
						ChildrenItems: nil,
					},
				},
			},
			{
				Content: "2",
				ChildrenItems: []EvaluationTableItem{
					{
						Content: "7",
						ChildrenItems: []EvaluationTableItem{
							{
								Content:       "12",
								ChildrenItems: nil,
							},
							{
								Content:       "13",
								ChildrenItems: nil,
							},
						},
					},
					{
						Content:       "8",
						ChildrenItems: nil,
					},
					{
						Content:       "9",
						ChildrenItems: nil,
					},
				},
			},
		},
	}

	except := []model.EvaluationTableItem{
		{
			Content: "1",
			Level:   1,
		},
		{
			Content: "5",
			Level:   2,
		},
		{
			Content: "10",
			Level:   3,
		},
		{
			Content: "14",
			Level:   4,
		},
		{
			Content: "15",
			Level:   4,
		},
		{
			Content: "11",
			Level:   3,
		},
		{
			Content: "6",
			Level:   2,
		},
		{
			Content: "2",
			Level:   1,
		},
		{
			Content: "7",
			Level:   2,
		},
		{
			Content: "12",
			Level:   3,
		},
		{
			Content: "13",
			Level:   3,
		},
		{
			Content: "8",
			Level:   2,
		},
		{
			Content: "9",
			Level:   2,
		},
	}

	t.Run("test 1", func(t *testing.T) {
		items := GetChildrenItems(evaluationTableService.TableItems, 1)

		assert.Equal(t, except, items)
	})
}
