package models

const (
	MetricType_Current = "current"
	MetricType_Range   = "range"
)

var (
	MetricsDict = []Metric{
		Metric{"cpuUtilization", MetricType_Current, `node_cpu_not_idle_rate{$filter}*100`},
		Metric{"cpuUtilizationAvg", MetricType_Range, `avg_over_time(node_cpu_not_idle_rate{$filter}[$duration])*100`},
		Metric{"cpuUtilizationMax", MetricType_Range, `max_over_time(node_cpu_not_idle_rate{$filter}[$duration])*100`},
		Metric{"cpuUtilizationMin", MetricType_Range, `min_over_time(node_cpu_not_idle_rate{$filter}[$duration])*100`},
		Metric{"memoryUtilization", MetricType_Current, `100 - ( node_memory_MemAvailable{$filter} OR node_memory_MemAvailable_ext{$filter})/ node_memory_MemTotal{$filter} * 100`},
		Metric{"memoryUtilizationAvg", MetricType_Range, `100 - (avg_over_time(node_memory_MemAvailable{$filter}[$duration]) OR avg_over_time(node_memory_MemAvailable_ext{$filter}[$duration])) / node_memory_MemTotal{$filter} * 100`},
		Metric{"memoryUtilizationMax", MetricType_Range, `100 - (max_over_time(node_memory_MemAvailable{$filter}[$duration]) OR max_over_time(node_memory_MemAvailable_ext{$filter}[$duration])) / node_memory_MemTotal{$filter} * 100`},
		Metric{"memoryUtilizationMin", MetricType_Range, `100 - (min_over_time(node_memory_MemAvailable{$filter}[$duration]) OR min_over_time(node_memory_MemAvailable_ext{$filter}[$duration])) / node_memory_MemTotal{$filter} * 100`},
		Metric{"diskUtilization", MetricType_Current, `100 - node_filesystem_free{$filter, fstype!~"proc|overlay|cgroup|rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs|iso.+"} / node_filesystem_size * 100`},
		Metric{"diskUtilizationMax", MetricType_Range, `100 - (max_over_time(node_filesystem_free{$filter, fstype!~"proc|overlay|cgroup|rootfs|selinuxfs|autofs|rpc_pipefs|tmpfs|iso.+"}[$duration])) / node_filesystem_size * 100`},
	}

	metricsDictMap = map[string]Metric{}
	MetricsCurrent = []Metric{}
	MetricsRange   = []Metric{}
)

type Metric struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Dsl  string `json:"dsl"`
}

func init() {
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
