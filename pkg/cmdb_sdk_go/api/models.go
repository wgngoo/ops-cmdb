package api

// InstanceApi_ListInstanceRequestParams 查询实例列表请求参数
type InstanceApi_ListInstanceRequestParams struct {
	OnlyRelationView bool   `json:"onlyRelationView,omitempty"`
	Page             int    `json:"page,omitempty"`
	PageSize         int    `json:"pageSize,omitempty"`
	RelationLimit    int    `json:"relationLimit,omitempty"`
	SelectRelations  string `json:"selectRelations,omitempty"`
}

// InstanceApi_ListInstanceResponseBody 查询实例列表响应
type InstanceApi_ListInstanceResponseBody struct {
	List     []map[string]interface{} `json:"list"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
	Total    int                      `json:"total"`
}

// InstanceApi_PostSearchV3RequestBody 搜索实例请求参数
type InstanceApi_PostSearchV3RequestBody struct {
	Fields                  []string                                               `json:"fields,omitempty"`
	IgnoreMissingFieldError bool                                                   `json:"ignoreMissingFieldError,omitempty"`
	Limitations             []InstanceApi_PostSearchV3RequestBody_limitations_item `json:"limitations,omitempty"`
	MetricsFilter           InstanceApi_PostSearchV3RequestBody_metrics_filter     `json:"metricsFilter,omitempty"`
	OnlyMyInstance          bool                                                   `json:"onlyMyInstance,omitempty"`
	Page                    int                                                    `json:"page,omitempty"`
	PageSize                int                                                    `json:"pageSize,omitempty"`
	Permission              []string                                               `json:"permission,omitempty"`
	Query                   map[string]interface{}                                 `json:"query,omitempty"`
	QueryContext            map[string]interface{}                                 `json:"queryContext,omitempty"`
	RelationLimit           int                                                    `json:"relationLimit,omitempty"`
	Sort                    []InstanceApi_PostSearchV3RequestBody_sort_item        `json:"sort,omitempty"`
}

// InstanceApi_PostSearchV3RequestBody_limitations_item 限制项
type InstanceApi_PostSearchV3RequestBody_limitations_item struct {
	Field string                                                           `json:"field,omitempty"`
	Limit int                                                              `json:"limit,omitempty"`
	Sort  []InstanceApi_PostSearchV3RequestBody_limitations_item_sort_item `json:"sort,omitempty"`
}

// InstanceApi_PostSearchV3RequestBody_limitations_item_sort_item 排序项
type InstanceApi_PostSearchV3RequestBody_limitations_item_sort_item struct {
	Key   string `json:"key,omitempty"`
	Order int    `json:"order,omitempty"`
}

// InstanceApi_PostSearchV3RequestBody_metrics_filter 指标过滤器
type InstanceApi_PostSearchV3RequestBody_metrics_filter struct {
	Limitations []InstanceApi_PostSearchV3RequestBody_metrics_filter_limitations_item `json:"limitations,omitempty"`
	TagsFilter  map[string]interface{}                                                `json:"tagsFilter,omitempty"`
	TimeRange   InstanceApi_PostSearchV3RequestBody_metrics_filter_time_range         `json:"timeRange,omitempty"`
}

// InstanceApi_PostSearchV3RequestBody_metrics_filter_limitations_item 指标限制项
type InstanceApi_PostSearchV3RequestBody_metrics_filter_limitations_item struct {
	Limit  int               `json:"limit,omitempty"`
	Metric string            `json:"metric,omitempty"`
	Sort   []ModelSearchSort `json:"sort,omitempty"`
}

// ModelSearchSort 搜索排序
type ModelSearchSort struct {
	Key   string `json:"key,omitempty"`
	Order int    `json:"order,omitempty"`
}

// InstanceApi_PostSearchV3RequestBody_metrics_filter_time_range 时间范围
type InstanceApi_PostSearchV3RequestBody_metrics_filter_time_range struct {
	EndTime   int `json:"endTime,omitempty"`
	StartTime int `json:"startTime,omitempty"`
}

// InstanceApi_PostSearchV3RequestBody_sort_item 排序项
type InstanceApi_PostSearchV3RequestBody_sort_item struct {
	Key   string `json:"key,omitempty"`
	Order int    `json:"order,omitempty"`
}

// InstanceApi_PostSearchV3ResponseBody 搜索实例响应
type InstanceApi_PostSearchV3ResponseBody struct {
	List     []map[string]interface{} `json:"list"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
	Total    int                      `json:"total"`
}
