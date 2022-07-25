package test_cover

import (
	"regexp"
	"strings"
)

// CheckPhoneNumber 校验手机号
func CheckPhoneNumber(phone string) bool {
	// 手机号长度11位
	if StringLength(phone) != 11 {
		return false
	}
	reg := regexp.MustCompile(`^1[3,4,5,6,7,8,9]{1}[0-9]{9}$`)
	return reg.MatchString(phone)
}

// StringLength 获取字符串长度
func StringLength(str string) int {
	return strings.Count(str, "") - 1
}
