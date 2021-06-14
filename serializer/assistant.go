package serializer

import "PingLeMe-Backend/model"

type Assistant struct {
	ID			uint `json:"assistant_id"`
	UID         string `json:"assistant_uid"`
	AssistantName string `json:"assistant_name"`
}

// Assistant 序列化助教
func BuildAssistant(user model.User) Assistant {
	return Assistant{
		ID: 			user.ID,
		UID:  			user.UID,
		AssistantName:	user.UserName,
	}
}

// BuildStudentResponse 序列化助教响应
func BuildAssistantResponse(user model.User) Response {
	return Response{
		Data: BuildAssistant(user),
	}
}

// BuildStudentList 序列化助教列表
func BuildAssistantList(user []model.User) []Assistant {
	var length = len(user)
	var assistantList []Assistant

	for i := 0; i < length; i++ {
		assistantList = append(assistantList, BuildAssistant(user[i]))
	}
	return assistantList
}

// BuildStudentList 序列化助教列表响应
func BuildAssistantListResponse(user []model.User) Response {
	return Response{
		Data: BuildAssistantList(user),
	}
}
