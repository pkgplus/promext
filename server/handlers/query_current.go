package handlers

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/promext/pkg/models"
	"github.com/xuebing1110/promext/pkg/prometheus"
	"github.com/xuebing1110/promext/pkg/response"
)

func QueryCurrentAllMetrics(ctx context.Context) {

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

	filter := fmt.Sprintf(`instance="%s"`, host)
	data, err := prometheus.QueryCurrent(metric, prometheus.GetDsl(m.Dsl, filter))
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(response.NewBadResponse(nil, "PrometheusErr", err))
		return
	}

	ctx.JSON(response.NewSucResponse(data))
}
