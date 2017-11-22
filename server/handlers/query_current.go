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

func QueryCurrentAllMetrics(ctx context.Context) {
	logger := ctx.Application().Logger()
	host := ctx.Params().Get("host")

	// filter
	var filter string
	if host != "" {
		filter = GetFilterFromParams(map[string]string{"filter_instance": host})
	} else {
		filter = GetFilterFromParams(ctx.URLParams())
	}

	// query
	data := make([]*model.Sample, 0)
	for _, m := range models.MetricsCurrent {
		ret, err := prometheus.QueryCurrent(m.Name, prometheus.GetDsl(m.Dsl, filter))
		if err != nil {
			logger.Errorf("query %s for %s failed: %v", m.Name, host, err)
			continue
		}

		for _, item := range ret {
			data = append(data, item)
		}
	}

	ctx.JSON(response.NewSucResponse(data))
}

func QueryCurrentMetric(ctx context.Context) {
	metric := ctx.Params().Get("metric")
	host := ctx.Params().Get("host")

	m, found := models.GetMetric(metric)
	if !found {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(response.NewBadResponse(nil, "MetricErr", fmt.Errorf("metric %s unavailable", metric)))
		return
	}

	// filter := fmt.Sprintf(`instance="%s"`, host)
	var filter string
	if host != "" {
		filter = GetFilterFromParams(map[string]string{"filter_instance": host})
	} else {
		filter = GetFilterFromParams(ctx.URLParams())
	}

	// query
	data, err := prometheus.QueryCurrent(metric, prometheus.GetDsl(m.Dsl, filter))
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(response.NewBadResponse(nil, "PrometheusErr", err))
		return
	}

	ctx.JSON(response.NewSucResponse(data))
}
