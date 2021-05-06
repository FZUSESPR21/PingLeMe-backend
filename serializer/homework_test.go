package serializer

import (
	"PingLeMe-Backend/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHomework(t *testing.T) {
	homeworkModel := model.Homework{
		ClassID: 1,
		Type: 1,
		Title: "pia pia pia",
		Content: "content pia pia",
		StartTime:  time.Now(),
		EndTime:  time.Now(),
		ScoringItems: []model.ScoringItem{
			{
				Description: "9",
				Level:   2,
			},
			{
				Description: "8",
				Level:   2,
			},
			{
				Description: "13",
				Level:   3,
			},
			{
				Description: "12",
				Level:   3,
			},
			{
				Description: "7",
				Level:   2,
			},
			{
				Description: "2",
				Level:   1,
			},
			{
				Description: "6",
				Level:   2,
			},
			{
				Description: "11",
				Level:   3,
			},
			{
				Description: "15",
				Level:   4,
			},
			{
				Description: "14",
				Level:   4,
			},
			{
				Description: "10",
				Level:   3,
			},
			{
				Description: "5",
				Level:   2,
			},
			{
				Description: "1",
				Level:   1,
			},
		},
	}

	homework := Homework{
		Type: 1,
		Title: "pia pia pia",
		Content: "content pia pia",
		StartTime:  time.Now(),
		EndTime:  time.Now(),
		ScoringItems: []ScoringItem{
			{
				Description: "1",
				ChildScoringItems: []ScoringItem{
					{
						Description: "5",
						ChildScoringItems: []ScoringItem{
							{
								Description: "10",
								ChildScoringItems: []ScoringItem{
									{
										Description:         "14",
										ChildScoringItems: nil,
									},
									{
										Description:         "15",
										ChildScoringItems: nil,
									},
								},
							},
							{
								Description:         "11",
								ChildScoringItems: nil,
							},
						},
					},
					{
						Description:         "6",
						ChildScoringItems: nil,
					},
				},
			},
			{
				Description: "2",
				ChildScoringItems: []ScoringItem{
					{
						Description: "7",
						ChildScoringItems: []ScoringItem{
							{
								Description:         "12",
								ChildScoringItems: nil,
							},
							{
								Description:         "13",
								ChildScoringItems: nil,
							},
						},
					},
					{
						Description:         "8",
						ChildScoringItems: nil,
					},
					{
						Description:         "9",
						ChildScoringItems: nil,
					},
				},
			},
		},
	}

	t.Run("BuildHomework", func(t *testing.T) {
		result := BuildHomework(homeworkModel)

		assert.Equal(t, homework, result)
	})
}
