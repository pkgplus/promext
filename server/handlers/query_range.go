package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/promext/pkg/models"
	"github.com/xuebing1110/promext/pkg/prometheus"
	"github.com/xuebing1110/promext/pkg/response"
)

func QueryRangeAllMetrics(ctx context.Context) {

}

func QueryRangeMetric(ctx context.Context) {
	metric := ctx.Params().Get("metric")
	host := ctx.Params().Get("host")

	// step
	step := ctx.URLParam("step")
	if step == "" {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(response.NewBadResponse(nil, "UrlParamErr", fmt.Errorf("\"step\" missing")))
		return
	}
	_, err := time.ParseDuration(step)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(response.NewBadResponse(nil, "UrlParamErr", fmt.Errorf("parse \"step\" failed")))
		return
	}

	// start & end
	var start, end time.Time
	duration := ctx.URLParam("duration")
	if duration != "" {
		d, err := time.ParseDuration(duration)
		if err != nil {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.JSON(response.NewBadResponse(nil, "UrlParamErr", fmt.Errorf("format duration failed: %v", err)))
			return
		}
		end = time.Now()
		start = end.Add(-d)
	} else {
		// start
		start_int64, err := ctx.URLParamInt64("start")
		if err != nil || start_int64 == 0 {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.JSON(response.NewBadResponse(nil, "UrlParamErr", fmt.Errorf("parse \"start\" to int64 failed")))
			return
		}
		start = time.Unix(start_int64, 0)

		// end
		end_int64, err := ctx.URLParamInt64("end")
		if err != nil || end_int64 == 0 {
			ctx.StatusCode(http.StatusBadRequest)
			ctx.JSON(response.NewBadResponse(nil, "UrlParamErr", fmt.Errorf("parse \"start\" to int64 failed")))
			return
		}
		end = time.Unix(end_int64, 0)
	}

	// metric
	m, found := models.GetMetric(metric)
	if !found {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(response.NewBadResponse(nil, "MetricErr", fmt.Errorf("metric %s unavailable", metric)))
		return
	}

	// query
	filter := fmt.Sprintf(`instance="%s"`, host)
	dsl := prometheus.GetRangeDsl(m.Dsl, filter, step)
	data, err := prometheus.QueryRange(metric, dsl, start, end, step)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(response.NewBadResponse(nil, "PrometheusErr", err))
		return
	}

	ctx.JSON(response.NewSucResponse(data))

}
