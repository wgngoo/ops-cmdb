package utils

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
