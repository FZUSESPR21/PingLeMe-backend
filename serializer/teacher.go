package serializer

import "PingLeMe-Backend/model"

type Teacher struct {
	ID          string `json:"teacher_id"`
	TeacherName string `json:"teacher_name"`
}

// BuildTeacher 序列化教师
func BuildTeacher(user model.User) Teacher {
	var teacher Teacher
	teacher.ID = user.UID
	teacher.TeacherName = user.UserName
	return teacher
}

func BuildTeacherResponse(user model.User) Response {
	return Response{
		Code: 0,
		Msg:  "Success",
		Data: BuildTeacher(user),
	}
}

// BuildTeacherList 序列化教师列表
func BuildTeacherList(user []model.User) []Teacher {
	var length = len(user)
	var teacher []Teacher
	for i := 0; i < length; i++ {
		teacher = append(teacher, BuildTeacher(user[i]))
	}
	return teacher
}

func BuildTeacherListResponse(user []model.User) Response {
	return Response{
		Code: 0,
		Msg:  "Success",
		Data: BuildTeacherList(user),
	}
}
