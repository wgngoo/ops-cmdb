package main

import (
	"fmt"
	"log"

	"github.com/wgngoo/ops-cmdb/pkg/cmdb_sdk_go/api"
	"github.com/wgngoo/ops-cmdb/pkg/cmdb_sdk_go/utils"
)

func main() {
	// 创建配置
	host := "ops-cmdb.api.xxx.com"
	accessKey := "xxx"
	secretKey := "xxx"

	config := &api.Config{
		AccessKey: api.String(accessKey),
		SecretKey: api.String(secretKey),
		Host:      api.String(host),
	}

	// 创建客户端
	client := api.NewCmdbClient(config)

	// 创建资源服务
	resourceService := api.NewResourceService(client)

	// 设置查询参数
	objectId := "DNSRECORD"
	fields := []string{"name"}

	// 构建查询条件
	queryStr := `{
		"$and": [{
			"$or": [{
				"name": {
					"$eq": "apk.tclclouds.com"
				}
			}]
		}]
	}`

	query, err := utils.StringFormat(queryStr)
	if err != nil {
		log.Fatalf("解析查询条件失败: %v", err)
	}

	// 获取主机信息
	result, err := resourceService.GetCmdbResource(objectId, fields, query)
	if err != nil {
		log.Fatalf("获取主机信息失败: %v", err)
	}
	fmt.Println(string(result.Data.List))

	available, err := resourceService.CheckDomainAvailability("apk.tclclouds.com")
	if err != nil {
		log.Fatalf("检查域名是否存在失败: %v", err)
	}
	fmt.Println(available)
}
