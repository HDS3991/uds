package utils

import "regexp"

// 邮政编码校验
func CheckPostcode(code string) bool {
	reg := `^\d{6}$`
	return regexp.MustCompile(reg).MatchString(code)
}
