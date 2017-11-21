package prometheus

import (
	"strings"
)

func GetDsl(template, filter string) string {
	return strings.Replace(template, "$filter", filter, -1)
}

func GetRangeDsl(template, filter, duration string) string {
	ret := strings.Replace(template, "$filter", filter, -1)
	return strings.Replace(ret, "$duration", duration, -1)
}
