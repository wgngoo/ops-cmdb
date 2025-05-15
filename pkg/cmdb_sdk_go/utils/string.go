package utils

import (
	"encoding/json"
	"log"
)

// String 将字符串转换为指针
func String(s string) *string {
	return &s
}

// StringValue 获取字符串指针的值
func StringValue(a *string) string {
	if a == nil {
		return ""
	}
	return *a
}

func StringFormat(queryStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	var err error
	if err = json.Unmarshal([]byte(queryStr), &result); err != nil {
		log.Fatalf("解析查询条件失败: %v", err)
		return result, err
	}
	return result, nil
}
