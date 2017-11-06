package models

const (
	MetricType_Scalar = "scalar"
	MetricType_Vector = "vector"
)

var (
	MetricsDict = []Metric{
		Metric{"cpuUtilization", MetricType_Scalar},
		Metric{"cpuUtilizationMax", MetricType_Scalar},
		Metric{"cpuUtilizationMin", MetricType_Scalar},
		Metric{"memoryUtilization", MetricType_Scalar},
		Metric{"memoryUtilizationMax", MetricType_Scalar},
		Metric{"memoryUtilizationMin", MetricType_Scalar},
		Metric{"diskUtilization", MetricType_Vector},
	}
)

type Metric struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
