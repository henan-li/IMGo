package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"log"
	"workplace/IMGo/model"
)

var DbEngine *xorm.Engine

func init() {
	driveName := "mysql"
	connConfig := "root:root@(127.0.0.1:3306)/chat?charset=utf8"

	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driveName, connConfig)

	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.User))

	fmt.Println("init DB...DB is running now")
}
