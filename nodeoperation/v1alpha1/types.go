package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metatypes "k8s.io/apimachinery/pkg/types"
)

const (
	NodeOffline NodeOperationType = "NodeOffline"
	NodeRecycle NodeOperationType = "NodeRecycle"
)

// These are the valid phases of node.
const (
	// NodeReady means the node is working, status is ready.
	NodeReady NodePhase = "Ready"
	// NodeTaint means the node has been tainted.
	NodeTaint NodePhase = "Taint"
	// NodeBlocking means the pod on this node, maybe pending in descheduler phase.
	NodeBlocking NodePhase = "Blocking"
	// NodeCanBeOperating means do operation is safe.
	NodeCanBeOperating NodePhase = "CanBeOperating"
	// NodeOfflinling means the node is removing from the cluster.
	NodeOfflining NodePhase = "Offlining"
	// NodeOfflined means the node is already removed from the cluster.
	NodeOfflined NodePhase = "Offlined"
	// NodeRecycle means the node is removed from the cluster and be recycled.
	NodeRecycled NodePhase = "Recycled"
	// NodeConflict means the node is in other nor current work queue.
	NodeConflict NodePhase = "Conflict"
)

const (
	// NodeAutoOperationMode means no need admin double check, operation will do automaticly.
	NodeAutoOperationMode OperationMode = "Auto"
	// NodeManualOperationMode means node operation need admin double check, after admin checked ,the mode will change to auto.
	NodeManualOperationMode OperationMode = "Manual"
)

const (
	// OperationPending means the operation is pending b/c not in work time range.
	OperationPending OperationStatus = "Pending"
	// OperationAnalysising means analysising the nodes.
	OperationAnalysising OperationStatus = "Analysising"
	// OperationPartialRunning means not all node is doing operation.
	OperationPartialRunning OperationStatus = "PartialRunning"
	// OperationRunning means all node is doing Operation.
	OperationRunning OperationStatus = "Running"
	// OperationFailed means the node operation is failed.
	OperationFailed OperationStatus = "Failed"
	// OperationSuccess means the node operation is success.
	OperationSuccess OperationStatus = "Success"
)

type NodeOperationType string

type NodePhase string

type AnalysisResultType string

type OperationMode string

type OperationStatus string

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster,shortName=nop
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NodeOperation is the custome resource defination, define what operation will apply to kubernetes node.
type NodeOperation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NodeOperationSpec   `json:"spec,omitempty"`
	Status            NodeOperationStatus `json:"status,omitempty"`
}

type NodeOperationSpec struct {
	// NodeOperation means a series of operations for specific kubernetes node.
	// +optional
	// +kubebuilder:validation:MinItems=1
	NodeOperation []NodeOperationType `json:"nodeOperation,omitempty"`
	// Nodes means set of nodes which will be operated.
	// +optional
	NodeNames []string `json:"nodeNames,omitempty"`
	// NodeSelector provides kubernetes origin node selector ,which will select set of nodes which will be operated.
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// OperableTime set the time range of node operation.
	// +optional
	OperableTime OperableTime `json:"operableTime,omitempty"`
	// OperationMode can be Auto or Manual.
	// +optional
	// +kubebuilder:validation:Eunm=Auto,Manual
	// +kubebuilder:default=Manual
	OperationMode OperationMode `json:"operationMode,omitempty"`
	// Notification specify what and how you need notify for your business user.
	Notification Notification `json:"notification,omitempty"`
}

type NodeOperationStatus struct {
	// OperationStatus means the current status of node operation.
	OperationStatus OperationStatus `json:"operationStatus,omitempty"`
	// NodesStatus is a map store erery node status information.
	NodesStatus map[string]NodeStatus `json:"nodesStatus,omitempty"`
}

// NodeStatus means before and after operation start , the status of the node.
type NodeStatus struct {
	// Name means the node name.
	Name metatypes.NodeName `json:"name,omitempty"`
	// Phase means the phase of the node when do node operation.
	Phase NodePhase `json:"phase,omitempty"`
	// PodsStatus means the set of pods status in the node.
	PodsStatus map[string]PodStatus `json:"podsStatus,omitempty"`
	// Store labels of node.
	Labels map[string]string `json:"labels,omitempty"`
}

type PodStatus struct {
	// Phase is a label for the condition of a pod at the current time.
	Phase corev1.PodPhase `json:"phase,omitempty"`
	// OwnerReference contains enough information to let you identify an owning
	// object. An owning object must be in the same namespace as the dependent, or
	// be cluster-scoped, so there is no namespace field.
	OwnerReference metav1.OwnerReference `json:"ownerReference,omitempty"`
	// AnalysisEvent store the result event of analysis job, example deschedulable analysis.
	AnalysisEvent AnalysisEvent `json:"analysisEvent,omitempty"`
	// Store labels of pod.
	Labels map[string]string `json:"labels,omitempty"`
	// Store annotations of pod.
	Annotations map[string]string `json:"annotations,omitempty"`
}

type AnalysisEvent struct {
	// AnalysisMessage store information after anaylsis.
	// +kubebuilder:validation:Type=string
	AnalysisMessage string `json:"analysisMessage,omitempty"`
	// Event is from kubernetes core api, provide useful information.
	Event Event `json:"event,omitempty"`
}

// Event from kubernetes
type Event struct {
	// This should be a short, machine understandable string that gives the reason
	// for the transition into the object's current status.
	// TODO: provide exact specification for format.
	// +optional
	Reason string `json:"reason,omitempty"`

	// A human-readable description of the status of this operation.
	// TODO: decide on maximum length.
	// +optional
	Message string `json:"message,omitempty"`

	// The time at which the most recent occurrence of this event was recorded.
	// +optional
	LastTimestamp metav1.Time `json:"lastTimestamp,omitempty"`
}

// OperableTime Operable time range
type OperableTime struct {
	// +optional
	StartYear *int `json:"startYear,omitempty"`
	// +optional
	EndYear *int `json:"endYear,omitempty"`
	// +optional
	StartMonth *int `json:"startMonth,omitempty"`
	// +optional
	EndMonth *int `json:"endMonth,omitempty"`
	// +optional
	StartDay *int `json:"startDay,omitempty"`
	// +optional
	EndDay *int `json:"endDay,omitempty"`
	// +optional
	StartHour *int `json:"startHour,omitempty"`
	// +optional
	EndHour *int `json:"endHour,omitempty"`
	// +optional
	StartMinute *int `json:"startMinute,omitempty"`
	// +optional
	EndMinute *int `json:"endMinute,omitempty"`
	// +optional
	StartSecond *int `json:"startSecond,omitempty"`
	// +optional
	EndSecond *int `json:"endSecond,omitempty"`
}

type Notification struct {
	// +optional
	Message string `json:"message,omitempty"`
	// Level can be Info, Warning
	// +optional
	// +kubebuilder:validation:Eunm=Info,Warning
	// +kubebuilder:default=Info
	Level string `json:"level,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NodeOperationList contains a list of NodeOperation
type NodeOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodeOperation `json:"items"`
}
