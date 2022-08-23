package model

import "k8s.io/apimachinery/pkg/api/resource"

const (
	// LabelJvm represents the service container is a JVM service
	LabelJvm string = "jvm"
	// LabelApp represents the group of the pods with same service container
	LabelApp string = "app"
	// AnnotationJvmInContainers indicates the JVM containers in the pod
	AnnotationJvmInContainers string = "autoscaling.k8s.io/jvm_in_containers"
	// AnnotationJvmMaxRamPercentage indicates the JVM max ram percentage of the max container memory limit
	AnnotationJvmMaxRamPercentage string = "autoscaling.k8s.io/jvm_max_ram_percentage"
	// DefaultJvmInContainer default container runs JVM
	DefaultJvmInContainer string = "app"
	// DefaultJvmMaxRamPercentage represents the default jvm max ram percentage
	DefaultJvmMaxRamPercentage int64 = 80
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
)

// CustomMetricNames is a map from CustomMetricName to the custom metric name in API
var CustomMetricNames = map[CustomMetricName]string{
	MetricJvmHeapUtil:       "namespace_app_pod_jvm_memory_heap_utilization",
	MetricJvmGcPauseSeconds: "namespace_app_pod_jvm_gc_pause_seconds_avg",
}

// CustomMetrics is a map from CustomMetricName to the corresponding CustomMetricValue
type CustomMetrics map[CustomMetricName]resource.Quantity
