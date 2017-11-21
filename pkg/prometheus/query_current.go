package prometheus

import (
	"context"
	"time"

	"github.com/prometheus/common/model"
)

func QueryCurrent(metric, dsl string) (mvs []*model.Sample, err error) {
	value, e := api.Query(context.Background(), dsl, time.Now())
	if e != nil {
		err = e
		return
	}

	vectorV := value.(model.Vector)
	mvs = []*model.Sample(vectorV)
	for _, mv := range mvs {
		mv.Metric["_name_"] = model.LabelValue(metric)
	}
	return mvs, nil
}
