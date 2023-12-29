package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = "host=192.168.2.142 port=5432 user=postgres dbname=ssa password=12345)(*&^%RFVwsx"
var Connect *gorm.DB

func init() {
	ConnectionPg()
}

func ConnectionPg() {
	var err error
	Connect, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("数据库连接失败...")
	}
}

func main() {
	fmt.Println(Connect)
}
