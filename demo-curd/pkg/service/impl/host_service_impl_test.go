package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-project/demo-curd/pkg/bean/model"
	"golang-project/demo-curd/pkg/service"
	"testing"
)

func TestHostServiceImpl_InsertHost(t *testing.T) {

	var hostService service.HostService = NewHostServiceImpl()

	ctx := context.Background()

	resource := &model.Resource{
		Id:          "1001",
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

	host := &model.Host{
		Resource:         resource,
		ResourceDescribe: resourceDescribe,
	}

	_, err := hostService.InsertHost(ctx, host)
	if err != nil {
		fmt.Println(err)
	}

}

func TestQueryHost(t *testing.T) {

	ctx := context.Background()

	var hostService service.HostService = NewHostServiceImpl()

	hostList, err := hostService.SelectOneHost(ctx, 1001)
	if err != nil {
		fmt.Println(err)
	}

	jsonByte, _ := json.Marshal(hostList)
	jsonData := string(jsonByte)
	fmt.Println(jsonData)
}

func TestName(t *testing.T) {
	fmt.Println("测试...")

	ctx := context.Background()

	resource := &model.Resource{
		Id:          "1001",
		Vendor:      2,
		Region:      "广东广州天河区",
		CreateAt:    1,
		ExpireAt:    2,
		Type:        "1",
		Name:        "test2",
		Description: "test2",
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

	host := &model.Host{
		Resource:         resource,
		ResourceDescribe: resourceDescribe,
	}

	var hostService service.HostService = NewHostServiceImpl()
	_, err := hostService.UpdateHost(ctx, host)
	if err != nil {
		fmt.Println(err)
	}
}
