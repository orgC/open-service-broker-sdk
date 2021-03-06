/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	broker "github.com/openshift/open-service-broker-sdk/pkg/apis/broker"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceInstances implements ServiceInstanceInterface
type FakeServiceInstances struct {
	Fake *FakeBroker
	ns   string
}

var serviceinstancesResource = schema.GroupVersionResource{Group: "sdkbroker.broker.k8s.io", Version: "", Resource: "serviceinstances"}

func (c *FakeServiceInstances) Create(serviceInstance *broker.ServiceInstance) (result *broker.ServiceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceinstancesResource, c.ns, serviceInstance), &broker.ServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*broker.ServiceInstance), err
}

func (c *FakeServiceInstances) Update(serviceInstance *broker.ServiceInstance) (result *broker.ServiceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceinstancesResource, c.ns, serviceInstance), &broker.ServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*broker.ServiceInstance), err
}

func (c *FakeServiceInstances) UpdateStatus(serviceInstance *broker.ServiceInstance) (*broker.ServiceInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceinstancesResource, "status", c.ns, serviceInstance), &broker.ServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*broker.ServiceInstance), err
}

func (c *FakeServiceInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceinstancesResource, c.ns, name), &broker.ServiceInstance{})

	return err
}

func (c *FakeServiceInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &broker.ServiceInstanceList{})
	return err
}

func (c *FakeServiceInstances) Get(name string, options v1.GetOptions) (result *broker.ServiceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceinstancesResource, c.ns, name), &broker.ServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*broker.ServiceInstance), err
}

func (c *FakeServiceInstances) List(opts v1.ListOptions) (result *broker.ServiceInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceinstancesResource, c.ns, opts), &broker.ServiceInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &broker.ServiceInstanceList{}
	for _, item := range obj.(*broker.ServiceInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceInstances.
func (c *FakeServiceInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceinstancesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched serviceInstance.
func (c *FakeServiceInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *broker.ServiceInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceinstancesResource, c.ns, name, data, subresources...), &broker.ServiceInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*broker.ServiceInstance), err
}
