package serializer

import "PingLeMe-Backend/model"

// EvaluationTable 评审表序列化器
type EvaluationTable struct {
	TableName  string                `json:"table_name"`
	TableItems []EvaluationTableItem `json:"table_items"`
}

// EvaluationTableItem 评审表项序列化器
type EvaluationTableItem struct {
	Content         string                `json:"content"`
	Score           int                   `json:"score"`
	Description     string                `json:"description"`
	ChildTableItems []EvaluationTableItem `json:"child_table_items"`
}

// BuildEvaluationTable 序列化评审表
func BuildEvaluationTable(tableModel model.EvaluationTable) EvaluationTable {
	var table EvaluationTable
	table.TableName = tableModel.TableName
	items := BuildTableItems(0, len(tableModel.TableItem)-1, tableModel.TableItem)
	return EvaluationTable{
		TableName:  tableModel.TableName,
		TableItems: items,
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
					Content:         tableItems[i].Content,
					Score:           tableItems[i].Score,
					ChildTableItems: BuildTableItems(b, e, tableItems),
				}}, items...)
				b = -1
				e = -1
			} else {
				items = append([]EvaluationTableItem{{
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
