package deployment

import (
	"github.com/taidevops/dashboard/src/app/backend/resource/dataselect"
	client "k8s.io/client-go/kubernetes"
)

func GetDeploymentPods(client client.Interface, dsQuery *dataselect.DataSelectQuery, namespace, deploymentName string) ()