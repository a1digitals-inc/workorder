package setting


import (
	"gopkg.in/ini.v1"
	"log"
	"strings"
	"time"
)

var (
	cfg        *ini.File
	RUN_MODE   string
	PAGE_SIZE  int
	JWT_SECRET string

	HTTP_PORT     int
	READ_TIMEOUT  time.Duration
	WRITE_TIMEOUT time.Duration

	DB_TYPE      string
	DB_USER      string
	DB_PASSWORD  string
	DB_HOST      string
	DB_PROT      int
	DB_NAME      string
	TABLE_PREFIX string
)

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(err.Error())
	}
	loadApp()
	loadServer()
	loadDataBase()
}

func loadServer() {
	sec, err := cfg.GetSection("server")
	if err != nil {
		log.Fatal(err.Error())
	}
	HTTP_PORT = sec.Key("HTTP_PORT").MustInt(8080)
	READ_TIMEOUT = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WRITE_TIMEOUT = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp() {
	sec, err := cfg.GetSection("app")
	if err != nil {
		log.Fatal(err.Error())
	}
	RUN_MODE = sec.Key("RUN_MODE").MustString("debug")
	PAGE_SIZE = sec.Key("PAGE_SIZE").MustInt(10)
	JWT_SECRET = sec.Key("JWT_SECRET").MustString("#####")
}

func loadDataBase() {
	sec, err := cfg.GetSection("database")
	if err != nil {
		log.Fatal(err.Error())
	}
	DB_TYPE = strings.ToLower(sec.Key("TYPE").MustString("mysql"))
	DB_USER = sec.Key("USER").MustString("")
	DB_PASSWORD = sec.Key("PASSWORD").MustString("")
	DB_HOST = sec.Key("HOST").MustString("")
	DB_PROT = sec.Key("PROT").MustInt(3306)
	// TODO 如果键不存在会怎么样
	DB_NAME = sec.Key("NAME").String()
	TABLE_PREFIX = sec.Key("TABLE_PREFIX").String()
}