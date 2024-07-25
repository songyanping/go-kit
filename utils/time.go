package utils

import (
	"fmt"
	"reflect"
	"sort"
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

// "2024-07-18T02:01:03.212Z",    // UTC时间
// "2024-07-18T10:44:22+08:00", // 包含时区信息的时间
// 转换成2024-07-18 10:01:03
func TimeStrFormatCST(timeStr string) (string, error) {
	// 解析原始时间字符串
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return "", fmt.Errorf("error parsing time: %w", err)
	}
	// 将解析出的时间转换为中国标准时间（UTC+8）
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", fmt.Errorf("error loading location: %w", err)
	}
	tInCST := t.In(loc)

	// 按照指定格式输出时间
	return tInCST.Format("2006-01-02 15:04:05"), nil
}

func TimeStrFormatTime(timeStr string) (time.Time, error) {
	// 定义时间格式
	layout := "2006-01-02 15:04:05"
	// 使用Parse解析时间字符串
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %v", err)
	}
	return t, nil
}

func TimeSortStructsByField(slice interface{}, fieldName string, asc bool) (interface{}, error) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice, got %T", slice)
	}

	sortedSlice := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Cap())
	reflect.Copy(sortedSlice, rv)

	sort.SliceStable(sortedSlice.Interface(), func(i, j int) bool {
		iVal, jVal := sortedSlice.Index(i), sortedSlice.Index(j)

		// 如果是指针类型，需要解引用到实际的对象
		if iVal.Kind() == reflect.Ptr {
			iVal = iVal.Elem()
		}
		if jVal.Kind() == reflect.Ptr {
			jVal = jVal.Elem()
		}

		iField := iVal.FieldByName(fieldName)
		jField := jVal.FieldByName(fieldName)

		if !iField.IsValid() || !jField.IsValid() || iField.Kind() != reflect.Struct || jField.Kind() != reflect.Struct {
			return false
		}

		iTime, okI := iField.Interface().(time.Time)
		jTime, okJ := jField.Interface().(time.Time)
		if !okI || !okJ {
			return false
		}

		if asc {
			return iTime.Before(jTime)
		} else {
			return jTime.Before(iTime)
		}
	})

	return sortedSlice.Interface(), nil
}
