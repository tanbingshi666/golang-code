package service

import (
	"context"
	"golang-project/demo-curd/pkg/bean/model"
)

type HostService interface {
	InsertHost(ctx context.Context, host *model.Host) (*model.Host, error)

	SelectOneHost(ctx context.Context, id int) (*model.HostList, error)

	UpdateHost(ctx context.Context, host *model.Host) (*model.Host, error)

	DeleteHost(ctx context.Context, id int) (*model.Host, error)
}
