package serializer

import "PingLeMe-Backend/model"

//助教
type Assis struct {
	AssistantID			uint `json:"assistant_id"`
	AssistantUID         string `json:"assistant_uid"`
	AssistantName string `json:"assistant_name"`
}

//助教和老师
type Assistant struct {
	TeacherID			uint `json:"teacher_id"`
	TeacherUID			string `json:"teacher_uid"`
	TeacherName 		string `json:"teacher_name"`
	AssisList 				[]Assis `json:"assis_list"`
}

// BuildAssis 序列化助教
func BuildAssis(user model.User) Assis {
	return Assis{
		AssistantID: 			user.ID,
		AssistantUID:  			user.UID,
		AssistantName:			user.UserName,
	}
}

// BuildAssistantResponse 序列化助教响应
func BuildAssistantResponse(user model.User) Response {
	return Response{
		Data: BuildAssis(user),
	}
}

// BuildAssistantList 序列化助教列表
func BuildAssistantList(assisList []model.User) []Assis {
	var length = len(assisList)
	var assistantList []Assis

	for i := 0; i < length; i++ {
		assistantList = append(assistantList, BuildAssis(assisList[i]))
	}
	return assistantList
}

// BuildAssistantListResponse 序列化助教列表响应
func BuildAssistantListResponse(assisList []model.User) Response {
	return Response{
		Data: BuildAssistantList(assisList),
	}
}

// BuildAssisAndTea 序列化助教和老师
func BuildAssisAndTea(teacher model.User,assis []model.User) Assistant {
	assisList := make([]Assis, len(assis))
	for i := range assisList {
		assisList[i] = BuildAssis(assis[i])
	}
	return Assistant{
		TeacherID: 			teacher.ID,
		TeacherUID:  		teacher.UID,
		TeacherName:		teacher.UserName,
		AssisList:			assisList,
	}
}

// BuildAssisAndTeaResponse 序列化助教和老师响应
func BuildAssisAndTeaResponse(teacher model.User,assis []model.User) Response {
	return Response{
		Data: BuildAssisAndTea(teacher,assis),
	}
}

