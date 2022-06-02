package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	AnnotationPrefixSchedulingExpansion = "expansion.scheduling.crane.io"
	AnnotationPrefixSchedulingBalanceLoad = "load.balance.scheduling.crane.io"
	AnnotationPrefixSchedulingBalanceTarget = "target.balance.scheduling.crane.io"

)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ClusterNodeResourcePolicyList is a list of ClusterNodeResourcePolicy resources
type ClusterNodeResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterNodeResourcePolicy `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster,shortName=cnrp
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterNodeResourcePolicy must be created in crane root namespace
type ClusterNodeResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterNodeResourcePolicySpec   `json:"spec,omitempty"`
	Status ClusterNodeResourcePolicyStatus `json:"status,omitempty"`
}

type NodeResourceApplyMode string

const (
	NodeResourceApplyModeCRD        NodeResourceApplyMode = "crd"
	NodeResourceApplyModeAnnotation NodeResourceApplyMode = "annotation"
)

type ClusterNodeResourcePolicySpec struct {
	// NodeSelector used to select target nodes which use the same resource policy param
	// +optional
	NodeSelector metav1.LabelSelector `json:"nodeSelector,omitempty"`
	// ApplyMode decides how the controller to get the policy data to kubernetes node object
	// annotation: ClusterNodeResourcePolicy controller will patch policy data & metric data to node annotation directly, do not create other NodeResourcePolicy crds
	// crd: ClusterNodeResourcePolicy controller will create NodeResourcePolicy crds to associate with node with the same name.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=annotation;crd
	// +kubebuilder:default=annotation
	ApplyMode NodeResourceApplyMode `json:"applyMode,omitempty"`
	// Template specify the resource policy which including node resource policy and so on.
	// +optional
	Template *NodeResourcePolicyTemplate `json:"template,omitempty"`
}

type ClusterNodeResourcePolicyStatus struct {
	// DesiredNumberCreated is desired number of created NodeResourcePolicys which associated nodes by nodeSelector
	// +optional
	DesiredNumberCreated int `json:"desiredNumberCreated,omitempty"`
	// DesiredNumberCreated is current number of created NodeResourcePolicys which associated nodes by nodeSelector
	// +optional
	CurrentNumberCreated int `json:"currentNumberCreated,omitempty"`
	// Conditions is the condition status of ClusterNodeResourcePolicy
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type NodeResourcePolicyTemplate struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeResourcePolicySpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourcePolicyList is a list of NodeResourcePolicy resources
type NodeResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodeResourcePolicy `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,shortName=nrp
// +kubebuilder:printcolumn:name="Name",type="string",JSONPath=".metadata.name",description="The name of the nrp."
// +kubebuilder:printcolumn:name="ResourceExpansionStrategy",type="string",JSONPath=".spec.resourceExpansionStrategyType",description="The resource expansion strategy of nrp."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."
// +kubebuilder:webhooks:path=/mutate-noderesourcepolicy,mutating=true,failurePolicy=fail,groups=scheduling.crane.io,resources=noderesourcepolicy,verbs=create;update,versions=v1alpha1,name=noderesourcepolicys.webhook.scheduling.crane.io,sideEffects=none,admissionReviewVersions=v1
// +kubebuilder:webhooks:verbs=create;update,path=/validate-noderesourcepolicy,mutating=false,failurePolicy=fail,groups=scheduling.crane.io,resources=noderesourcepolicy,versions=v1,name=noderesourcepolicys.webhook.scheduling.crane.io,sideEffects=none,admissionReviewVersions=v1
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourcePolicy
type NodeResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeResourcePolicySpec   `json:"spec,omitempty"`
	Status NodeResourcePolicyStatus `json:"status,omitempty"`
}

type ResourceExpansionStrategyType string

const (
	ResourceExpansionStrategyTypeTypeStatic ResourceExpansionStrategyType = "static"
	ResourceExpansionStrategyTypeTypeAuto   ResourceExpansionStrategyType = "auto"
)

type NodeResourcePolicySpec struct {
	// TargetRef is the target node the NodeResourcePolicy associated
	// +optional
	TargetRef *autoscalingv2.CrossVersionObjectReference `json:"targetRef"`
	// ResourceExpansionStrategy is the node resource capacity expansion strategy
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=static;auto
	// +kubebuilder:default=static
	ResourceExpansionStrategy ResourceExpansionStrategyType `json:"resourceExpansionStrategyType,omitempty"`
	// StaticResourceExpansion is the static resource expansion strategy, user specify the ratio of the node resource capacity expansion
	// +optional
	StaticResourceExpansion *StaticResourceExpansion `json:"staticResourceExpansion,omitempty"`
	// AutoResourceExpansion is the auto resource expansion strategy, crane auto expansion the node resource
	// +optional
	AutoResourceExpansion *AutoResourceExpansion `json:"autoResourceExpansion,omitempty"`

	// TargetLoadThreshold is the target utilization percent
	TargetLoadThreshold *TargetLoadThreshold `json:"targetLoadThreshold,omitempty"`
}

type TargetLoadThreshold struct {
	// Percents is the node load threshold for each resource, such as 60, it is a percentage
	// +optional
	Percents map[v1.ResourceName]int64 `json:"percents,omitempty"`
}

type StaticResourceExpansion struct {
	// Ratios define the static expansion ratio for node resource, such as cpu/memory. 1.5 means expansion 1.5 for the node resource capacity
	// +optional
	Ratios map[v1.ResourceName]string `json:"ratios,omitempty"`
}

type AutoResourceExpansion struct {
}

type NodeResourcePolicyStatus struct {
	// Conditions of the NodeResourcePolicyStatus
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// MetricSyncPolicy is the metric data syncing policy, it is a map of metric and its sync period.
// each metric has its own sync period. such as max_avg_5m is syncing each 5m
/** todo: later use, it is a global configuration for annotator controller to decide what metric and how the metric is syncing to target node annotation,
  now we keep the original configuration file way, it is no need to keep it associated with each node, global is better,
  when we control by each node, it will be very complicated and make the traffic to datasource is very large and out of control
*/
type MetricSyncPolicy struct {
	// SyncPeriod is the metric syncing period, controller will sync the metric value to get the latest metric each period
	// +optional
	SyncPeriod map[string]metav1.Duration `json:"syncPeriod,omitempty"`

	// GenHotValueMetric used to gen hot value metric for node
	GenHotValueMetric *HotValueMetric `json:"genHotValueMetric,omitempty"`
}

// this is just a metric computed by controller, so it belongs to controller to watch this config
type HotValueMetric struct {
	// HotValueWeights used to compute a hot value score for the node
	// +optional
	HotValueWeights []HotValueWeight `json:"hotValueWeights,omitempty"`
}

type HotValueWeight struct {
	// TimeRange is the time duration to statistic the binding counts from now to now-TimeRange
	// +optional
	TimeRange metav1.Duration `json:"timeRange,omitempty"`
	// Count is the weight used to compute the hot value
	// +optional
	Count int `json:"count,omitempty"`
}

type CraneSchedulerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CraneSchedulerConfigurationSpec `json:"spec,omitempty"`
}

// ScheduleApplyScope is a global config, use a crd is not suitable, use configmap instead. so now it will be not used.
type CraneSchedulerConfigurationSpec struct {
	// ScheduleApplyScope
	// +optional
	ScheduleApplyScope ScheduleApplyScope `json:"scheduleApplyScope,omitempty"`
	// SchedulePolicy is the policy used to do scheduling by crane scheduler
	// +optional
	SchedulePolicy SchedulePolicy `json:"schedulePolicy,omitempty"`
}


type NamespaceSelector struct {
	// Namespaces is applied namespaces which means all newly pods in these namespaces will use newly predicates & priorities to do scheduling
	// if it there is no namespaces, then all newly pods will keep the original scheduling logic
	// +optional
	Namespaces map[string]bool `json:"namespaces,omitempty"`
}

type PodSelector struct {
	Namespace string `json:"namespace,omitempty"`
	Selector  metav1.LabelSelector `json:"selector,omitempty"`
}

type ScheduleApplyScope struct {
	// ClusterScope means the newly scheduling policy will applied to the whole cluster, all newly pods
	// default is false which means scheduler will keep the original logic to schedule pods
	// and when ClusterScope is true, Selectors is not valid, its priority is highest
	// +optional
	ClusterScope bool `json:"clusterScope,omitempty"`
	// NamespaceSelector means the newly scheduling policy will be applied to all the namespaces specified
	// +optional
	NamespaceScopeSelector *NamespaceSelector `json:"namespaceScopeSelector,omitempty"`
	// PodScopeSelector means the newly scheduling policy will be applied to all selected pods satisfied the selector
	// +optional
	PodScopeSelector *[]PodSelector `json:"podScopeSelector,omitempty"`
}

// todo: later use, now scheduler do not support crd, use the original configuration file way
type SchedulePolicy struct {
	Predicates map[string]PredicatePolicy `json:"predicates,omitempty"`
	Priorities map[string]PriorityPolicy  `json:"priorities,omitempty"`
}

type PredicatePolicy struct {
	// Name of the predicate policy. such as overLoadFit
	Name string `json:"name,omitempty"`
	// ScheduleRuleType is the predicate schedule rule type.
	// default is a hard-code way to do some computing. same as the original configured policy. This is the default
	// expression is a go template expression way, user can write the expression, then scheduler will evaluate dynamically.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=default;expression
	// +kubebuilder:default=default
	ScheduleRuleType       PredicateScheduleRuleType        `json:"scheduleRuleType,omitempty"`
	DefaultScheduleRule    *DefaultPredicateScheduleRule    `json:"defaultScheduleRule,omitempty"`
	ExpressionScheduleRule *ExpressionPredicateScheduleRule `json:"expressionScheduleRule,omitempty"`
}

type PriorityPolicy struct {
	// Name of the predicate policy. such as safeLoadBalanced
	Name string `json:"name,omitempty"`
	// ScheduleRuleType is the priority schedule rule type.
	// default is a hard-code way to do some computing. same as the original configured policy. This is the default
	// expression is a go template expression way, user can write the expression, then scheduler will evaluate dynamically.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=default;expression
	// +kubebuilder:default=default
	ScheduleRuleType       PriorityScheduleRuleType        `json:"scheduleRuleType,omitempty"`
	DefaultScheduleRule    *DefaultPriorityScheduleRule    `json:"defaultScheduleRule,omitempty"`
	ExpressionScheduleRule *ExpressionPriorityScheduleRule `json:"expressionScheduleRule,omitempty"`
}

type PredicateScheduleRuleType string

type PriorityScheduleRuleType string

const (
	PredicateScheduleRuleTypeDefault    = "default"
	PredicateScheduleRuleTypeExpression = "expression"

	PriorityScheduleRuleTypeDefault    = "default"
	PriorityScheduleRuleTypeExpression = "expression"
)

type DefaultPredicateScheduleRule struct {
}

type ExpressionPredicateScheduleRule struct {
	// Name is the name of the schedule rule
	Name string `json:"name,omitempty"`
	// Expression is the expression of the schedule rule used to do computing by evaluate an expression
	// for example:
	//  max_avg_5m > 0.6 && max_avg_1h > 0.6  in predicates
	Expression string `json:"expression,omitempty"`
}

type DefaultPriorityScheduleRule struct {
}

type ExpressionPriorityScheduleRule struct {
	// Name is the name of the schedule rule
	Name string `json:"name,omitempty"`
	// Expression is the expression of the schedule rule used to do computing by evalute an expression
	// for example:
	//  ((1 - max_avg_5m) * w1 + (1- max_avg_1h) * w2 + (1- max_avg_1d) * w3)/(w1+w2+w3) - node_hot_value in priorities
	Expression string `json:"expression,omitempty"`
}
