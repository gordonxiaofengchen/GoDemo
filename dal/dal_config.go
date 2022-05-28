package dal

import (
	"demo/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
)

var OrmDb *gorm.DB

func init(){
	username := config.GetString("db.username")
	password := config.GetString("db.password")
	databaseName := config.GetString("db.databaseName")
	host := config.GetString("db.host")
	extraEnv := config.GetString("db.extraEnv")
	port := config.GetInt("db.port")

	var gormLogger logger.Interface
	if config.Has("db.gorm.logger") {
		var logLevel logger.LogLevel
		switch l := strings.ToUpper(config.GetString("db.gorm.logger.logLevel")); l {
		case "INFO": logLevel = logger.Info
		case "WARN": logLevel = logger.Warn
		case "ERROR": logLevel = logger.Error
		case "SILENT": logLevel = logger.Silent
		default: log.Fatal("invalid logLevel: ", l)
		}
		gormLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:              config.GetDuration("db.gorm.logger.slowThreshold"),   // Slow SQL threshold
				LogLevel:                   logLevel, // Log level
				IgnoreRecordNotFoundError: config.GetBool("db.gorm.logger.ignoreRecordNotFoundError"),           // Ignore ErrRecordNotFound error for logger
				Colorful:                  config.GetBool("db.gorm.logger.colorful"),          // Disable color
			},
		)
	}else{
		gormLogger = logger.Default
	}


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password,
		host, port, databaseName, extraEnv)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	OrmDb = db
}