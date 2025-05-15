# CMDB SDK Go

这是一个用于访问 CMDB 服务的 Go SDK。

## 安装

```bash
go get github.com/ops-cmdb/cmdb_sdk_go
```

## 使用示例

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/ops-cmdb/cmdb_sdk_go/pkg/cmdb_sdk_go/api"
	"github.com/ops-cmdb/cmdb_sdk_go/pkg/cmdb_sdk_go/models"
	"github.com/ops-cmdb/cmdb_sdk_go/pkg/cmdb_sdk_go/utils"
)

func main() {
	// 创建配置
	host := "ops-cmdb.api.leiniao.com"
	accessKey := "your_access_key"
	secretKey := "your_secret_key"

	config := &models.Config{
		AccessKey: utils.String(accessKey),
		SecretKey: utils.String(secretKey),
		Host:      utils.String(host),
	}

	// 创建客户端
	client := api.NewCmdbClient(config)

	// 创建资源服务
	resourceService := api.NewResourceService(client)

	// 设置查询参数
	objectId := "DNSRECORD"
	fields := []string{"name", "ip", "hostname"}
	
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
	
	// 使用 StringFormat 格式化查询条件
	query, err := utils.StringFormat(queryStr)
	if err != nil {
		log.Fatalf("格式化查询条件失败: %v", err)
	}

	// 获取主机信息
	if err := resourceService.GetHostForCmdb(objectId, fields, query); err != nil {
		log.Fatalf("获取主机信息失败: %v", err)
	}
}
```

## 功能特性

- 支持 CMDB 实例查询
- 支持资源管理
- 支持标签管理
- 支持区域管理

## 文档

详细的 API 文档请参考 [API 文档](docs/api.md)。

## 许可证

MIT 