//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

import "PingLeMe-Backend/model"

// EvaluationTable 评审表序列化器
type EvaluationTable struct {
	TableID    uint                  `json:"table_id"`
	TableName  string                `json:"table_name"`
	TableItems []EvaluationTableItem `json:"table_items"`
}

// EvaluationTableItem 评审表项序列化器
type EvaluationTableItem struct {
	ItemID          uint                  `json:"item_id"`
	Content         string                `json:"content"`
	Score           float32               `json:"score"`
	ChildTableItems []EvaluationTableItem `json:"child_table_items"`
}

// EvaluationTableList 评审表列表序列化器
type EvaluationTableList struct {
	List []EvaluationTableListItem `json:"list"`
}

type EvaluationTableListItem struct {
	TableID   uint   `json:"table_id"`
	TableName string `json:"table_name"`
}

// BuildEvaluationTable 序列化评审表
func BuildEvaluationTable(tableModel model.EvaluationTable) EvaluationTable {
	var table EvaluationTable
	table.TableName = tableModel.TableName
	items := BuildTableItems(0, len(tableModel.TableItem)-1, tableModel.TableItem)
	return EvaluationTable{
		TableID:    tableModel.ID,
		TableName:  tableModel.TableName,
		TableItems: items,
	}
}

// BuildEvaluationTableResponse 序列化评审表
func BuildEvaluationTableResponse(table model.EvaluationTable) Response {
	return Response{
		Code: 0,
		Data: BuildEvaluationTable(table),
	}
}

// BuildEvaluationTableList 序列化评审表列表
func BuildEvaluationTableList(table []model.EvaluationTable) EvaluationTableList {
	var list EvaluationTableList
	list.List = make([]EvaluationTableListItem, len(table))
	for i, t := range table {
		item := EvaluationTableListItem{
			TableID:   t.ID,
			TableName: t.TableName,
		}
		list.List[i] = item
	}
	return list
}

// BuildEvaluationTableListResponse 序列化评审表列表
func BuildEvaluationTableListResponse(table []model.EvaluationTable) Response {
	return Response{
		Code:  0,
		Data:  BuildEvaluationTableList(table),
	}
}

func BuildTableItems(begin, end int, tableItems []model.EvaluationTableItem) []EvaluationTableItem {
	level := tableItems[begin].Level

	b := -1
	e := -1
	heads := make([]EvaluationTableItem, 0)
	items := make([]EvaluationTableItem, 0)
	i := begin
	for i <= end {
		if tableItems[i].Level > level {
			e = i
			if b == -1 {
				b = i
			}
		}

		if tableItems[i].Level == level {
			if b != -1 {
				items = append([]EvaluationTableItem{{
					ItemID:          tableItems[i].ID,
					Content:         tableItems[i].Content,
					Score:           tableItems[i].Score,
					ChildTableItems: BuildTableItems(b, e, tableItems),
				}}, items...)
				b = -1
				e = -1
			} else {
				items = append([]EvaluationTableItem{{
					ItemID:          tableItems[i].ID,
					Content:         tableItems[i].Content,
					Score:           tableItems[i].Score,
					ChildTableItems: nil,
				}}, items...)
			}
		}

		if tableItems[i].Level < level && tableItems[i].Level == 1 {
			childItems := make([]EvaluationTableItem, len(items))
			copy(childItems, items)
			heads = append([]EvaluationTableItem{{
				ItemID:          tableItems[i].ID,
				Content:         tableItems[i].Content,
				Score:           tableItems[i].Score,
				ChildTableItems: childItems,
			}}, heads...)
			items = make([]EvaluationTableItem, 0)
		}
		i = i + 1
	}

	if i < len(tableItems) {
		return items
	} else {
		return heads
	}
}
