//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"PingLeMe-Backend/util"
	"errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Repository struct {
	DB *gorm.DB
}

var Repo Repository

var AdminDefaultPasswd string

// Database 在中间件中初始化mysql链接
func Database(connString string, logLevel logger.LogLevel) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logLevel,
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})

	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", zap.Error(err))
	}
	//设置连接池
	sqlDB, err1 := db.DB()
	if err1 != nil {
		util.Log().Panic("连接池设置失败", zap.Error(err))
	} else {
		//空闲
		sqlDB.SetMaxIdleConns(50)
		//打开
		sqlDB.SetMaxOpenConns(100)
		//超时
		sqlDB.SetConnMaxLifetime(time.Second * 30)
	}

	Repo.DB = db

	migration()

	firstInit()
}

func firstInit() {
	util.Log().Info("initializing admin account...")
	newAdmin := User{
		UID:      "admin",
		UserName: "admin",
		Role:     RoleAdmin,
	}

	var userResult User
	if result := Repo.DB.Where("role = ?", RoleAdmin).First(&userResult); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if err := newAdmin.SetPassword(AdminDefaultPasswd); err != nil {
				util.Log().Panic("Default admin account init error: SetPassword failed.", zap.Error(err))
			}
			err := Repo.SetUser(newAdmin)
			if err != nil {
				util.Log().Panic("Default admin account init error.", zap.Error(result.Error))
			}
			util.Log().Info("admin default account successfully created.")
		} else {
			util.Log().Panic("Default admin account init error.", zap.Error(result.Error))
		}
	} else {
		newAdmin = userResult
	}

	if roles, err := Repo.GetUserRoles(newAdmin.ID); err != nil {
		util.Log().Panic("Default admin account init error.", zap.Error(err))
	} else {
		has := false
		for _, role := range roles {
			if role.Type == RoleAdmin {
				has = true
				break
			}
		}
		if !has {
			err := Repo.SetUserRole(RoleAdmin, newAdmin)
			if err != nil {
				util.Log().Panic("Default admin account init error: SetUserRole failed.", zap.Error(err))
			}
		}
	}
}
