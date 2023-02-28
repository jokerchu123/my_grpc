package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString(("mysql.host"))
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	fmt.Println(dsn)
	err := Database(dsn)
	if err != nil {
		panic(err)
	}

}
func Database(dsn string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true, //禁用datatime的精度，mysql5.6不兼容
		DontSupportRenameIndex:   true, //重命名索引的时候采用删除并新建的方法，mysql5.7版本之前不支持重命名
		DontSupportRenameColumn:  true,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
	return err
}
