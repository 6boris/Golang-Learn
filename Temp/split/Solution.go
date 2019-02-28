package Solution

import (
	"strings"
)

//	截取指定字符串前面的字符
func Split(str, sp string) string {
	return str[:strings.Index(str, sp)]
}
