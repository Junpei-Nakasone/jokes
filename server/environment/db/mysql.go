package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateDBConnection returns db
func CreateDBConnection() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASSWORD := "root"
	// ローカルのGoからDockerのMySQLに繋ぐ時はコンテナ名ではなくlocalhostで指定する
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "test_database"
	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println("error occured at db connection")
		return nil
	}
	db.LogMode(true)

	return db
}
