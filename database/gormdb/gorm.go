package gormdb

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"mariadbapp/config"
)

func NewGormConfig(appConfig *config.AppConfig) (*gorm.DB, error) {
	c := &mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", appConfig.Db.Host, appConfig.Db.Port),
		DBName:               appConfig.Db.DBName,
		User:                 appConfig.Db.Username,
		Passwd:               appConfig.Db.Password,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	//dns := `appConfig.Db.Username + ":" +
	//	appConfig.Db.Password + "@tcp" +"(" +
	//	appConfig.Db.Host + ":" +
	//	appConfig.Db.Port + ")/" +
	//	appConfig.Db.DBName + "?" + "parseTime=true"`
	return gorm.Open("mysql", c.FormatDSN())
}
