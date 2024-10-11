package configmaps

import (
	"context"
	"log"
	"strings"

	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func BackupConfigMaps(
	_ *rest.Config,
	kubeCl kubernetes.Interface,
	_ dynamic.Interface,
	namespaces []string,
) ([]runtime.Object, error) {
	namespaces = lo.Filter(namespaces, func(item string, _ int) bool {
		return strings.HasPrefix(item, "d8-") || strings.HasPrefix(item, "kube-")
	})

	configmaps := parallel.Map(namespaces, func(namespace string, _ int) []runtime.Object {
		list, err := kubeCl.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatalf("Failed to list configmaps from : %v", err)
		}

		return lo.Map(list.Items, func(item corev1.ConfigMap, _ int) runtime.Object {
			return &item
		})
	})

	return lo.Flatten(configmaps), nil
}
