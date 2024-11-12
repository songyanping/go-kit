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

func TimeSortStructsByFieldTime(slice interface{}, fieldName string, asc bool) (interface{}, error) {
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

func TimeSortStructsByFieldString(slice interface{}, fieldName string, asc bool) (interface{}, error) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice, got %T", slice)
	}

	sortedSlice := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Cap())
	reflect.Copy(sortedSlice, rv)

	sort.SliceStable(sortedSlice.Interface(), func(i, j int) bool {
		iVal, jVal := sortedSlice.Index(i), sortedSlice.Index(j)

		// 如果是指针类型，需要解引用
		if iVal.Kind() == reflect.Ptr {
			iVal = iVal.Elem()
		}
		if jVal.Kind() == reflect.Ptr {
			jVal = jVal.Elem()
		}

		iField := iVal.FieldByName(fieldName)
		jField := jVal.FieldByName(fieldName)

		// 检查字段是否存在
		if !iField.IsValid() || !jField.IsValid() {
			return false
		}

		// 将字段值作为字符串读取
		iStr, okI := iField.Interface().(string)
		jStr, okJ := jField.Interface().(string)
		if !okI || !okJ {
			return false
		}

		// 解析字符串到 time.Time
		layout := "2006-01-02 15:04:05" // 确保这个格式与你的时间字符串匹配
		iTime, errI := time.Parse(layout, iStr)
		jTime, errJ := time.Parse(layout, jStr)
		if errI != nil || errJ != nil {
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

// ConvertTimestampMillisToCST 将以毫秒为单位的int64格式的UNIX时间戳转换为CST时间的字符串表示
func TimeConvertMillisString(timestampMillis int64) (string, error) {
	// 将毫秒转换为秒和纳秒
	seconds := timestampMillis / 1000
	nanoseconds := (timestampMillis % 1000) * 1000000

	// 使用time.Unix函数将时间戳转换为time.Time格式
	t := time.Unix(seconds, nanoseconds)

	// 加载CST时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", fmt.Errorf("加载时区失败: %v", err)
	}

	// 转换到CST时区
	cstTime := t.In(loc)

	// 返回转换后的CST时间的字符串表示
	return cstTime.Format(time.DateTime), nil
}

func TimeConvertSecondsString(timestampSeconds int64) (string, error) {
	// 使用time.Unix函数将时间戳转换为time.Time格式
	t := time.Unix(timestampSeconds, 0)

	// 加载CST时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", fmt.Errorf("加载时区失败: %v", err)
	}

	// 转换到CST时区
	cstTime := t.In(loc)

	// 返回转换后的CST时间的字符串表示
	return cstTime.Format(time.DateTime), nil
}

// getStartAndEndOfWeek 返回给定时间所在周的开始和结束时间
func GetStartAndEndOfWeek(t time.Time) (start, end time.Time) {
	// 确定周一为每周的开始
	start = time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	end = start.AddDate(0, 0, 6) // 向后加6天得到周日
	return
}
