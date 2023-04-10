//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *API) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIList) DeepCopyInto(out *APIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]API, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIList.
func (in *APIList) DeepCopy() *APIList {
	if in == nil {
		return nil
	}
	out := new(APIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicy) DeepCopyInto(out *APIPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicy.
func (in *APIPolicy) DeepCopy() *APIPolicy {
	if in == nil {
		return nil
	}
	out := new(APIPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicyList) DeepCopyInto(out *APIPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicyList.
func (in *APIPolicyList) DeepCopy() *APIPolicyList {
	if in == nil {
		return nil
	}
	out := new(APIPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicySpec) DeepCopyInto(out *APIPolicySpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(PolicySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(PolicySpec)
		(*in).DeepCopyInto(*out)
	}
	in.TargetRef.DeepCopyInto(&out.TargetRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicySpec.
func (in *APIPolicySpec) DeepCopy() *APIPolicySpec {
	if in == nil {
		return nil
	}
	out := new(APIPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicyStatus) DeepCopyInto(out *APIPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicyStatus.
func (in *APIPolicyStatus) DeepCopy() *APIPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(APIPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRateLimitPolicy) DeepCopyInto(out *APIRateLimitPolicy) {
	*out = *in
	out.RateLimit = in.RateLimit
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRateLimitPolicy.
func (in *APIRateLimitPolicy) DeepCopy() *APIRateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(APIRateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	if in.Production != nil {
		in, out := &in.Production, &out.Production
		*out = make([]EnvConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sandbox != nil {
		in, out := &in.Sandbox, &out.Sandbox
		*out = make([]EnvConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStatus) DeepCopyInto(out *APIStatus) {
	*out = *in
	if in.TransitionTime != nil {
		in, out := &in.TransitionTime, &out.TransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStatus.
func (in *APIStatus) DeepCopy() *APIStatus {
	if in == nil {
		return nil
	}
	out := new(APIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
	out.JWT = in.JWT
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthSpec) DeepCopyInto(out *AuthSpec) {
	*out = *in
	in.ExternalService.DeepCopyInto(&out.ExternalService)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthSpec.
func (in *AuthSpec) DeepCopy() *AuthSpec {
	if in == nil {
		return nil
	}
	out := new(AuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Authentication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationList) DeepCopyInto(out *AuthenticationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Authentication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationList.
func (in *AuthenticationList) DeepCopy() *AuthenticationList {
	if in == nil {
		return nil
	}
	out := new(AuthenticationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthenticationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
	in.Default.DeepCopyInto(&out.Default)
	in.Override.DeepCopyInto(&out.Override)
	in.TargetRef.DeepCopyInto(&out.TargetRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationStatus) DeepCopyInto(out *AuthenticationStatus) {
	*out = *in
	if in.TransitionTime != nil {
		in, out := &in.TransitionTime, &out.TransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationStatus.
func (in *AuthenticationStatus) DeepCopy() *AuthenticationStatus {
	if in == nil {
		return nil
	}
	out := new(AuthenticationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigs) DeepCopyInto(out *BackendConfigs) {
	*out = *in
	in.TLS.DeepCopyInto(&out.TLS)
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = make([]SecurityConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigs.
func (in *BackendConfigs) DeepCopy() *BackendConfigs {
	if in == nil {
		return nil
	}
	out := new(BackendConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendList) DeepCopyInto(out *BackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendList.
func (in *BackendList) DeepCopy() *BackendList {
	if in == nil {
		return nil
	}
	out := new(BackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in BackendMapping) DeepCopyInto(out *BackendMapping) {
	{
		in := &in
		*out = make(BackendMapping, len(*in))
		for key, val := range *in {
			var outVal *ResolvedBackend
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ResolvedBackend)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendMapping.
func (in BackendMapping) DeepCopy() BackendMapping {
	if in == nil {
		return nil
	}
	out := new(BackendMapping)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendReference) DeepCopyInto(out *BackendReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendReference.
func (in *BackendReference) DeepCopy() *BackendReference {
	if in == nil {
		return nil
	}
	out := new(BackendReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendSpec) DeepCopyInto(out *BackendSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]Service, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = make([]SecurityConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendSpec.
func (in *BackendSpec) DeepCopy() *BackendSpec {
	if in == nil {
		return nil
	}
	out := new(BackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendStatus) DeepCopyInto(out *BackendStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendStatus.
func (in *BackendStatus) DeepCopy() *BackendStatus {
	if in == nil {
		return nil
	}
	out := new(BackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicSecurityConfig) DeepCopyInto(out *BasicSecurityConfig) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicSecurityConfig.
func (in *BasicSecurityConfig) DeepCopy() *BasicSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(BasicSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRateLimitPolicy) DeepCopyInto(out *CustomRateLimitPolicy) {
	*out = *in
	out.RateLimit = in.RateLimit
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRateLimitPolicy.
func (in *CustomRateLimitPolicy) DeepCopy() *CustomRateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(CustomRateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvConfig) DeepCopyInto(out *EnvConfig) {
	*out = *in
	if in.HTTPRouteRefs != nil {
		in, out := &in.HTTPRouteRefs, &out.HTTPRouteRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvConfig.
func (in *EnvConfig) DeepCopy() *EnvConfig {
	if in == nil {
		return nil
	}
	out := new(EnvConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtAuthService) DeepCopyInto(out *ExtAuthService) {
	*out = *in
	out.ServiceRef = in.ServiceRef
	if in.AuthTypes != nil {
		in, out := &in.AuthTypes, &out.AuthTypes
		*out = make([]Auth, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtAuthService.
func (in *ExtAuthService) DeepCopy() *ExtAuthService {
	if in == nil {
		return nil
	}
	out := new(ExtAuthService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPQuery) DeepCopyInto(out *HTTPQuery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPQuery.
func (in *HTTPQuery) DeepCopy() *HTTPQuery {
	if in == nil {
		return nil
	}
	out := new(HTTPQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorConfig) DeepCopyInto(out *InterceptorConfig) {
	*out = *in
	out.BackendRef = in.BackendRef
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]InterceptorInclusion, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorConfig.
func (in *InterceptorConfig) DeepCopy() *InterceptorConfig {
	if in == nil {
		return nil
	}
	out := new(InterceptorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JWTAuth) DeepCopyInto(out *JWTAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTAuth.
func (in *JWTAuth) DeepCopy() *JWTAuth {
	if in == nil {
		return nil
	}
	out := new(JWTAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	in.RequestQueryModifier.DeepCopyInto(&out.RequestQueryModifier)
	if in.RequestInterceptor != nil {
		in, out := &in.RequestInterceptor, &out.RequestInterceptor
		*out = new(InterceptorConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ResponseInterceptor != nil {
		in, out := &in.ResponseInterceptor, &out.ResponseInterceptor
		*out = new(InterceptorConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimit) DeepCopyInto(out *RateLimit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimit.
func (in *RateLimit) DeepCopy() *RateLimit {
	if in == nil {
		return nil
	}
	out := new(RateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitAPIPolicy) DeepCopyInto(out *RateLimitAPIPolicy) {
	*out = *in
	out.API = in.API
	out.Custom = in.Custom
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitAPIPolicy.
func (in *RateLimitAPIPolicy) DeepCopy() *RateLimitAPIPolicy {
	if in == nil {
		return nil
	}
	out := new(RateLimitAPIPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicy) DeepCopyInto(out *RateLimitPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicy.
func (in *RateLimitPolicy) DeepCopy() *RateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyList) DeepCopyInto(out *RateLimitPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimitPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyList.
func (in *RateLimitPolicyList) DeepCopy() *RateLimitPolicyList {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicySpec) DeepCopyInto(out *RateLimitPolicySpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(RateLimitAPIPolicy)
		**out = **in
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(RateLimitAPIPolicy)
		**out = **in
	}
	in.TargetRef.DeepCopyInto(&out.TargetRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicySpec.
func (in *RateLimitPolicySpec) DeepCopy() *RateLimitPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyStatus) DeepCopyInto(out *RateLimitPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyStatus.
func (in *RateLimitPolicyStatus) DeepCopy() *RateLimitPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefConfig) DeepCopyInto(out *RefConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefConfig.
func (in *RefConfig) DeepCopy() *RefConfig {
	if in == nil {
		return nil
	}
	out := new(RefConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestQueryModifier) DeepCopyInto(out *RequestQueryModifier) {
	*out = *in
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make([]HTTPQuery, len(*in))
		copy(*out, *in)
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestQueryModifier.
func (in *RequestQueryModifier) DeepCopy() *RequestQueryModifier {
	if in == nil {
		return nil
	}
	out := new(RequestQueryModifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedBackend) DeepCopyInto(out *ResolvedBackend) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]Service, len(*in))
		copy(*out, *in)
	}
	in.TLS.DeepCopyInto(&out.TLS)
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = make([]ResolvedSecurityConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedBackend.
func (in *ResolvedBackend) DeepCopy() *ResolvedBackend {
	if in == nil {
		return nil
	}
	out := new(ResolvedBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedBasicSecurityConfig) DeepCopyInto(out *ResolvedBasicSecurityConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedBasicSecurityConfig.
func (in *ResolvedBasicSecurityConfig) DeepCopy() *ResolvedBasicSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(ResolvedBasicSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedSecurityConfig) DeepCopyInto(out *ResolvedSecurityConfig) {
	*out = *in
	out.Basic = in.Basic
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedSecurityConfig.
func (in *ResolvedSecurityConfig) DeepCopy() *ResolvedSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(ResolvedSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedTLSConfig) DeepCopyInto(out *ResolvedTLSConfig) {
	*out = *in
	if in.AllowedSANs != nil {
		in, out := &in.AllowedSANs, &out.AllowedSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedTLSConfig.
func (in *ResolvedTLSConfig) DeepCopy() *ResolvedTLSConfig {
	if in == nil {
		return nil
	}
	out := new(ResolvedTLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scope) DeepCopyInto(out *Scope) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scope.
func (in *Scope) DeepCopy() *Scope {
	if in == nil {
		return nil
	}
	out := new(Scope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Scope) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeList) DeepCopyInto(out *ScopeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Scope, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeList.
func (in *ScopeList) DeepCopy() *ScopeList {
	if in == nil {
		return nil
	}
	out := new(ScopeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScopeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeSpec) DeepCopyInto(out *ScopeSpec) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeSpec.
func (in *ScopeSpec) DeepCopy() *ScopeSpec {
	if in == nil {
		return nil
	}
	out := new(ScopeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeStatus) DeepCopyInto(out *ScopeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeStatus.
func (in *ScopeStatus) DeepCopy() *ScopeStatus {
	if in == nil {
		return nil
	}
	out := new(ScopeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfig) DeepCopyInto(out *SecurityConfig) {
	*out = *in
	out.Basic = in.Basic
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfig.
func (in *SecurityConfig) DeepCopy() *SecurityConfig {
	if in == nil {
		return nil
	}
	out := new(SecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRef) DeepCopyInto(out *ServiceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRef.
func (in *ServiceRef) DeepCopy() *ServiceRef {
	if in == nil {
		return nil
	}
	out := new(ServiceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(RefConfig)
		**out = **in
	}
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(RefConfig)
		**out = **in
	}
	if in.AllowedSANs != nil {
		in, out := &in.AllowedSANs, &out.AllowedSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}
