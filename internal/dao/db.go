package dao

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func InitDB() {
	if DBInstance == nil {
		mysqlConfig := viper.GetString("mysql")
		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second,   // 慢 SQL 阈值
				LogLevel:                  logger.Silent, // 日志级别
				IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,         // 禁用彩色打印
			},
		)
		db, err := gorm.Open(mysql.Open(mysqlConfig), &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			log.Println("connect to database error", err.Error())
			panic(err)
		}
		// 连接池
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		DBInstance = db
	}
}
