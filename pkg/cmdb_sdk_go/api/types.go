package api

import "encoding/json"

// Config CMDB 配置
type Config struct {
	Host      *string `json:"host"`
	AccessKey *string `json:"accessKey"`
	SecretKey *string `json:"secretKey"`
}

// CmdbClient CMDB 客户端
type CmdbClient struct {
	Config
}

// ResourceService 资源服务
type ResourceService struct {
	cmdbClient *CmdbClient
}

// CmdbTAG CMDB 标签
type CmdbTAG struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

// CmdbRegion CMDB 区域
type CmdbRegion struct {
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	RegionExtId string `json:"regionExtId"`
}

// CmdbOrigin CMDB 源
type CmdbOrigin struct {
	Origin     string `json:"origin"`
	Protocol   string `json:"protocol"`
	ServerName string `json:"serverName"`
}

// HostForCmdbResponse 主机响应
type HostForCmdbResponse struct {
	Hostname          string    `json:"hostname"`
	Ip                string    `json:"ip"`
	PublicIpAddresses []string  `json:"publicIpAddresses"`
	Account           string    `json:"account"`
	CloudRegion       string    `json:"cloudRegion"`
	TAG               []CmdbTAG `json:"tag"`
	InstanceId        string    `json:"instanceId"`
	InstanceName      string    `json:"instanceName"`
	CloudBrand        string    `json:"cloudBrand"`
	CloudId           string    `json:"cloudId"`
}

// OssForCmdbResponse OSS响应
type OssForCmdbResponse struct {
	InstanceId   string             `json:"instanceId"`
	InstanceName string             `json:"name"`
	Account      string             `json:"account"`
	Location     string             `json:"location"`
	AccessUrls   []CmdbOssAccessUrl `json:"accessUrls"`
	TAG          []CmdbTAG          `json:"tag"`
	CloudBrand   string             `json:"cloudBrand"`
}

// CmdbOssAccessUrl OSS访问URL
type CmdbOssAccessUrl struct {
	Primary     bool   `json:"primary"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

// CmdbListResponse CMDB列表响应
type CmdbListResponse struct {
	Data struct {
		List     json.RawMessage `json:"list"`
		Total    int             `json:"total"`
		Page     int             `json:"page"`
		PageSize int             `json:"pageSize"`
	} `json:"data"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// LbForCmdbResponse 负载均衡响应
type LbForCmdbResponse struct {
	InstanceId   string       `json:"instanceId"`
	InstanceName string       `json:"name"`
	Account      string       `json:"account"`
	CloudBrand   string       `json:"cloudBrand"`
	TAG          []CmdbTAG    `json:"tag"`
	Region       []CmdbRegion `json:"region"`
	Ip           string       `json:"ip"`
}

// CdnForCmdbResponse CDN响应
type CdnForCmdbResponse struct {
	Account      string       `json:"account"`
	InstanceName string       `json:"instanceId"`
	Cname        string       `json:"cname"`
	Origins      []CmdbOrigin `json:"origins"`
	TAG          []CmdbTAG    `json:"tag"`
	CloudBrand   string       `json:"cloudBrand"`
}
