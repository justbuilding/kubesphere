// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "istio.io/client-go/pkg/clientset/versioned/typed/rbac/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRbacV1alpha1 struct {
	*testing.Fake
}

func (c *FakeRbacV1alpha1) ClusterRbacConfigs() v1alpha1.ClusterRbacConfigInterface {
	return &FakeClusterRbacConfigs{c}
}

func (c *FakeRbacV1alpha1) RbacConfigs(namespace string) v1alpha1.RbacConfigInterface {
	return &FakeRbacConfigs{c, namespace}
}

func (c *FakeRbacV1alpha1) ServiceRoles(namespace string) v1alpha1.ServiceRoleInterface {
	return &FakeServiceRoles{c, namespace}
}

func (c *FakeRbacV1alpha1) ServiceRoleBindings(namespace string) v1alpha1.ServiceRoleBindingInterface {
	return &FakeServiceRoleBindings{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRbacV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}