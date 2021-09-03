package database

import (
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var Client = NewDatabase(utils.GetHomeDir()+"/.sakurabot-manager.db", false)

func NewDatabase(path string, useMySql bool) *gorm.DB {
	var db *gorm.DB
	var err error
	if useMySql {
		db, err = gorm.Open(mysql.Open(os.ExpandEnv("${MYSQL_USER_NAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DBNAME}?charset=utf8mb4&parseTime=True&loc=Local")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	} else {
		db, err = gorm.Open(sqlite.Open(path), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Tag{}, &User{}, &String{}, &Charge{})
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}
	return db
}
