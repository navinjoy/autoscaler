package model

import "k8s.io/apimachinery/pkg/api/resource"

const (
	// LabelJvm represents the service container is a JVM service
	LabelJvm string = "jvm"
	// LabelApp represents the group of the pods with same service container
	LabelApp string = "app"
	// AnnotationAppContainer indicates the JVM containers in the pod
	AnnotationAppContainer string = "autoscaling.k8s.io/app_container"
	// DefaultAppContainer default container runs JVM
	DefaultAppContainer string = "app"
)

// CustomMetricName represents the name of the custom metric
type CustomMetricName string

// namespace_app_pod_jvm_memory_heap_utilization
// namespace_app_pod_jvm_gc_pause_seconds_avg
const (
	// MetricJvmHeapUtil represents Jvm Heap Utilization percentage.
	MetricJvmHeapUtil CustomMetricName = "jvmHeapUtil"
	// MetricJvmGcPauseSeconds represents Jvm GC pause seconds
	MetricJvmGcPauseSeconds CustomMetricName = "jvmGcPauseSeconds"
	// MetricAppContainerMaxMemoryLimits represents the maximum memory limit across pods with same app label
	MetricAppContainerMaxMemoryLimits CustomMetricName = "appContainerMaxMemoryLimits"
	// MetricAppContainerMinMemoryLimits represents the minimum memory limit across pods with same app label
	MetricAppContainerMinMemoryLimits CustomMetricName = "appContainerMinMemoryLimits"
)

// CustomMetricNames is a map from CustomMetricName to the custom metric name in API
var CustomMetricNames = map[CustomMetricName]string{
	MetricJvmHeapUtil:                 "namespace_app_pod_jvm_memory_heap_utilization",
	MetricJvmGcPauseSeconds:           "namespace_app_pod_jvm_gc_pause_seconds_avg",
	MetricAppContainerMaxMemoryLimits: "namespace_app_container_app_max_memory_limits",
	MetricAppContainerMinMemoryLimits: "namespace_app_container_app_min_memory_limits",
}

// CustomMetrics is a map from CustomMetricName to the corresponding CustomMetricValue
type CustomMetrics map[CustomMetricName]resource.Quantity
