package service

import (
	"github.com/taidevops/dashboard/src/app/backend/resource/dataselect"
	v1 "k8s.io/api/core/v1"
)

type ServiceCell v1.Service

func (self ServiceCell) GetProperty(name dataselect.PropertyName) dataselect.ComparableValue {
	switch name {
	case dataselect.NameProperty:
		return dataselect.StdComparableString
	}
}