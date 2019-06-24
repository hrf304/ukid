package db

import (
	"github.com/go-xorm/xorm"
	"fmt"
)

var DB *xorm.Engine

func InitDB(){
	var err error = nil
	DB, err = xorm.NewEngine("mysql", "root:aadmin123@tcp(localhost:3306)/test?charset=utf8&loc=Local")
	if err != nil{
		fmt.Println(err)
		return
	}
	DB.ShowSQL(true)
}