# CMDB SDK Go

这是一个用于访问 CMDB 服务的 Go SDK，提供了简单易用的接口来管理配置管理数据库（CMDB）资源。

## 功能特性

- 支持 CMDB 实例查询和管理
- 支持多种资源类型（主机、OSS、负载均衡、CDN等）
- 支持标签管理
- 支持区域管理
- 支持分页查询
- 支持复杂的查询条件构建
- 支持域名可用性检查

## 安装

```bash
go get github.com/ops-cmdb/cmdb_sdk_go
```

## 快速开始

### 1. 创建配置

```go
config := &api.Config{
	AccessKey: api.String("your_access_key"),
	SecretKey: api.String("your_secret_key"),
	Host:      api.String("ops-cmdb.api.xxx.com"),
}
```

### 2. 初始化客户端

```go
client := api.NewCmdbClient(config)
resourceService := api.NewResourceService(client)
```

### 3. 查询资源

```go
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

// 使用 StringFormat 格式化查询条件
query, err := utils.StringFormat(queryStr)
if err != nil {
	log.Fatalf("解析查询条件失败: %v", err)
}

// 获取资源信息
result, err := resourceService.GetCmdbResource(objectId, fields, query)
if err != nil {
	log.Fatalf("获取资源信息失败: %v", err)
}
fmt.Println(string(result.Data.List))
```

### 4. 检查域名可用性

```go
available, err := resourceService.CheckDomainAvailability("apk.tclclouds.com")
if err != nil {
	log.Fatalf("检查域名是否存在失败: %v", err)
}
fmt.Println(available)
```

## 支持的资源类型

- 主机（Host）
- 对象存储（OSS）
- 负载均衡（LB）
- CDN
- DNS记录

## 查询条件构建

SDK 支持多种查询操作符：

- `$eq`: 等于
- `$ne`: 不等于
- `$gt`: 大于
- `$gte`: 大于等于
- `$lt`: 小于
- `$lte`: 小于等于
- `$in`: 在列表中
- `$nin`: 不在列表中
- `$and`: 与操作
- `$or`: 或操作

## 错误处理

SDK 使用标准的 Go 错误处理机制，所有可能失败的操作都会返回 error。建议始终检查返回的错误：

```go
if err := resourceService.GetCmdbResource(objectId, fields, query); err != nil {
	log.Printf("操作失败: %v", err)
	return
}
```

## 开发指南

### 项目结构

```
cmdb_sdk_go/
├── pkg/
│   └── cmdb_sdk_go/
│       ├── api/        # API 实现
│       ├── models/     # 数据模型
│       └── utils/      # 工具函数
├── examples/           # 示例代码
└── docs/              # 文档
```

### 添加新的资源类型

1. 在 `models` 包中定义新的资源类型结构体
2. 在 `api` 包中实现相应的查询方法
3. 在 `ResourceService` 中添加新的服务方法

## 贡献指南

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 许可证

MIT License

## 联系方式

如有问题或建议，请提交 Issue 或 Pull Request。 