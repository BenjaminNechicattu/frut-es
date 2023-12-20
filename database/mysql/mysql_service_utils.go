package mysqlservice

import (
	"log"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host             string `json:"host" validate:"required"`
	Port             int    `json:"port" validate:"required,numeric"`
	Database         string `json:"database" validate:"required"`
	User             string `json:"user" validate:"required"`
	Password         string `json:"password" validate:"required"`
	DSN              string `json:"dsn"`
	ConnMaxOpen      int
	ConnMaxIdleTime  int64
	ConnMaxIdleConns int
}

func GetMysqlConfig(config *MysqlConfig) *mysql.Config {
	mysql_config, err := mysql.ParseDSN(config.DSN)
	log.Println(mysql_config, err)
	if err != nil {
		mysql_config = &mysql.Config{User: config.User, Passwd: config.Password, Net: "tcp",
			Addr: config.Host + ":" + strconv.Itoa(config.Port), DBName: config.Database, Timeout: time.Second * 5, ReadTimeout: time.Second * 5,
			WriteTimeout: time.Second * 5, MultiStatements: true, InterpolateParams: true, AllowNativePasswords: true}
	}
	mysql_config.User = config.User
	mysql_config.Passwd = config.Password
	mysql_config.Addr = config.Host + ":" + strconv.Itoa(config.Port)
	mysql_config.DBName = config.Database

	return mysql_config
}
