// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/riotkit-org/backup-maker-operator/pkg/apis/riotkit/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScheduledBackupLister helps list ScheduledBackups.
// All objects returned here must be treated as read-only.
type ScheduledBackupLister interface {
	// List lists all ScheduledBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledBackup, err error)
	// ScheduledBackups returns an object that can list and get ScheduledBackups.
	ScheduledBackups(namespace string) ScheduledBackupNamespaceLister
	ScheduledBackupListerExpansion
}

// scheduledBackupLister implements the ScheduledBackupLister interface.
type scheduledBackupLister struct {
	indexer cache.Indexer
}

// NewScheduledBackupLister returns a new ScheduledBackupLister.
func NewScheduledBackupLister(indexer cache.Indexer) ScheduledBackupLister {
	return &scheduledBackupLister{indexer: indexer}
}

// List lists all ScheduledBackups in the indexer.
func (s *scheduledBackupLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledBackup))
	})
	return ret, err
}

// ScheduledBackups returns an object that can list and get ScheduledBackups.
func (s *scheduledBackupLister) ScheduledBackups(namespace string) ScheduledBackupNamespaceLister {
	return scheduledBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScheduledBackupNamespaceLister helps list and get ScheduledBackups.
// All objects returned here must be treated as read-only.
type ScheduledBackupNamespaceLister interface {
	// List lists all ScheduledBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledBackup, err error)
	// Get retrieves the ScheduledBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScheduledBackup, error)
	ScheduledBackupNamespaceListerExpansion
}

// scheduledBackupNamespaceLister implements the ScheduledBackupNamespaceLister
// interface.
type scheduledBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScheduledBackups in the indexer for a given namespace.
func (s scheduledBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledBackup))
	})
	return ret, err
}

// Get retrieves the ScheduledBackup from the indexer for a given namespace and name.
func (s scheduledBackupNamespaceLister) Get(name string) (*v1alpha1.ScheduledBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scheduledbackup"), name)
	}
	return obj.(*v1alpha1.ScheduledBackup), nil
}
