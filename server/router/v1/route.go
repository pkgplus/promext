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

	// /console/node/<NODEIP>/metrics
	party.Get("/console/node/:host/metrics", handlers.QueryAllMetrics)
	party.Get("/console/node/:host/metrics/:metric", handlers.QueryMetric)

	// /graph/node/<NODEIP>/metrics
	party.Get("/graph/node/:host/metrics", handlers.QueryRangeAllMetrics)
	party.Get("/graph/node/:host/metrics/:metric", handlers.QueryRangeMetric)
}
