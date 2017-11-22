package handlers

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/context"
	"github.com/prometheus/common/model"

	"github.com/xuebing1110/promext/pkg/models"
	"github.com/xuebing1110/promext/pkg/prometheus"
	"github.com/xuebing1110/promext/pkg/response"
)

func QueryRangeAllMetrics(ctx context.Context) {
	logger := ctx.Application().Logger()
	host := ctx.Params().Get("host")

	// parameter
	start, end, step, err := GetStartEndStep(ctx)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(response.NewBadResponse(nil, "UrlParamErr", err))
		return
	}

	// filter
	var filter string
	if host != "" {
		filter = GetFilterFromParams(map[string]string{"filter_instance": host})
	} else {
		filter = GetFilterFromParams(ctx.URLParams())
	}

	// query
	data := make([]*model.SampleStream, 0)
	for _, m := range models.MetricsRange {
		dsl := prometheus.GetRangeDsl(m.Dsl, filter, step)
		ret, err := prometheus.QueryRange(m.Name, dsl, start, end, step)
		if err != nil {
			logger.Errorf("query %s for %s failed: %v", m.Name, host, err)
			continue
		}

		for _, item := range ret {
			data = append(data, item)
		}
	}

	// response
	ctx.JSON(response.NewSucResponse(data))
}

func QueryRangeMetric(ctx context.Context) {
	metric := ctx.Params().Get("metric")
	host := ctx.Params().Get("host")

	// parameter
	start, end, step, err := GetStartEndStep(ctx)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(response.NewBadResponse(nil, "UrlParamErr", err))
		return
	}

	// metric
	m, found := models.GetMetric(metric)
	if !found {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(response.NewBadResponse(nil, "MetricErr", fmt.Errorf("metric %s unavailable", metric)))
		return
	}

	// filter
	var filter string
	if host != "" {
		filter = GetFilterFromParams(map[string]string{"filter_instance": host})
	} else {
		filter = GetFilterFromParams(ctx.URLParams())
	}

	// query
	dsl := prometheus.GetRangeDsl(m.Dsl, filter, step)
	data, err := prometheus.QueryRange(metric, dsl, start, end, step)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(response.NewBadResponse(nil, "PrometheusErr", err))
		return
	}

	ctx.JSON(response.NewSucResponse(data))

}
