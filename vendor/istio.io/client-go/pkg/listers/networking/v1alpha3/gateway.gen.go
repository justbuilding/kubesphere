// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GatewayLister helps list Gateways.
type GatewayLister interface {
	// List lists all Gateways in the indexer.
	List(selector labels.Selector) (ret []*v1alpha3.Gateway, err error)
	// Gateways returns an object that can list and get Gateways.
	Gateways(namespace string) GatewayNamespaceLister
	GatewayListerExpansion
}

// gatewayLister implements the GatewayLister interface.
type gatewayLister struct {
	indexer cache.Indexer
}

// NewGatewayLister returns a new GatewayLister.
func NewGatewayLister(indexer cache.Indexer) GatewayLister {
	return &gatewayLister{indexer: indexer}
}

// List lists all Gateways in the indexer.
func (s *gatewayLister) List(selector labels.Selector) (ret []*v1alpha3.Gateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.Gateway))
	})
	return ret, err
}

// Gateways returns an object that can list and get Gateways.
func (s *gatewayLister) Gateways(namespace string) GatewayNamespaceLister {
	return gatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GatewayNamespaceLister helps list and get Gateways.
type GatewayNamespaceLister interface {
	// List lists all Gateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha3.Gateway, err error)
	// Get retrieves the Gateway from the indexer for a given namespace and name.
	Get(name string) (*v1alpha3.Gateway, error)
	GatewayNamespaceListerExpansion
}

// gatewayNamespaceLister implements the GatewayNamespaceLister
// interface.
type gatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Gateways in the indexer for a given namespace.
func (s gatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.Gateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.Gateway))
	})
	return ret, err
}

// Get retrieves the Gateway from the indexer for a given namespace and name.
func (s gatewayNamespaceLister) Get(name string) (*v1alpha3.Gateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("gateway"), name)
	}
	return obj.(*v1alpha3.Gateway), nil
}
