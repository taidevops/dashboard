package common

import (
	metricapi "github.com/taidevops/dashboard/src/app/backend/integration/metric/api"
)

type SortableInt64 []int64

func (a SortableInt64) Len() int           { return len(a) }
func (a SortableInt64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableInt64) Less(i, j int) bool { return a[i] < a[j] }

func AggregateData(metricList []metricapi.Metric, metricName string,
	aggregationName metricapi.AggregationMode) metricapi.Metric {

	return metricapi.Metric{}
}

func AggregatingMapFromDataList(metricList []metricapi.Metric, metricName string) (
	map[int64][]int64, metricapi.Label) {
	newLabel := metricapi.Label{}

	aggrMap := make(map[int64][]int64, 0)
	return aggrMap, newLabel
}

func AggregateMetricPromises(metricPromises metricapi.MetricPromises, metricName string,
	aggregations metricapi.AggregationModes, forceLabel metricapi.Label) metricapi.MetricPromises {
	return metricPromises
}
