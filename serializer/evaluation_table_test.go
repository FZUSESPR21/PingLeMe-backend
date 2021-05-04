package serializer

import (
	"PingLeMe-Backend/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluationTable(t *testing.T) {
	evaluationTableModel := model.EvaluationTable{
		TableName: "bala bala",
		TableItem: []model.EvaluationTableItem{
			{
				Content: "9",
				Level:   2,
			},
			{
				Content: "8",
				Level:   2,
			},
			{
				Content: "13",
				Level:   3,
			},
			{
				Content: "12",
				Level:   3,
			},
			{
				Content: "7",
				Level:   2,
			},
			{
				Content: "2",
				Level:   1,
			},
			{
				Content: "6",
				Level:   2,
			},
			{
				Content: "11",
				Level:   3,
			},
			{
				Content: "15",
				Level:   4,
			},
			{
				Content: "14",
				Level:   4,
			},
			{
				Content: "10",
				Level:   3,
			},
			{
				Content: "5",
				Level:   2,
			},
			{
				Content: "1",
				Level:   1,
			},
		},
	}

	evaluationTable := EvaluationTable{
		TableName: "bala bala",
		TableItems: []EvaluationTableItem{
			{
				Content: "1",
				ChildTableItems: []EvaluationTableItem{
					{
						Content: "5",
						ChildTableItems: []EvaluationTableItem{
							{
								Content: "10",
								ChildTableItems: []EvaluationTableItem{
									{
										Content:         "14",
										ChildTableItems: nil,
									},
									{
										Content:         "15",
										ChildTableItems: nil,
									},
								},
							},
							{
								Content:         "11",
								ChildTableItems: nil,
							},
						},
					},
					{
						Content:         "6",
						ChildTableItems: nil,
					},
				},
			},
			{
				Content: "2",
				ChildTableItems: []EvaluationTableItem{
					{
						Content: "7",
						ChildTableItems: []EvaluationTableItem{
							{
								Content:         "12",
								ChildTableItems: nil,
							},
							{
								Content:         "13",
								ChildTableItems: nil,
							},
						},
					},
					{
						Content:         "8",
						ChildTableItems: nil,
					},
					{
						Content:         "9",
						ChildTableItems: nil,
					},
				},
			},
		},
	}

	t.Run("BuildEvaluationTable", func(t *testing.T) {
		table := BuildEvaluationTable(evaluationTableModel)

		assert.Equal(t, evaluationTable, table)
	})
}
