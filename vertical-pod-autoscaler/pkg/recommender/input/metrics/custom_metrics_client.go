package metrics

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
	"k8s.io/klog/v2"
	customClient "k8s.io/metrics/pkg/client/custom_metrics"
	"time"
)

// CustomMetricsSnapshot contains information about usage of certain container within defined time window.
type CustomMetricsSnapshot struct {
	//AppName represents the group of pods in the namespace
	AppName string
	// End time of the measurement interval.
	SnapshotTime time.Time
	// Duration of the measurement interval, which is [SnapshotTime - SnapshotWindow, SnapshotTime].
	SnapshotWindow time.Duration
	// Actual custom metrics
	CustomMetrics model.CustomMetrics
}

// CustomMetricsClient provides custom metrics on resources usage on pod level.
type CustomMetricsClient interface {
	// GetCustomMetrics returns a CustomMetricsSnapshot,
	// representing custom metrics for every running pod in the cluster
	GetCustomMetrics(namespace string, appName string) (*CustomMetricsSnapshot, error)
}

type customMetricsClient struct {
	metricsGetter     customClient.NamespacedMetricsGetter
	namespaceSelector string
}

// NewCustomMetricsClient creates new instance of CustomMetricsClient, which is used by recommender.
// It requires an instance of NamespacedMetricsGetter, which is used for underlying communication with metrics server.
// namespace limits queries to particular namespace, use k8sapiv1.NamespaceAll to select all namespaces.
func NewCustomMetricsClient(metricsGetter customClient.NamespacedMetricsGetter, namespaceSelector string) CustomMetricsClient {
	return &customMetricsClient{
		metricsGetter:     metricsGetter,
		namespaceSelector: namespaceSelector,
	}
}

func (c *customMetricsClient) GetCustomMetrics(namespace string, appName string) (*CustomMetricsSnapshot, error) {
	var metricsSnapshot = &CustomMetricsSnapshot{
		AppName:       appName,
		CustomMetrics: make(model.CustomMetrics),
	}

	metricsInterface := c.metricsGetter.NamespacedMetrics(namespace)

	for metric, metricName := range model.CustomMetricNames {
		klog.V(3).Infof("Getting metric value %s, apps=%s ", metricName, appName)

		metricValue, err := metricsInterface.GetForObject(schema.GroupKind{
			Group: "apps",
			Kind:  "Deployment",
		}, appName, metricName, labels.Everything())

		if err != nil {
			klog.V(1).Infof("Query metric %s error %v", metricName, err)
		} else {
			metricsSnapshot.SnapshotTime = metricValue.Timestamp.Time
			metricsSnapshot.SnapshotWindow = time.Duration(*metricValue.WindowSeconds)
			metricsSnapshot.CustomMetrics[metric] = model.CustomMetricValue(metricValue.Value)
		}
	}

	return metricsSnapshot, nil
}
