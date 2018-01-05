package prometheus

import (
	"strings"
)

func GetDsl(template, filter string) string {
	if filter == "" {
		template = strings.Replace(template, "$filter,", "", -1)
		return strings.Replace(template, "$filter", filter, -1)
	} else {
		return strings.Replace(template, "$filter", filter, -1)
	}
}

func GetRangeDsl(template, filter, duration string) string {
	ret := GetDsl(template, filter)
	return strings.Replace(ret, "$duration", duration, -1)
}
