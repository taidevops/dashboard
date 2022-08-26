package dataselect

type DataSelectQuery struct {
	PaginationQuery *PaginationQuery
	SortQuery       *SortQuery
	FilterQuery     *FilterQuery
	MetricQuery     *MetricQuery
}

var MoMetrics = 

type MetricQuery struct {
	MetricName []string

	Aggregations metricapi.AggregationModes
}

func NewMetricQuery(metricNames []string, aggregations metricapi.AggregationModes) *MetricQuery {
	return &MetricQuery{
		MetricNames:  metricNames,
		Aggregations: aggregations,
	}
}

type SortQuery struct {
	SortByList []SortBy
}

type SortBy struct {
	Property  PropertyName
	Ascending bool
}

var NoSort = &SortQuery{
	SortByList: []SortBy{},
}

type FilterQuery struct {
	FilterByList []FilterBy
}

type FilterBy struct {
	Property PropertyName
	Value    ComparableValue
}

var NoFilter = &FilterQuery{
	FilterByList: []FilterBy{},
}

