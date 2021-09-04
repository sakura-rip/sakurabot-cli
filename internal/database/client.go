package database

import (
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DefaultClient = NewDatabase(utils.GetHomeDir()+"/.sakurabot-manager.db", false)

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
	err = db.AutoMigrate(&Tag{}, &User{}, &String{}, &Charge{}, &Token{}, &Proxy{})
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}
	return db
}

// Where add conditions
func Where(query interface{}, args ...interface{}) *gorm.DB {
	return DefaultClient.Where(query, args)
}

// Limit specify the number of records to be retrieved
func Limit(c int) *gorm.DB {
	return DefaultClient.Limit(c)
}

// Create insert the value into database
func Create(value interface{}) *gorm.DB {
	return DefaultClient.Create(value)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func Save(value interface{}) *gorm.DB {
	return DefaultClient.Save(value)
}

// Model specify the model you would like to run db operations
func Model(value interface{}) *gorm.DB {
	return DefaultClient.Model(value)
}

// Preload associations with given conditions
func Preload(query string, args ...interface{}) *gorm.DB {
	return DefaultClient.Preload(query, args)
}
