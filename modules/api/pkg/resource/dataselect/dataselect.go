package dataselect

import (
	metricapi "github.com/taidevops/dashboard/src/app/backend/integration/metric/api"
)

// GenericDataCell describes the interface of the data cell that contains all the necessary methods needed to perform
// complex data selection
// GenericDataSelect takes a list of these interfaces and performs selection operation.
// Therefore as long as the list is composed of GenericDataCells you can perform any data selection!
type DataCell interface {
	// GetPropertyAtIndex returns the property of this data cell.
	// Value returned has to have Compare method which is required by Sort functionality of DataSelect.
	GetProperty(PropertyName) ComparableValue
}

// MetricDataCell extends interface of DataCells and additionally supports metric download.
type MetricDataCell interface {
	DataCell
	// GetResourceSelector returns ResourceSelector for this resource. The ResourceSelector can be used to get.
	// HeapsterSelector which in turn can be used to download metrics.
	GetResourceSelector() *metricapi.ResourceSelector
}

// ComparableValue hold any value that can be compared to its own kind.
type ComparableValue interface {
	// Compares self with other value. Returns 1 if other value is smaller, 0 if they are the same, -1 if other is larger.
	Compare(ComparableValue) int
	// Returns true if self value contains or is equal to other value, false otherwise.
	Contains(ComparableValue) bool
}

// SelectableData contains all the required data to perform data selection.
// It implements sort.Interface so its sortable under sort.Sort
// You can use its Select method to get selected GenericDataCell list.
type DataSelector struct {
	// GenericDataList hold generic data cells that are being selected.
	GenericDataList []DataCell
	// DataSelectQuery holds instructions for data select.
	DataSelectQuery *DataSelectQuery
	// CachedResources stores resources that may be needed during data selection process
	CachedResources *metricapi.CachedResources
	// CumulativeMetricsPromises is a list of promises holding aggregated metrics for resources in GenericDataList.
	// The metrics will be calculated after calling GetCumulativeMetrics method.
	CumulativeMetricsPromises metricapi.MetricPromises
	// MetricsPromises is a list of promises holding metrics for resources in GenericDataList.
	// The metrics will be calculated after calling GetMetrics method. Metric will not be
	// aggregated and can are used to display sparklines on pod list.
	MetricsPromises metricapi.MetricPromises
}

// Implementation of sort.Interface so that we can use built-in sort function (sort.Sort) for sorting SelectableData

// Len returns the length of data inside SelectableData.
func (self DataSelector) Len() int { return len(self.GenericDataList) }

// Swap swaps 2 indices inside SelectableData.
func (self DataSelector) Swap(i, j int) {
	self.GenericDataList[i], self.GenericDataList[j] = self.GenericDataList[j], self.GenericDataList[i]
}

func (self DataSelector) Less(i, j int) bool {
	return false
}

func (self *DataSelector) Sort() *DataSelector {
	return self
}

func (self *DataSelector) Filter() *DataSelector {
	return self
}

func (self *DataSelector) GetMetrics(metricClient metricapi.MetricClient) *DataSelector {
	return self
}

func PodListMetrics(dataList []DataCell, dsQuery *DataSelectQuery,
	metricClient metricapi.MetricClient) metricapi.MetricPromises {
	selectableData := DataSelector{
		GenericDataList: dataList,
		DataSelectQuery: dsQuery,
		CachedResources: metricapi.NoResourceCache,
	}

	processed := selectableData.GetMetrics(metricClient)
	return processed.MetricsPromises
}
