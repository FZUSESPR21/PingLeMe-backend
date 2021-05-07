//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

import (
	"PingLeMe-Backend/model"
	"time"
)

// Homework 作业序列化器
type Homework struct {
	Type         uint8         `json:"type"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	StartTime    time.Time     `json:"start_time"`
	EndTime      time.Time     `json:"end_time"`
	ScoringItems []ScoringItem `json:"scoring_items"`
}

// HomeworkList 作业列表序列化器
type HomeworkListItem struct {
	ClassID		 uint		   `json:"class_id"`
	Type         uint8         `json:"type"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	StartTime    time.Time     `json:"start_time"`
	EndTime      time.Time     `json:"end_time"`
}

// ScoringItem 评分项序列化器
type ScoringItem struct {
	Description string `json:"description"`
	Score       int    `json:"score"`
	Option      uint8  `json:"option"`
	Note        string `json:"note"`
	//AssistantID       uint          `json:"assistant_id"`
	ChildScoringItems []ScoringItem `json:"child_scoring_items"`
}

// BuildHomeworkList 序列化作业列表
func BuildHomeworkList(homeworkList []model.Homework) []HomeworkListItem {
	items := make([]HomeworkListItem, len(homeworkList))
	for i := range items {
		items[i].ClassID = homeworkList[i].ClassID
		items[i].Type = homeworkList[i].Type
		items[i].Title = homeworkList[i].Title
		items[i].Content = homeworkList[i].Content
		items[i].StartTime = homeworkList[i].StartTime
		items[i].EndTime = homeworkList[i].EndTime
	}
	return items
}

// BuildHomework 序列化作业
func BuildHomework(homeworkModel model.Homework) Homework {
	var homework Homework
	homework.Type = homeworkModel.Type
	homework.Title = homeworkModel.Title
	homework.Content = homeworkModel.Content
	homework.StartTime = homeworkModel.StartTime
	homework.EndTime = homeworkModel.EndTime
	items := BuildScoringItems(0, len(homeworkModel.ScoringItems)-1, homeworkModel.ScoringItems)
	return Homework{
		Type:         homeworkModel.Type,
		Title:        homeworkModel.Title,
		Content:      homeworkModel.Content,
		StartTime:    homeworkModel.StartTime,
		EndTime:      homeworkModel.EndTime,
		ScoringItems: items,
	}
}

func BuildScoringItems(begin, end int, scoringItems []model.ScoringItem) []ScoringItem {
	level := scoringItems[begin].Level

	b := -1
	e := -1
	heads := make([]ScoringItem, 0)
	items := make([]ScoringItem, 0)
	i := begin
	for i <= end {
		if scoringItems[i].Level > level {
			e = i
			if b == -1 {
				b = i
			}
		}

		if scoringItems[i].Level == level {
			if b != -1 {
				items = append([]ScoringItem{{
					Description:       scoringItems[i].Description,
					Score:             scoringItems[i].Score,
					Option:            scoringItems[i].Option,
					Note:              scoringItems[i].Note,
					ChildScoringItems: BuildScoringItems(b, e, scoringItems),
				}}, items...)
				b = -1
				e = -1
			} else {
				items = append([]ScoringItem{{
					Description:       scoringItems[i].Description,
					Score:             scoringItems[i].Score,
					Option:            scoringItems[i].Option,
					Note:              scoringItems[i].Note,
					ChildScoringItems: nil,
				}}, items...)
			}
		}

		if scoringItems[i].Level < level && scoringItems[i].Level == 1 {
			childItems := make([]ScoringItem, len(items))
			copy(childItems, items)
			heads = append([]ScoringItem{{
				Description:       scoringItems[i].Description,
				Score:             scoringItems[i].Score,
				Option:            scoringItems[i].Option,
				Note:              scoringItems[i].Note,
				ChildScoringItems: childItems,
			}}, heads...)
			items = make([]ScoringItem, 0)
		}
		i = i + 1
	}

	if i < len(scoringItems) {
		return items
	} else {
		return heads
	}
}
