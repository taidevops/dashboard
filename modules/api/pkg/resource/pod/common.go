package pod

import (
	v1 "k8s.io/api/core/v1"
)

// getRestartCount return the restart count of given pod (total number of its containers restarts).
func getRestartCount(pod v1.Pod) int32 {
	var restartCount int32 = 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restartCount += containerStatus.RestartCount
	}
	return restartCount
}
