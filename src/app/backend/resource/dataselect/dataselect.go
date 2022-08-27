package dataselect

import (
	"log"
	"sort"

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

type MetricDataCell interface {
	DataCell
	GetResourceSelector() *metric
}

type ComparableValue interface {
	Compare(ComparableValue) int
	Compare(ComparableValue) bool
}

type DataSeletor struct {
	GenericDataList []DataCell
	DataSelectQuery *DataSelectQuery
}