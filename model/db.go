package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
	"viaBlog/utils"
)

var db *gorm.DB
var err error

func InitDb()  {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm 日志模式: silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表明明
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库连接失败，请检查参数:", err)
		os.Exit(1)
	}

	// 迁移数据表，在没有数据表结构变更的时候，建议注释不执行
	_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDB, _ := db.DB()
	// 设置连接池中的最大闲置连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置数据库最大连接数量
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
