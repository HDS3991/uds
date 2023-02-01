package utils

import (
	"fmt"
	"regexp"
)

// 邮政编码校验
func CheckPostcode[T comparable](code T) bool {
	reg := `^\d{6}$`
	return regexp.MustCompile(reg).MatchString(fmt.Sprintf("%v", code))
}
