package dataselect

import (
	"log"
	"sort"

)

type DataCell interface {
	GetProperty(PropertyName) ComparableValue
}

type MetricDataCell interface {

}

type ComparableValue interface {
	Compare(ComparableValue) int
	Compare(ComparableValue) bool
}

type DataSeletor struct {
	GenericDataList []DataCell
	DataSelectQuery *DataSelectQuery
}