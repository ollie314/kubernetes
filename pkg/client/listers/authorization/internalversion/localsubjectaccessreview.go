/*
Copyright 2016 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package internalversion

import (
	"k8s.io/kubernetes/pkg/api/errors"
	authorization "k8s.io/kubernetes/pkg/apis/authorization"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// LocalSubjectAccessReviewLister helps list LocalSubjectAccessReviews.
type LocalSubjectAccessReviewLister interface {
	// List lists all LocalSubjectAccessReviews in the indexer.
	List(selector labels.Selector) (ret []*authorization.LocalSubjectAccessReview, err error)
	// LocalSubjectAccessReviews returns an object that can list and get LocalSubjectAccessReviews.
	LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewNamespaceLister
	LocalSubjectAccessReviewListerExpansion
}

// localSubjectAccessReviewLister implements the LocalSubjectAccessReviewLister interface.
type localSubjectAccessReviewLister struct {
	indexer cache.Indexer
}

// NewLocalSubjectAccessReviewLister returns a new LocalSubjectAccessReviewLister.
func NewLocalSubjectAccessReviewLister(indexer cache.Indexer) LocalSubjectAccessReviewLister {
	return &localSubjectAccessReviewLister{indexer: indexer}
}

// List lists all LocalSubjectAccessReviews in the indexer.
func (s *localSubjectAccessReviewLister) List(selector labels.Selector) (ret []*authorization.LocalSubjectAccessReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*authorization.LocalSubjectAccessReview))
	})
	return ret, err
}

// LocalSubjectAccessReviews returns an object that can list and get LocalSubjectAccessReviews.
func (s *localSubjectAccessReviewLister) LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewNamespaceLister {
	return localSubjectAccessReviewNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LocalSubjectAccessReviewNamespaceLister helps list and get LocalSubjectAccessReviews.
type LocalSubjectAccessReviewNamespaceLister interface {
	// List lists all LocalSubjectAccessReviews in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*authorization.LocalSubjectAccessReview, err error)
	// Get retrieves the LocalSubjectAccessReview from the indexer for a given namespace and name.
	Get(name string) (*authorization.LocalSubjectAccessReview, error)
	LocalSubjectAccessReviewNamespaceListerExpansion
}

// localSubjectAccessReviewNamespaceLister implements the LocalSubjectAccessReviewNamespaceLister
// interface.
type localSubjectAccessReviewNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LocalSubjectAccessReviews in the indexer for a given namespace.
func (s localSubjectAccessReviewNamespaceLister) List(selector labels.Selector) (ret []*authorization.LocalSubjectAccessReview, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*authorization.LocalSubjectAccessReview))
	})
	return ret, err
}

// Get retrieves the LocalSubjectAccessReview from the indexer for a given namespace and name.
func (s localSubjectAccessReviewNamespaceLister) Get(name string) (*authorization.LocalSubjectAccessReview, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(authorization.Resource("localsubjectaccessreview"), name)
	}
	return obj.(*authorization.LocalSubjectAccessReview), nil
}
