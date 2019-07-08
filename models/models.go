package models

import (
	"fmt"
	"ginapi/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

var (
	db *gorm.DB
)

// init the connection of database
func init() {
	var err error
	db, err = gorm.Open(
		setting.DB_TYPE,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			setting.DB_USER,
			setting.DB_PASSWORD,
			setting.DB_HOST,
			setting.DB_PROT,
			setting.DB_NAME, ))

	if err != nil {
		log.Println(err.Error())
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.TABLE_PREFIX + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
