package service

import (
	"PingLeMe-Backend/model"
	"github.com/stretchr/testify/assert"
	"time"

	"testing"
)

func TestGetChildScoringItems(t *testing.T) {
	homeworkService := HomeworkService{
		ClassID:   1,
		Type:      1,
		Title:     "pia pia pia",
		Content:   "content pia pia",
		StartTime: time.Now(),
		EndTime:   time.Now(),
		ScoringItems: []ScoringItem{
			{
				Description: "1",
				ChildrenItems: []ScoringItem{
					{
						Description: "5",
						ChildrenItems: []ScoringItem{
							{
								Description: "10",
								ChildrenItems: []ScoringItem{
									{
										Description:   "14",
										ChildrenItems: nil,
									},
									{
										Description:   "15",
										ChildrenItems: nil,
									},
								},
							},
							{
								Description:   "11",
								ChildrenItems: nil,
							},
						},
					},
					{
						Description:   "6",
						ChildrenItems: nil,
					},
				},
			},
			{
				Description: "2",
				ChildrenItems: []ScoringItem{
					{
						Description: "7",
						ChildrenItems: []ScoringItem{
							{
								Description:   "12",
								ChildrenItems: nil,
							},
							{
								Description:   "13",
								ChildrenItems: nil,
							},
						},
					},
					{
						Description:   "8",
						ChildrenItems: nil,
					},
					{
						Description:   "9",
						ChildrenItems: nil,
					},
				},
			},
		},
	}

	except := []model.ScoringItem{
		{
			Description: "1",
			Level:       1,
		},
		{
			Description: "5",
			Level:       2,
		},
		{
			Description: "10",
			Level:       3,
		},
		{
			Description: "14",
			Level:       4,
		},
		{
			Description: "15",
			Level:       4,
		},
		{
			Description: "11",
			Level:       3,
		},
		{
			Description: "6",
			Level:       2,
		},
		{
			Description: "2",
			Level:       1,
		},
		{
			Description: "7",
			Level:       2,
		},
		{
			Description: "12",
			Level:       3,
		},
		{
			Description: "13",
			Level:       3,
		},
		{
			Description: "8",
			Level:       2,
		},
		{
			Description: "9",
			Level:       2,
		},
	}

	t.Run("test 1", func(t *testing.T) {
		items := GetChildScoringItems(homeworkService.ScoringItems, 1)

		assert.Equal(t, except, items)
	})
}
