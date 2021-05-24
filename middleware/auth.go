//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package middleware

import (
	"PingLeMe-Backend/auth"
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/serializer"
	"PingLeMe-Backend/util"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.Repo.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// LoginRequired 需要登录
func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusOK, serializer.CheckLogin())
		c.Abort()
	}
}

// PermissionRequired 需要权限
func PermissionRequired(permissionDesc string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if u, ok := user.(*model.User); ok {
				authService := auth.RBACAuth{RBACRepositoryInterface: &model.Repo}
				has, err := authService.CheckUserPermission(*u, permissionDesc)
				if err != nil {
					util.Log().Error("middleware/authService.go/PermissionRequired", zap.Error(err))
					c.JSON(http.StatusOK, serializer.ServerInnerErr("", err))
					c.Abort()
				}
				if has {
					c.Next()
					return
				}
				c.JSON(http.StatusOK, serializer.PermissionDenied())
				c.Abort()
			}
		}

		c.JSON(http.StatusOK, serializer.CheckLogin())
		c.Abort()
	}
}
