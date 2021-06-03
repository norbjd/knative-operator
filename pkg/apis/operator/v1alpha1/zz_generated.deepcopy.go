// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwssqsSourceConfiguration) DeepCopyInto(out *AwssqsSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwssqsSourceConfiguration.
func (in *AwssqsSourceConfiguration) DeepCopy() *AwssqsSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(AwssqsSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephSourceConfiguration) DeepCopyInto(out *CephSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephSourceConfiguration.
func (in *CephSourceConfiguration) DeepCopy() *CephSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(CephSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSpec) DeepCopyInto(out *CommonSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(ConfigMapData, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	in.Registry.DeepCopyInto(&out.Registry)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceRequirementsOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentOverride != nil {
		in, out := &in.DeploymentOverride, &out.DeploymentOverride
		*out = make([]DeploymentOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make([]Manifest, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalManifests != nil {
		in, out := &in.AdditionalManifests, &out.AdditionalManifests
		*out = make([]Manifest, len(*in))
		copy(*out, *in)
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailability)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSpec.
func (in *CommonSpec) DeepCopy() *CommonSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ConfigMapData) DeepCopyInto(out *ConfigMapData) {
	{
		in := &in
		*out = make(ConfigMapData, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapData.
func (in ConfigMapData) DeepCopy() ConfigMapData {
	if in == nil {
		return nil
	}
	out := new(ConfigMapData)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContourIngressConfiguration) DeepCopyInto(out *ContourIngressConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContourIngressConfiguration.
func (in *ContourIngressConfiguration) DeepCopy() *ContourIngressConfiguration {
	if in == nil {
		return nil
	}
	out := new(ContourIngressConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchdbSourceConfiguration) DeepCopyInto(out *CouchdbSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchdbSourceConfiguration.
func (in *CouchdbSourceConfiguration) DeepCopy() *CouchdbSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(CouchdbSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomCerts) DeepCopyInto(out *CustomCerts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomCerts.
func (in *CustomCerts) DeepCopy() *CustomCerts {
	if in == nil {
		return nil
	}
	out := new(CustomCerts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentOverride) DeepCopyInto(out *DeploymentOverride) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentOverride.
func (in *DeploymentOverride) DeepCopy() *DeploymentOverride {
	if in == nil {
		return nil
	}
	out := new(DeploymentOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubSourceConfiguration) DeepCopyInto(out *GithubSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubSourceConfiguration.
func (in *GithubSourceConfiguration) DeepCopy() *GithubSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(GithubSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitlabSourceConfiguration) DeepCopyInto(out *GitlabSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitlabSourceConfiguration.
func (in *GitlabSourceConfiguration) DeepCopy() *GitlabSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(GitlabSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailability) DeepCopyInto(out *HighAvailability) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailability.
func (in *HighAvailability) DeepCopy() *HighAvailability {
	if in == nil {
		return nil
	}
	out := new(HighAvailability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressConfigs) DeepCopyInto(out *IngressConfigs) {
	*out = *in
	in.Istio.DeepCopyInto(&out.Istio)
	out.Kourier = in.Kourier
	out.Contour = in.Contour
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressConfigs.
func (in *IngressConfigs) DeepCopy() *IngressConfigs {
	if in == nil {
		return nil
	}
	out := new(IngressConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioGatewayOverride) DeepCopyInto(out *IstioGatewayOverride) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioGatewayOverride.
func (in *IstioGatewayOverride) DeepCopy() *IstioGatewayOverride {
	if in == nil {
		return nil
	}
	out := new(IstioGatewayOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioIngressConfiguration) DeepCopyInto(out *IstioIngressConfiguration) {
	*out = *in
	if in.KnativeIngressGateway != nil {
		in, out := &in.KnativeIngressGateway, &out.KnativeIngressGateway
		*out = new(IstioGatewayOverride)
		(*in).DeepCopyInto(*out)
	}
	if in.KnativeLocalGateway != nil {
		in, out := &in.KnativeLocalGateway, &out.KnativeLocalGateway
		*out = new(IstioGatewayOverride)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioIngressConfiguration.
func (in *IstioIngressConfiguration) DeepCopy() *IstioIngressConfiguration {
	if in == nil {
		return nil
	}
	out := new(IstioIngressConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceConfiguration) DeepCopyInto(out *KafkaSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceConfiguration.
func (in *KafkaSourceConfiguration) DeepCopy() *KafkaSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeEventing) DeepCopyInto(out *KnativeEventing) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeEventing.
func (in *KnativeEventing) DeepCopy() *KnativeEventing {
	if in == nil {
		return nil
	}
	out := new(KnativeEventing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KnativeEventing) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeEventingList) DeepCopyInto(out *KnativeEventingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KnativeEventing, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeEventingList.
func (in *KnativeEventingList) DeepCopy() *KnativeEventingList {
	if in == nil {
		return nil
	}
	out := new(KnativeEventingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KnativeEventingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeEventingSpec) DeepCopyInto(out *KnativeEventingSpec) {
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(SourceConfigs)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeEventingSpec.
func (in *KnativeEventingSpec) DeepCopy() *KnativeEventingSpec {
	if in == nil {
		return nil
	}
	out := new(KnativeEventingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeEventingStatus) DeepCopyInto(out *KnativeEventingStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeEventingStatus.
func (in *KnativeEventingStatus) DeepCopy() *KnativeEventingStatus {
	if in == nil {
		return nil
	}
	out := new(KnativeEventingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServing) DeepCopyInto(out *KnativeServing) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServing.
func (in *KnativeServing) DeepCopy() *KnativeServing {
	if in == nil {
		return nil
	}
	out := new(KnativeServing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KnativeServing) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServingList) DeepCopyInto(out *KnativeServingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KnativeServing, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServingList.
func (in *KnativeServingList) DeepCopy() *KnativeServingList {
	if in == nil {
		return nil
	}
	out := new(KnativeServingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KnativeServingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServingSpec) DeepCopyInto(out *KnativeServingSpec) {
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	in.DeprecatedKnativeIngressGateway.DeepCopyInto(&out.DeprecatedKnativeIngressGateway)
	in.DeprecatedClusterLocalGateway.DeepCopyInto(&out.DeprecatedClusterLocalGateway)
	out.ControllerCustomCerts = in.ControllerCustomCerts
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(IngressConfigs)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServingSpec.
func (in *KnativeServingSpec) DeepCopy() *KnativeServingSpec {
	if in == nil {
		return nil
	}
	out := new(KnativeServingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServingStatus) DeepCopyInto(out *KnativeServingStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServingStatus.
func (in *KnativeServingStatus) DeepCopy() *KnativeServingStatus {
	if in == nil {
		return nil
	}
	out := new(KnativeServingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KourierIngressConfiguration) DeepCopyInto(out *KourierIngressConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KourierIngressConfiguration.
func (in *KourierIngressConfiguration) DeepCopy() *KourierIngressConfiguration {
	if in == nil {
		return nil
	}
	out := new(KourierIngressConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Manifest) DeepCopyInto(out *Manifest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Manifest.
func (in *Manifest) DeepCopy() *Manifest {
	if in == nil {
		return nil
	}
	out := new(Manifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatssSourceConfiguration) DeepCopyInto(out *NatssSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatssSourceConfiguration.
func (in *NatssSourceConfiguration) DeepCopy() *NatssSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(NatssSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSourceConfiguration) DeepCopyInto(out *PrometheusSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSourceConfiguration.
func (in *PrometheusSourceConfiguration) DeepCopy() *PrometheusSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(PrometheusSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqSourceConfiguration) DeepCopyInto(out *RabbitmqSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqSourceConfiguration.
func (in *RabbitmqSourceConfiguration) DeepCopy() *RabbitmqSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(RabbitmqSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSourceConfiguration) DeepCopyInto(out *RedisSourceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSourceConfiguration.
func (in *RedisSourceConfiguration) DeepCopy() *RedisSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(RedisSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirementsOverride) DeepCopyInto(out *ResourceRequirementsOverride) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirementsOverride.
func (in *ResourceRequirementsOverride) DeepCopy() *ResourceRequirementsOverride {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirementsOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceConfigs) DeepCopyInto(out *SourceConfigs) {
	*out = *in
	out.Awssqs = in.Awssqs
	out.Ceph = in.Ceph
	out.Couchdb = in.Couchdb
	out.Github = in.Github
	out.Gitlab = in.Gitlab
	out.Kafka = in.Kafka
	out.Natss = in.Natss
	out.Prometheus = in.Prometheus
	out.Rabbitmq = in.Rabbitmq
	out.Redis = in.Redis
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceConfigs.
func (in *SourceConfigs) DeepCopy() *SourceConfigs {
	if in == nil {
		return nil
	}
	out := new(SourceConfigs)
	in.DeepCopyInto(out)
	return out
}
