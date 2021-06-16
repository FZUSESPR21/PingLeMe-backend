//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package api

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/service"
	"PingLeMe-Backend/util"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const StudentImportFileDst = "./.student_import/"

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.Login(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	err := s.Save()
	if err != nil {
		util.Log().Error("保存Session错误", zap.Error(err))
	}
	c.JSON(http.StatusOK, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

func UserMe(c *gin.Context) {
	var service service.UserInfoService
	service.PairRepositoryInterface = &model.Repo
	service.UserRepositoryInterface = &model.Repo
	res := service.Information(CurrentUser(c).ID)
	c.JSON(http.StatusOK, res)
}

// UserInfo 用户信息接口
func UserInfo(c *gin.Context) {
	var service service.UserInfoService
	service.PairRepositoryInterface = &model.Repo
	service.UserRepositoryInterface = &model.Repo
	userID := c.Param("id")
	user, err := strconv.Atoi(userID)
	if err != nil {
		res := serializer.ParamErr("", err)
		c.JSON(http.StatusOK, res)
	} else if user < 0 {
		res := serializer.ParamErr("用户ID错误", nil)
		c.JSON(http.StatusOK, res)
	} else {
		res := service.Information(uint(user))
		c.JSON(http.StatusOK, res)
	}
}

func GetTeacherList(c *gin.Context) {
	var service service.TeacherListService
	service.UserRepositoryInterface = &model.Repo
	res := service.GetTeacherList()
	c.JSON(http.StatusOK, res)
}

func GetAssistantList(c *gin.Context) {
	var service service.TeacherListService
	service.UserRepositoryInterface = &model.Repo
	res := service.GetAssistantList()
	c.JSON(http.StatusOK, res)
}

func AddTeachers(c *gin.Context) {
	var service service.AddTeacherService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.AddTeacher(false)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func AddAss(c *gin.Context) {
	var service service.AddTeacherService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.AddTeacher(true)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func GetTeachers(c *gin.Context) {
	var service service.TeacherListService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.GetTeacherList()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// StudentImport 文件导入学生（Excel）
func StudentImport(c *gin.Context) {
	file, _ := c.FormFile("file")
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(int(CurrentUser(c).ID)))
	builder.WriteString(time.Now().String())
	builder.WriteString(util.RandStringRunes(5))
	file.Filename = builder.String()

	err := c.SaveUploadedFile(file, StudentImportFileDst)
	if err != nil {
		c.JSON(http.StatusOK, serializer.ServerInnerErr("", err))
	}

	var service service.StudentImportService
	res := service.Import(StudentImportFileDst + file.Filename)
	c.JSON(http.StatusOK, res)
}

func SubmitWorks(c *gin.Context) {
	id, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err1.Error()))
		return
	}
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	files := form.File["files"]

	_base := "./blog"
	exist, err := PathExists(_base)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}

	if !exist {
		// 创建文件夹
		err := os.Mkdir(_base, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}

	var service service.WorkSubmissionService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		service.HomeworkRepositoryInterface = &model.Repo
		service.WorkSubmissionRepositoryInterface = &model.Repo
		service.TeamRepositoryInterface = &model.Repo
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	reg := regexp.MustCompile(`.*-.*-[(](.*)[)]`)

	for _, file := range files {
		filename := "blog/" + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		results := reg.FindStringSubmatch(filename)
		if len(results) <= 1 {
			continue
		}

		//重命名
		newName := results[1] + "_" + GetRandomString(8) + ".pdf"
		if er := os.Rename(filename, newName); er != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("rename file err: %s", er.Error()))
			return
		}
		service.SubmitWork(results[1], newName, uint(id))
	}
	c.JSON(http.StatusOK,
		fmt.Sprintf("Uploaded successfully %d files", len(files)))

}

func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var service service.ChangePasswordService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		res := service.ChangePassword()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// AddStudents 批量添加学生
func AddStudents(c *gin.Context) {
	var service service.AddStudentsService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRepositoryInterface = &model.Repo
		service.ClassRepositoryInterface = &model.Repo
		res := service.AddStudents()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
