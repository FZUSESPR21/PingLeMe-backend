package serializer

import "PingLeMe-Backend/model"

type Student struct {
	ID			uint `json:"student_id"`
	UID         string `json:"student_uid"`
	StudentName string `json:"student_name"`
	PairUID		string `json:"pair_uid"`
	PairName	string `json:"pair_name"`
	TeamID		uint `json:"team_id"`
}

// BuildStudent 序列化学生
func BuildStudent(user model.User, pairUID string, pairName string, teamID uint) Student {
	return Student{
		ID:				user.ID,
		UID:  			user.UID,
		StudentName:	user.UserName,
		PairUID:   		pairUID,
		PairName:  		pairName,
		TeamID:    		teamID,
	}
}

// BuildStudentResponse 序列化学生响应
func BuildStudentResponse(user model.User, pairUID string, pairName string, teamID uint) Response {
	return Response{
		Data: BuildStudent(user, pairUID, pairName, teamID),
	}
}

// BuildStudentList 序列化学生列表
func BuildStudentList(user []model.User) []Student {
	var length = len(user)
	var studentList []Student

	for i := 0; i < length; i++ {
		studentList = append(studentList, BuildStudent(user[i], "",  "",  0))
	}
	return studentList
}

// BuildStudentList 序列化学生列表响应
func BuildStudentListResponse(user []model.User) Response {
	return Response{
		Data: BuildStudentList(user),
	}
}
