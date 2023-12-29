package model

// Vendor 定义枚举
type Vendor int

// 枚举值集合
const (
	PrivateIDC Vendor = iota
	Tencent
	AliYun
	HuaWei
)

// Resource 定义资源对象
type Resource struct {
	Id          string            `json:"id"`          // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`      // 厂商
	Region      string            `json:"region"`      // 地域
	CreateAt    int64             `json:"create_at"`   // 创建时间
	ExpireAt    int64             `json:"expire_at"`   // 过期时间
	Type        string            `json:"type"`        // 规格
	Name        string            `json:"name"`        // 名称
	Description string            `json:"description"` // 描述
	Status      string            `json:"status"`      // 服务商中的状态
	Tags        map[string]string `json:"tags"`        // 标签
	UpdateAt    int64             `json:"update_at"`   // 更新时间
	SyncAt      int64             `json:"sync_at"`     // 同步时间
	Account     string            `json:"account"`     // 资源的所属账号
	PublicIP    string            `json:"public_ip"`   // 公网IP
	PrivateIP   string            `json:"private_ip"`  // 内网IP
}

// ResourceDescribe 定义资源描述对象
type ResourceDescribe struct {
	SerialNumber string `json:"serial_number"` // 序列 ID
	ResourceId   string `json:"resource_id"`   // 资源 ID
	CPU          int    `json:"cpu"`           // 核数
	Memory       int    `json:"memory"`        // 内存
	GPUAmount    int    `json:"gpu_amount"`    // GPU数量
	GPUSpec      string `json:"gpu_spec"`      // GPU类型
	OSType       string `json:"os_type"`       // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`       // 操作系统名称
}

// Host 定义主机信息对象
type Host struct {
	*Resource
	*ResourceDescribe
}

func NewHost() *Host {
	return &Host{
		Resource:         &Resource{},
		ResourceDescribe: &ResourceDescribe{},
	}
}

// HostList 定义 Host 列表对象
type HostList struct {
	Items []*Host
}

func NewHostList() *HostList {
	return &HostList{
		Items: []*Host{},
	}
}

func (HostList *HostList) Add(host *Host) {
	HostList.Items = append(HostList.Items, host)
}
