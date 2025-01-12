/*
Copyright 2018 Replicated.

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
// Code generated by main. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/replicatedhq/gatekeeper/pkg/apis/policies/v1alpha2"
	scheme "github.com/replicatedhq/gatekeeper/pkg/client/gatekeeperclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AdmissionPoliciesGetter has a method to return a AdmissionPolicyInterface.
// A group's client should implement this interface.
type AdmissionPoliciesGetter interface {
	AdmissionPolicies(namespace string) AdmissionPolicyInterface
}

// AdmissionPolicyInterface has methods to work with AdmissionPolicy resources.
type AdmissionPolicyInterface interface {
	Create(*v1alpha2.AdmissionPolicy) (*v1alpha2.AdmissionPolicy, error)
	Update(*v1alpha2.AdmissionPolicy) (*v1alpha2.AdmissionPolicy, error)
	UpdateStatus(*v1alpha2.AdmissionPolicy) (*v1alpha2.AdmissionPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.AdmissionPolicy, error)
	List(opts v1.ListOptions) (*v1alpha2.AdmissionPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AdmissionPolicy, err error)
	AdmissionPolicyExpansion
}

// admissionPolicies implements AdmissionPolicyInterface
type admissionPolicies struct {
	client rest.Interface
	ns     string
}

// newAdmissionPolicies returns a AdmissionPolicies
func newAdmissionPolicies(c *PoliciesV1alpha2Client, namespace string) *admissionPolicies {
	return &admissionPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the admissionPolicy, and returns the corresponding admissionPolicy object, and an error if there is any.
func (c *admissionPolicies) Get(name string, options v1.GetOptions) (result *v1alpha2.AdmissionPolicy, err error) {
	result = &v1alpha2.AdmissionPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("admissionpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AdmissionPolicies that match those selectors.
func (c *admissionPolicies) List(opts v1.ListOptions) (result *v1alpha2.AdmissionPolicyList, err error) {
	result = &v1alpha2.AdmissionPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("admissionpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested admissionPolicies.
func (c *admissionPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("admissionpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a admissionPolicy and creates it.  Returns the server's representation of the admissionPolicy, and an error, if there is any.
func (c *admissionPolicies) Create(admissionPolicy *v1alpha2.AdmissionPolicy) (result *v1alpha2.AdmissionPolicy, err error) {
	result = &v1alpha2.AdmissionPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("admissionpolicies").
		Body(admissionPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a admissionPolicy and updates it. Returns the server's representation of the admissionPolicy, and an error, if there is any.
func (c *admissionPolicies) Update(admissionPolicy *v1alpha2.AdmissionPolicy) (result *v1alpha2.AdmissionPolicy, err error) {
	result = &v1alpha2.AdmissionPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("admissionpolicies").
		Name(admissionPolicy.Name).
		Body(admissionPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *admissionPolicies) UpdateStatus(admissionPolicy *v1alpha2.AdmissionPolicy) (result *v1alpha2.AdmissionPolicy, err error) {
	result = &v1alpha2.AdmissionPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("admissionpolicies").
		Name(admissionPolicy.Name).
		SubResource("status").
		Body(admissionPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the admissionPolicy and deletes it. Returns an error if one occurs.
func (c *admissionPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("admissionpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *admissionPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("admissionpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched admissionPolicy.
func (c *admissionPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AdmissionPolicy, err error) {
	result = &v1alpha2.AdmissionPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("admissionpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
