// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubeflow/xgboost-1/pkg/apis/xgboost/v1alpha1"
	scheme "github.com/kubeflow/xgboost-1/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// XGBoostJobsGetter has a method to return a XGBoostJobInterface.
// A group's client should implement this interface.
type XGBoostJobsGetter interface {
	XGBoostJobs(namespace string) XGBoostJobInterface
}

// XGBoostJobInterface has methods to work with XGBoostJob resources.
type XGBoostJobInterface interface {
	Create(*v1alpha1.XGBoostJob) (*v1alpha1.XGBoostJob, error)
	Update(*v1alpha1.XGBoostJob) (*v1alpha1.XGBoostJob, error)
	UpdateStatus(*v1alpha1.XGBoostJob) (*v1alpha1.XGBoostJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.XGBoostJob, error)
	List(opts v1.ListOptions) (*v1alpha1.XGBoostJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.XGBoostJob, err error)
	XGBoostJobExpansion
}

// xGBoostJobs implements XGBoostJobInterface
type xGBoostJobs struct {
	client rest.Interface
	ns     string
}

// newXGBoostJobs returns a XGBoostJobs
func newXGBoostJobs(c *KubeflowV1alpha1Client, namespace string) *xGBoostJobs {
	return &xGBoostJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the xGBoostJob, and returns the corresponding xGBoostJob object, and an error if there is any.
func (c *xGBoostJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.XGBoostJob, err error) {
	result = &v1alpha1.XGBoostJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("xgboostjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of XGBoostJobs that match those selectors.
func (c *xGBoostJobs) List(opts v1.ListOptions) (result *v1alpha1.XGBoostJobList, err error) {
	result = &v1alpha1.XGBoostJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("xgboostjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested xGBoostJobs.
func (c *xGBoostJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("xgboostjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a xGBoostJob and creates it.  Returns the server's representation of the xGBoostJob, and an error, if there is any.
func (c *xGBoostJobs) Create(xGBoostJob *v1alpha1.XGBoostJob) (result *v1alpha1.XGBoostJob, err error) {
	result = &v1alpha1.XGBoostJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("xgboostjobs").
		Body(xGBoostJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a xGBoostJob and updates it. Returns the server's representation of the xGBoostJob, and an error, if there is any.
func (c *xGBoostJobs) Update(xGBoostJob *v1alpha1.XGBoostJob) (result *v1alpha1.XGBoostJob, err error) {
	result = &v1alpha1.XGBoostJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("xgboostjobs").
		Name(xGBoostJob.Name).
		Body(xGBoostJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *xGBoostJobs) UpdateStatus(xGBoostJob *v1alpha1.XGBoostJob) (result *v1alpha1.XGBoostJob, err error) {
	result = &v1alpha1.XGBoostJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("xgboostjobs").
		Name(xGBoostJob.Name).
		SubResource("status").
		Body(xGBoostJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the xGBoostJob and deletes it. Returns an error if one occurs.
func (c *xGBoostJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("xgboostjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *xGBoostJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("xgboostjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched xGBoostJob.
func (c *xGBoostJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.XGBoostJob, err error) {
	result = &v1alpha1.XGBoostJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("xgboostjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
