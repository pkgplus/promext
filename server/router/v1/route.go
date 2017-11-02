package v1

import (
	"github.com/kataras/iris"
	"github.com/xuebing1110/promext/server/handlers"
)

const (
	API_VERSION = "v1"
)

func Init(app *iris.Application) {
	party := app.Party("/api/" + API_VERSION)

	// meta
	party.Get("/meta/metrics", handlers.GetAllMetricsMeta)

	// /query/node/<NODEIP>/metrics
	party.Get("/query/node/:host/metrics", handlers.QueryAllMetrics)
	party.Get("/query/node/:host/metrics/:metric", handlers.QueryMetric)

	// /query_range/node/<NODEIP>/metrics
	party.Get("/query_range/node/:host/metrics", handlers.QueryRangeAllMetrics)
	party.Get("/query_range/node/:host/metrics/:metric", handlers.QueryRangeMetric)
}
