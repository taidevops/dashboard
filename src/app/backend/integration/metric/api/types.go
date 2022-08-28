package api

import (
	"time"

	"github.com/taidevops/dashboard/src/app/backend/api"
	"k8s.io/apimachinery/pkg/types"
)

type MetricClient interface {
	DownloadMetric(selectors []ResourceSelector, metricName string)
}

// AggregationMode informs how data should be aggregated (sum, min, max)
type AggregationMode string

// Aggregation modes which should be used for data aggregation. Eg. [sum, min, max].
const (
	SumAggregation     = "sum"
	MaxAggregation     = "max"
	MinAggregation     = "min"
	DefaultAggregation = SumAggregation
)

type AggregationModes []AggregationMode

var OnlySumAggregation = AggregationModes{SumAggregation}
var OnlyDefaultAggregation = AggregationModes{DefaultAggregation}

// ResourceSelector is a structure used to quickly and uniquely identify given resource.
// This struct can be later used for heapster data download etc.
type ResourceSelector struct {
	// Namespace of this resource.
	Namespace string
	// Type of this resource
	ResourceType api.ResourceKind
	// Name of this resource.
	ResourceName string
	// Selector used to identify this resource (should be used only for Deployments!).
	Selector map[string]string
	// UID is resource unique identifier.
	UID types.UID
}

const (
	CpuUsage    = "cpu/usage_rate"
	MemoryUsage = "memory/usage"
)

type DataPoints []DataPoint

type DataPoint struct {
	X int64 `json:"x"`
	Y int64 `json:y`
}

type MetricPoint struct {
	Timestamp time.Time `json:"timestamp"`
	Value     uint64    `json:"value"`
}

// Label stores information about identity of resources (UIDs) described by metric.
type Label map[api.ResourceKind][]types.UID

// AddMetricLabel returns a unique combined Label of self and other resource.
// New label describes both resources.
func (self Label) AddMetricLabel(other Label) Label {
	if other == nil {
		return self
	}

	uniqueMap := map[types.UID]bool{}
	for _, v := range self {
		for _, t := range v {
			uniqueMap[t] = true
		}
	}

	for k, v := range other {
		for _, t := range v {
			if _, exists := uniqueMap[t]; !exists {
				self[k] = append(self[k], t)
			}
		}
	}
	return self
}
