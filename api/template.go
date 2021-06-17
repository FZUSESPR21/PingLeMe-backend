package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const StudentImportTemplateFilePath = "./template/student_import.xls"

// StudentImportTemplate 下载文件
func StudentImportTemplate(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=学生导入模版.xls")
	c.Header("Content-Type", "application/x-xls")
	c.File(StudentImportTemplateFilePath)
}
