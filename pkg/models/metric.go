package models

import (
	"io/ioutil"
	"os"
	"github.com/go-yaml/yaml"
)

const (
	MetricType_Current = "current"
	MetricType_Range   = "range"
)

var (
	MetricsDict    []Metric
	metricsDictMap = map[string]Metric{}
	MetricsCurrent []Metric
	MetricsRange   []Metric
)

type Metric struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Dsl  string `json:"dsl"`
}

func initConfig() {
	yamlFile, err := ioutil.ReadFile("./config/" + os.Getenv("job") + ".yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &MetricsDict)
	if err != nil {
		panic(err)
	}
}

func init() {
	initConfig()
	MetricsCurrent = make([]Metric, 0)
	metricsDictMap = make(map[string]Metric)
	for _, m := range MetricsDict {
		metricsDictMap[m.Name] = m
		if m.Type == MetricType_Current {
			MetricsCurrent = append(MetricsCurrent, m)
		} else {
			MetricsRange = append(MetricsRange, m)
		}
	}
}

func GetMetric(name string) (Metric, bool) {
	m, found := metricsDictMap[name]
	return m, found
}
