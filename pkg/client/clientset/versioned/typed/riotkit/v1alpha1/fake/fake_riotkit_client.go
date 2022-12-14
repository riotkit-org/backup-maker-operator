// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/riotkit-org/backup-maker-controller/pkg/client/clientset/versioned/typed/riotkit/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRiotkitV1alpha1 struct {
	*testing.Fake
}

func (c *FakeRiotkitV1alpha1) ClusterBackupProcedureTemplates(namespace string) v1alpha1.ClusterBackupProcedureTemplateInterface {
	return &FakeClusterBackupProcedureTemplates{c, namespace}
}

func (c *FakeRiotkitV1alpha1) RequestedBackupActions(namespace string) v1alpha1.RequestedBackupActionInterface {
	return &FakeRequestedBackupActions{c, namespace}
}

func (c *FakeRiotkitV1alpha1) ScheduledBackups(namespace string) v1alpha1.ScheduledBackupInterface {
	return &FakeScheduledBackups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRiotkitV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
