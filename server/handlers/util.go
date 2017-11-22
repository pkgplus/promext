package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/kataras/iris/context"
)

func GetStartEndStep(ctx context.Context) (start, end time.Time, step string, err error) {
	// step
	step = ctx.URLParam("step")
	if step == "" {
		err = fmt.Errorf("\"step\" missing")
		return
	}
	_, err = time.ParseDuration(step)
	if err != nil {
		err = fmt.Errorf("parse \"step\" failed")
	}

	// start & end
	duration := ctx.URLParam("duration")
	if duration != "" {
		d, err_tmp := time.ParseDuration(duration)
		if err_tmp != nil {
			err = fmt.Errorf("format duration failed: %v", err_tmp)
			return
		}
		end = time.Now()
		start = end.Add(-d)
	} else {
		// start
		start_int64, err_tmp := ctx.URLParamInt64("start")
		if err_tmp != nil || start_int64 == 0 {
			err = fmt.Errorf("parse \"start\" to int64 failed: %v", err_tmp)
			return
		}
		start = time.Unix(start_int64, 0)

		// end
		end_int64, err_tmp := ctx.URLParamInt64("end")
		if err_tmp != nil || end_int64 == 0 {
			err = fmt.Errorf("parse \"start\" to int64 failed: %v", err_tmp)
			return
		}
		end = time.Unix(end_int64, 0)
	}

	return
}

func GetFilterFromParams(params map[string]string) string {
	var filter string
	for name, value := range params {
		if !strings.HasPrefix(name, "filter_") {
			continue
		}

		labelName := string([]byte(name)[7:])
		if filter == "" {
			filter = fmt.Sprintf("%s=\"%s\"", labelName, value)
		} else {
			filter = filter + fmt.Sprintf(",%s=%s", labelName, value)
		}
	}
	return filter
}
