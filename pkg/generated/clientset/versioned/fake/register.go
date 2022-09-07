// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	analysisv1alpha1 "git.woa.com/crane/api/analysis/v1alpha1"
	autoscalingv1alpha1 "git.woa.com/crane/api/autoscaling/v1alpha1"
	ensurancev1alpha1 "git.woa.com/crane/api/ensurance/v1alpha1"
	predictionv1alpha1 "git.woa.com/crane/api/prediction/v1alpha1"
	schedulingv1alpha1 "git.woa.com/crane/api/scheduling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	analysisv1alpha1.AddToScheme,
	autoscalingv1alpha1.AddToScheme,
	ensurancev1alpha1.AddToScheme,
	predictionv1alpha1.AddToScheme,
	schedulingv1alpha1.AddToScheme,
}

var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
