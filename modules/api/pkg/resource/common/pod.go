package common

func FilterDeploymentPodsByOnwerRefrence(deployment apps.Deployment, allRS []apps.ReplicaSet,
	allPods []v1.Pod) []v1.Pod {
	var matchingPods []v1.Pod
	for _, rs := range allRS {
		if metav1.IsControlledBy(&rs, &deployment) {
			matchingPods = append(matchingPods, Fil)
		}
	}
}