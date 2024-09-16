//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllAtOnceRolloutStrategy) DeepCopyInto(out *AllAtOnceRolloutStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllAtOnceRolloutStrategy.
func (in *AllAtOnceRolloutStrategy) DeepCopy() *AllAtOnceRolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(AllAtOnceRolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualRolloutStrategy) DeepCopyInto(out *ManualRolloutStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualRolloutStrategy.
func (in *ManualRolloutStrategy) DeepCopy() *ManualRolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(ManualRolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatistics) DeepCopyInto(out *QueueStatistics) {
	*out = *in
	out.ApproximateBacklogAge = in.ApproximateBacklogAge
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatistics.
func (in *QueueStatistics) DeepCopy() *QueueStatistics {
	if in == nil {
		return nil
	}
	out := new(QueueStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStep) DeepCopyInto(out *RolloutStep) {
	*out = *in
	out.PauseDuration = in.PauseDuration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStep.
func (in *RolloutStep) DeepCopy() *RolloutStep {
	if in == nil {
		return nil
	}
	out := new(RolloutStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStrategy) DeepCopyInto(out *RolloutStrategy) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]RolloutStep, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStrategy.
func (in *RolloutStrategy) DeepCopy() *RolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(RolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalConnection) DeepCopyInto(out *TemporalConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalConnection.
func (in *TemporalConnection) DeepCopy() *TemporalConnection {
	if in == nil {
		return nil
	}
	out := new(TemporalConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalConnectionList) DeepCopyInto(out *TemporalConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemporalConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalConnectionList.
func (in *TemporalConnectionList) DeepCopy() *TemporalConnectionList {
	if in == nil {
		return nil
	}
	out := new(TemporalConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalConnectionSpec) DeepCopyInto(out *TemporalConnectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalConnectionSpec.
func (in *TemporalConnectionSpec) DeepCopy() *TemporalConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(TemporalConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalConnectionStatus) DeepCopyInto(out *TemporalConnectionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalConnectionStatus.
func (in *TemporalConnectionStatus) DeepCopy() *TemporalConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(TemporalConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalWorker) DeepCopyInto(out *TemporalWorker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalWorker.
func (in *TemporalWorker) DeepCopy() *TemporalWorker {
	if in == nil {
		return nil
	}
	out := new(TemporalWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalWorker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalWorkerList) DeepCopyInto(out *TemporalWorkerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemporalWorker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalWorkerList.
func (in *TemporalWorkerList) DeepCopy() *TemporalWorkerList {
	if in == nil {
		return nil
	}
	out := new(TemporalWorkerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemporalWorkerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalWorkerSpec) DeepCopyInto(out *TemporalWorkerSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.ProgressDeadlineSeconds != nil {
		in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
		*out = new(int32)
		**out = **in
	}
	in.RolloutStrategy.DeepCopyInto(&out.RolloutStrategy)
	out.WorkerOptions = in.WorkerOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalWorkerSpec.
func (in *TemporalWorkerSpec) DeepCopy() *TemporalWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(TemporalWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemporalWorkerStatus) DeepCopyInto(out *TemporalWorkerStatus) {
	*out = *in
	if in.TargetVersion != nil {
		in, out := &in.TargetVersion, &out.TargetVersion
		*out = new(VersionedDeployment)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(VersionedDeployment)
		(*in).DeepCopyInto(*out)
	}
	if in.DeprecatedVersions != nil {
		in, out := &in.DeprecatedVersions, &out.DeprecatedVersions
		*out = make([]*VersionedDeployment, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(VersionedDeployment)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.VersionConflictToken != nil {
		in, out := &in.VersionConflictToken, &out.VersionConflictToken
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemporalWorkerStatus.
func (in *TemporalWorkerStatus) DeepCopy() *TemporalWorkerStatus {
	if in == nil {
		return nil
	}
	out := new(TemporalWorkerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionedDeployment) DeepCopyInto(out *VersionedDeployment) {
	*out = *in
	if in.HealthySince != nil {
		in, out := &in.HealthySince, &out.HealthySince
		*out = (*in).DeepCopy()
	}
	if in.CompatibleBuildIDs != nil {
		in, out := &in.CompatibleBuildIDs, &out.CompatibleBuildIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RampPercentage != nil {
		in, out := &in.RampPercentage, &out.RampPercentage
		*out = new(uint8)
		**out = **in
	}
	if in.Statistics != nil {
		in, out := &in.Statistics, &out.Statistics
		*out = new(QueueStatistics)
		**out = **in
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionedDeployment.
func (in *VersionedDeployment) DeepCopy() *VersionedDeployment {
	if in == nil {
		return nil
	}
	out := new(VersionedDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerOptions) DeepCopyInto(out *WorkerOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerOptions.
func (in *WorkerOptions) DeepCopy() *WorkerOptions {
	if in == nil {
		return nil
	}
	out := new(WorkerOptions)
	in.DeepCopyInto(out)
	return out
}
