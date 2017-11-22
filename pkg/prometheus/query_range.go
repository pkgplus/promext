package prometheus

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func QueryRange(metric, dsl string, start, end time.Time, step string) (mvs []*model.SampleStream, err error) {
	logger.Infof("dsl: %s, start=%s,end=%s,step=%s", dsl, start.String(), end.String(), step)

	var step_d time.Duration
	step_d, err = time.ParseDuration(step)
	if err != nil {
		return
	}

	value, e := api.QueryRange(
		context.Background(),
		dsl,
		v1.Range{Start: start, End: end, Step: step_d},
	)
	if e != nil {
		err = e
		return
	}

	matrixV := value.(model.Matrix)
	mvs = []*model.SampleStream(matrixV)
	for _, mv := range mvs {
		mv.Metric["_name_"] = model.LabelValue(metric)
	}
	return mvs, nil
}
