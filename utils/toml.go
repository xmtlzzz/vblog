package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// toml文件存在多个表，结构体嵌套，单个表可以直接DB
type Config struct {
	Http HttpClient `toml:"http"`
	DBs  DB         `toml:"datasource"` // 指定toml的表名
	Logs Log        `toml:"log"`
}

type HttpClient struct {
	Host string
	Port int
}
type DB struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Debug    bool
}

type Log struct {
	Level   string
	Console LogConsole
	File    LogFile
}

type LogConsole struct {
	Enable  bool
	Nocolor bool
}

type LogFile struct {
	Enable bool
}

func DecodeToml() *Config {
	var conf Config
	_, err := toml.DecodeFile(fmt.Sprintf("%v%v", os.Getenv("workdir"), "\\etc\\application.toml"), &conf)
	if err != nil {
		log.Fatal(err)
	}
	return &conf
}

func NewDBConnecter() *gorm.DB {
	db := DecodeToml().DBs
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", db.Username, db.Password, db.Host, db.Port, db.Database)
	fmt.Println(dsn)
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if db.Debug == true {
		gdb = gdb.Debug()
	}
	return gdb
}
