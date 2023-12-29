package dao

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-project/demo-curd/pkg/bean/model"
	"testing"
)

func TestMysql(t *testing.T) {

	var err error
	mysqlConnect, err := getMySQLConnection()
	if err != nil {
		_ = fmt.Errorf("初始化 MySQL 失败, %s", err)
		panic(err)
		return
	}
	resource := &model.Resource{
		Id:          "1002",
		Vendor:      1,
		Region:      "广东",
		CreateAt:    1,
		ExpireAt:    1,
		Type:        "1",
		Name:        "test",
		Description: "test",
		Status:      "1",
		Tags:        map[string]string{"key": "value"},
		UpdateAt:    1,
		SyncAt:      1,
		Account:     "test",
		PublicIP:    "test",
		PrivateIP:   "test",
	}

	resourceDescribe := &model.ResourceDescribe{
		SerialNumber: "10001",
		ResourceId:   "1001",
		CPU:          1,
		Memory:       1,
		GPUAmount:    1,
		GPUSpec:      "test",
		OSType:       "Linux",
		OSName:       "Linux",
	}

	// 测试插入 MySQL 数据库
	ctx := context.Background()
	tx, err := mysqlConnect.BeginTx(ctx, nil)
	if err != nil {
		_ = fmt.Errorf("开启 MySQL 事务失败, %s", err)
		return
	}

	// 通过 defer 提交事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				_ = fmt.Errorf("回滚 MySQL 事务, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				_ = fmt.Errorf("提交 MySQL 事务失败, %s", err)
			}
		}
	}()

	insertResourcePrepareContext, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(prepareContext *sql.Stmt) {
		err := prepareContext.Close()
		if err != nil {
			_ = fmt.Errorf("关闭 MySQL 预编译 SQL 失败, %s", err)
		}
	}(insertResourcePrepareContext)

	_, err = insertResourcePrepareContext.ExecContext(ctx,
		resource.Id,
		resource.Vendor,
		resource.Region,
		resource.CreateAt,
		resource.ExpireAt,
		resource.Type,
		resource.Name,
		resource.Description,
		resource.Status,
		func() string {
			marshal, _ := json.Marshal(resource.Tags)
			return string(marshal)
		}(),
		resource.UpdateAt,
		resource.SyncAt,
		resource.Account,
		resource.PublicIP,
		resource.PrivateIP,
	)

	insertDescribePrepareContext, err := tx.PrepareContext(ctx, InsertResourceDescribeSQL)
	if err != nil {
		fmt.Println(err)
	}
	_, err = insertDescribePrepareContext.ExecContext(ctx,
		resourceDescribe.SerialNumber,
		resourceDescribe.ResourceId,
		resourceDescribe.CPU,
		resourceDescribe.Memory,
		resourceDescribe.GPUAmount,
		resourceDescribe.GPUSpec,
		resourceDescribe.OSType,
		resourceDescribe.OSName,
	)

	if err != nil {
		_ = fmt.Errorf("%s", err)
	}

}
