package dao

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-project/demo-curd/pkg/conf"
	"time"
)

var (
	MysqlHost       = conf.GetConfig().Mysql.Host
	MysqlPort       = conf.GetConfig().Mysql.Port
	MysqlUsername   = conf.GetConfig().Mysql.Username
	MysqlPassword   = conf.GetConfig().Mysql.Password
	MysqlDatabase   = conf.GetConfig().Mysql.Database
	MysqlConnection *sql.DB
)

// 初始化 mysql 连接
func init() {

	mysqlConnect, err := getMySQLConnection()
	if err != nil {
		_ = fmt.Errorf("初始化 MySQL 失败, %s", err)
		panic(err)
		return
	}
	MysqlConnection = mysqlConnect
}

func getMySQLConnection() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&multiStatements=true", MysqlUsername, MysqlPassword, MysqlHost, MysqlPort, MysqlDatabase)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Second * time.Duration(60))
	db.SetConnMaxIdleTime(time.Second * time.Duration(60))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}
