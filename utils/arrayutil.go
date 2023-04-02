package utils

import (
	"fmt"
	"strings"
)

type ArrayUtil struct{}

func NewArrayUtil() *ArrayUtil {
	return &ArrayUtil{}
}

/**
 * @description:
 * @param {[]string} array 字符串数组，如["jack","even","Eyck","Geoff"]
 * @param {string} value 需要判断是否在数组中的值，如Geoff
 * @return {bool} if array contain value,return true, else return false.
 */
func (u *ArrayUtil) IsInArray(array []string, value string) bool {
	var flag bool = false
	if len(array) == 0 {
		return false
	}
	for _, v := range array {
		if v == value {
			flag = true
			break
		}
	}
	return flag
}

/**
 * @description: 将数组转换成逗号相隔的字符串（逗号间无空格）
 * @param {[]string} array_or_slice
 * @return {string} 逗号相隔的字符串。
 */
func (u *ArrayUtil) ArrayToStr(array_or_slice []string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
}
