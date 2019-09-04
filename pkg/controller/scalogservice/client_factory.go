package scalogservice

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
	newClientDeployment creates a Kubernetes Deployment used for client test
*/
func newClientDeployment() *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-client-deployment",
			Namespace: "scalog",
			Labels: map[string]string{
				"app": "scalog-client",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "scalog-client",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "scalog-client",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "scalog-service-account",
					Containers: []corev1.Container{
						corev1.Container{
							Name:            "scalog-client-node",
							Image:           "scalog/scalog:latest",
							Command:         []string{"./scalog"},
							Args:            []string{"k8sclient"},
							ImagePullPolicy: "Always",
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{ContainerPort: 21024},
							},
							Env: []corev1.EnvVar{
								corev1.EnvVar{
									Name: "NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
								corev1.EnvVar{
									Name: "NAMESPACE",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
								corev1.EnvVar{
									Name: "POD_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "status.podIP",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
