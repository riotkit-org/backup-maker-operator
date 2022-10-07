// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/riotkit-org/backup-maker-operator/pkg/apis/riotkit/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRestoredBackups implements RestoredBackupInterface
type FakeRestoredBackups struct {
	Fake *FakeRiotkitV1alpha1
	ns   string
}

var restoredbackupsResource = schema.GroupVersionResource{Group: "riotkit.org", Version: "v1alpha1", Resource: "restoredbackups"}

var restoredbackupsKind = schema.GroupVersionKind{Group: "riotkit.org", Version: "v1alpha1", Kind: "RestoredBackup"}

// Get takes name of the restoredBackup, and returns the corresponding restoredBackup object, and an error if there is any.
func (c *FakeRestoredBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RestoredBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(restoredbackupsResource, c.ns, name), &v1alpha1.RestoredBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestoredBackup), err
}

// List takes label and field selectors, and returns the list of RestoredBackups that match those selectors.
func (c *FakeRestoredBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RestoredBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(restoredbackupsResource, restoredbackupsKind, c.ns, opts), &v1alpha1.RestoredBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RestoredBackupList{ListMeta: obj.(*v1alpha1.RestoredBackupList).ListMeta}
	for _, item := range obj.(*v1alpha1.RestoredBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested restoredBackups.
func (c *FakeRestoredBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(restoredbackupsResource, c.ns, opts))

}

// Create takes the representation of a restoredBackup and creates it.  Returns the server's representation of the restoredBackup, and an error, if there is any.
func (c *FakeRestoredBackups) Create(ctx context.Context, restoredBackup *v1alpha1.RestoredBackup, opts v1.CreateOptions) (result *v1alpha1.RestoredBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(restoredbackupsResource, c.ns, restoredBackup), &v1alpha1.RestoredBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestoredBackup), err
}

// Update takes the representation of a restoredBackup and updates it. Returns the server's representation of the restoredBackup, and an error, if there is any.
func (c *FakeRestoredBackups) Update(ctx context.Context, restoredBackup *v1alpha1.RestoredBackup, opts v1.UpdateOptions) (result *v1alpha1.RestoredBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(restoredbackupsResource, c.ns, restoredBackup), &v1alpha1.RestoredBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestoredBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRestoredBackups) UpdateStatus(ctx context.Context, restoredBackup *v1alpha1.RestoredBackup, opts v1.UpdateOptions) (*v1alpha1.RestoredBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(restoredbackupsResource, "status", c.ns, restoredBackup), &v1alpha1.RestoredBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestoredBackup), err
}

// Delete takes name of the restoredBackup and deletes it. Returns an error if one occurs.
func (c *FakeRestoredBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(restoredbackupsResource, c.ns, name), &v1alpha1.RestoredBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRestoredBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(restoredbackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RestoredBackupList{})
	return err
}

// Patch applies the patch and returns the patched restoredBackup.
func (c *FakeRestoredBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RestoredBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(restoredbackupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RestoredBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestoredBackup), err
}
