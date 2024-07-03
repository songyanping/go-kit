package utils

import (
	"time"
)

func TimeFormatNow(layout string) (string, error) {
	// 加载上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", err // 如果加载时区失败，则返回错误
	}

	// 获取当前时间，并应用上海时区
	now := time.Now().In(loc)
	// 格式化时间,layout参数可以是"time.Datetime"或"time.RFC3339",或自定义格式
	return now.Format(layout), nil
}
