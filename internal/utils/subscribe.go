package utils

import (
	"strings"

	"github.com/gogf/gf/v2/util/grand"
)

// GetRandomString 获取随机字符串，a,b,c,d,e 这样的格式随机获取一个
func GetRandomString(v string) string {
	if strings.Contains(v, ",") {
		vs := strings.Split(v, ",")
		return vs[grand.Intn(len(vs))]
	}
	return v
}
