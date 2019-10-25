// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/apis/mcm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorks implements WorkInterface
type FakeWorks struct {
	Fake *FakeMcmV1alpha1
	ns   string
}

var worksResource = schema.GroupVersionResource{Group: "mcm.ibm.com", Version: "v1alpha1", Resource: "works"}

var worksKind = schema.GroupVersionKind{Group: "mcm.ibm.com", Version: "v1alpha1", Kind: "Work"}

// Get takes name of the work, and returns the corresponding work object, and an error if there is any.
func (c *FakeWorks) Get(name string, options v1.GetOptions) (result *v1alpha1.Work, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(worksResource, c.ns, name), &v1alpha1.Work{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Work), err
}

// List takes label and field selectors, and returns the list of Works that match those selectors.
func (c *FakeWorks) List(opts v1.ListOptions) (result *v1alpha1.WorkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(worksResource, worksKind, c.ns, opts), &v1alpha1.WorkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkList{ListMeta: obj.(*v1alpha1.WorkList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested works.
func (c *FakeWorks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(worksResource, c.ns, opts))

}

// Create takes the representation of a work and creates it.  Returns the server's representation of the work, and an error, if there is any.
func (c *FakeWorks) Create(work *v1alpha1.Work) (result *v1alpha1.Work, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(worksResource, c.ns, work), &v1alpha1.Work{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Work), err
}

// Update takes the representation of a work and updates it. Returns the server's representation of the work, and an error, if there is any.
func (c *FakeWorks) Update(work *v1alpha1.Work) (result *v1alpha1.Work, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(worksResource, c.ns, work), &v1alpha1.Work{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Work), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorks) UpdateStatus(work *v1alpha1.Work) (*v1alpha1.Work, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(worksResource, "status", c.ns, work), &v1alpha1.Work{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Work), err
}

// Delete takes name of the work and deletes it. Returns an error if one occurs.
func (c *FakeWorks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(worksResource, c.ns, name), &v1alpha1.Work{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(worksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkList{})
	return err
}

// Patch applies the patch and returns the patched work.
func (c *FakeWorks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Work, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(worksResource, c.ns, name, pt, data, subresources...), &v1alpha1.Work{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Work), err
}
