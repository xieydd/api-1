//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalysisEvent) DeepCopyInto(out *AnalysisEvent) {
	*out = *in
	in.Event.DeepCopyInto(&out.Event)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalysisEvent.
func (in *AnalysisEvent) DeepCopy() *AnalysisEvent {
	if in == nil {
		return nil
	}
	out := new(AnalysisEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeOperation) DeepCopyInto(out *NodeOperation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeOperation.
func (in *NodeOperation) DeepCopy() *NodeOperation {
	if in == nil {
		return nil
	}
	out := new(NodeOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeOperation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeOperationList) DeepCopyInto(out *NodeOperationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeOperation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeOperationList.
func (in *NodeOperationList) DeepCopy() *NodeOperationList {
	if in == nil {
		return nil
	}
	out := new(NodeOperationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeOperationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeOperationSpec) DeepCopyInto(out *NodeOperationSpec) {
	*out = *in
	if in.NodeOperation != nil {
		in, out := &in.NodeOperation, &out.NodeOperation
		*out = make([]NodeOperationType, len(*in))
		copy(*out, *in)
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.OperableTime.DeepCopyInto(&out.OperableTime)
	out.Notification = in.Notification
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeOperationSpec.
func (in *NodeOperationSpec) DeepCopy() *NodeOperationSpec {
	if in == nil {
		return nil
	}
	out := new(NodeOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeOperationStatus) DeepCopyInto(out *NodeOperationStatus) {
	*out = *in
	if in.NodesStatus != nil {
		in, out := &in.NodesStatus, &out.NodesStatus
		*out = make(map[string]NodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeOperationStatus.
func (in *NodeOperationStatus) DeepCopy() *NodeOperationStatus {
	if in == nil {
		return nil
	}
	out := new(NodeOperationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	if in.PodsStatus != nil {
		in, out := &in.PodsStatus, &out.PodsStatus
		*out = make(map[string]PodStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notification) DeepCopyInto(out *Notification) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notification.
func (in *Notification) DeepCopy() *Notification {
	if in == nil {
		return nil
	}
	out := new(Notification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperableTime) DeepCopyInto(out *OperableTime) {
	*out = *in
	if in.StartYear != nil {
		in, out := &in.StartYear, &out.StartYear
		*out = new(int)
		**out = **in
	}
	if in.EndYear != nil {
		in, out := &in.EndYear, &out.EndYear
		*out = new(int)
		**out = **in
	}
	if in.StartMonth != nil {
		in, out := &in.StartMonth, &out.StartMonth
		*out = new(int)
		**out = **in
	}
	if in.EndMonth != nil {
		in, out := &in.EndMonth, &out.EndMonth
		*out = new(int)
		**out = **in
	}
	if in.StartDay != nil {
		in, out := &in.StartDay, &out.StartDay
		*out = new(int)
		**out = **in
	}
	if in.EndDay != nil {
		in, out := &in.EndDay, &out.EndDay
		*out = new(int)
		**out = **in
	}
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(int)
		**out = **in
	}
	if in.EndHour != nil {
		in, out := &in.EndHour, &out.EndHour
		*out = new(int)
		**out = **in
	}
	if in.StartMinute != nil {
		in, out := &in.StartMinute, &out.StartMinute
		*out = new(int)
		**out = **in
	}
	if in.EndMinute != nil {
		in, out := &in.EndMinute, &out.EndMinute
		*out = new(int)
		**out = **in
	}
	if in.StartSecond != nil {
		in, out := &in.StartSecond, &out.StartSecond
		*out = new(int)
		**out = **in
	}
	if in.EndSecond != nil {
		in, out := &in.EndSecond, &out.EndSecond
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperableTime.
func (in *OperableTime) DeepCopy() *OperableTime {
	if in == nil {
		return nil
	}
	out := new(OperableTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
	in.OwnerReference.DeepCopyInto(&out.OwnerReference)
	in.AnalysisEvent.DeepCopyInto(&out.AnalysisEvent)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}
