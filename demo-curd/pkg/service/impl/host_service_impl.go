package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-project/demo-curd/pkg/bean/model"
	"golang-project/demo-curd/pkg/dao"
)

type HostServiceImpl struct {
	MysqlConnection *sql.DB
}

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		MysqlConnection: dao.MysqlConnection,
	}
}

func (hostServiceImpl *HostServiceImpl) InsertHost(ctx context.Context, host *model.Host) (*model.Host, error) {

	var err error
	tx, err := hostServiceImpl.MysqlConnection.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

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

	_, err = tx.ExecContext(ctx, dao.InsertResourceSQL,
		host.Resource.Id,
		host.Resource.Vendor,
		host.Resource.Region,
		host.Resource.CreateAt,
		host.Resource.ExpireAt,
		host.Resource.Type,
		host.Resource.Name,
		host.Resource.Description,
		host.Resource.Status,
		func() string {
			marshal, err := json.Marshal(host.Resource.Tags)
			if err != nil {
				return ""
			}
			return string(marshal)
		}(),
		host.Resource.UpdateAt,
		host.Resource.SyncAt,
		host.Resource.Account,
		host.Resource.PublicIP,
		host.Resource.PrivateIP,
	)

	_, err = tx.ExecContext(ctx, dao.InsertResourceDescribeSQL,
		host.ResourceDescribe.SerialNumber,
		host.ResourceDescribe.ResourceId,
		host.ResourceDescribe.CPU,
		host.ResourceDescribe.Memory,
		host.ResourceDescribe.GPUAmount,
		host.ResourceDescribe.GPUSpec,
		host.ResourceDescribe.OSType,
		host.ResourceDescribe.OSName)

	return nil, err
}

func (hostServiceImpl *HostServiceImpl) SelectOneHost(ctx context.Context, id int) (*model.HostList, error) {

	var err error
	rows, err := hostServiceImpl.MysqlConnection.Query(dao.QueryOneHostSQL, id)
	if err != nil {
		return nil, err
	}

	hostList := model.NewHostList()

	var tags string
	for rows.Next() {
		host := model.NewHost()
		err = rows.Scan(
			&host.Resource.Id,
			&host.Resource.Vendor,
			&host.Resource.Region,
			&host.Resource.CreateAt,
			&host.Resource.ExpireAt,
			&host.Resource.Type,
			&host.Resource.Name,
			&host.Resource.Description,
			&host.Resource.Status,
			&tags,
			&host.Resource.UpdateAt,
			&host.Resource.SyncAt,
			&host.Resource.Account,
			&host.Resource.PublicIP,
			&host.Resource.PrivateIP,

			&host.ResourceDescribe.SerialNumber,
			&host.ResourceDescribe.ResourceId,
			&host.ResourceDescribe.CPU,
			&host.ResourceDescribe.Memory,
			&host.ResourceDescribe.GPUAmount,
			&host.ResourceDescribe.GPUSpec,
			&host.ResourceDescribe.OSType,
			&host.ResourceDescribe.OSName,
		)

		tagsMap := make(map[string]string)

		err := json.Unmarshal([]byte(tags), &tagsMap)
		if err != nil {
			fmt.Println("tags 字符串转为 map 失败")
		} else {
			host.Resource.Tags = tagsMap
		}

		hostList.Add(host)
	}

	return hostList, err
}

func (hostServiceImpl *HostServiceImpl) UpdateHost(ctx context.Context, host *model.Host) (*model.Host, error) {

	var err error
	// 一般需要加事务
	_, err = hostServiceImpl.MysqlConnection.ExecContext(ctx, dao.UpdateOneResourceSQL,
		host.Resource.Vendor,
		host.Resource.Region,
		host.Resource.ExpireAt,
		host.Resource.Name,
		host.Resource.Description,
		host.Resource.Id,
	)

	return nil, err
}

func (hostServiceImpl *HostServiceImpl) DeleteHost(ctx context.Context, id int) (*model.Host, error) {

	var err error
	// 一般需要加事务同时删除 resource 和 resource_describe 双表
	_, err = hostServiceImpl.MysqlConnection.Exec(dao.DeleteOneResourceSQL)
	return nil, err
}
