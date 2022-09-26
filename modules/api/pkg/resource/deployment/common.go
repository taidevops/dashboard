package deployment

import (
	"github.com/taidevops/dashboard/src/app/backend/resource/dataselect"
	apps "k8s.io/api/apps/v1"
)

// The code below allows to perform complex data section on Deployment

type DeploymentCell apps.Deployment

func (self DeploymentCell) GetProperty(name dataselect.PropertyName) dataselect.ComparableValue {
	switch name {
	case dataselect.NameProperty:
		return dataselect.StdComparableString(self.ObjectMeta.Name)
	case dataselect.CreationTimestampProperty:
		return dataselect.StdComparableTime(self.ObjectMeta.CreationTimestamp.Time)
	case dataselect.NamespaceProperty:
		return dataselect.StdComparableString(self.ObjectMeta.Namespace)
	default:
		// if name is not supported then just returns a constant dummy value, sort will have no effect.
		return nil
	}
}