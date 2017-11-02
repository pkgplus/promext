package handlers

import (
	"github.com/kataras/iris/context"
	"github.com/xuebing1110/promext/pkg/models"
	"github.com/xuebing1110/promext/pkg/response"
)

func GetAllMetricsMeta(ctx context.Context) {
	ctx.JSON(response.NewSucResponse(models.MetricsDict))
}
