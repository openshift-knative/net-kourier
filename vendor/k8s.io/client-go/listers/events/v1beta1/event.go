/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "k8s.io/api/events/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// EventLister helps list Events.
// All objects returned here must be treated as read-only.
type EventLister interface {
	// List lists all Events in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Event, err error)
	// Events returns an object that can list and get Events.
	Events(namespace string) EventNamespaceLister
	EventListerExpansion
}

// eventLister implements the EventLister interface.
type eventLister struct {
	listers.ResourceIndexer[*v1beta1.Event]
}

// NewEventLister returns a new EventLister.
func NewEventLister(indexer cache.Indexer) EventLister {
	return &eventLister{listers.New[*v1beta1.Event](indexer, v1beta1.Resource("event"))}
}

// Events returns an object that can list and get Events.
func (s *eventLister) Events(namespace string) EventNamespaceLister {
	return eventNamespaceLister{listers.NewNamespaced[*v1beta1.Event](s.ResourceIndexer, namespace)}
}

// EventNamespaceLister helps list and get Events.
// All objects returned here must be treated as read-only.
type EventNamespaceLister interface {
	// List lists all Events in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Event, err error)
	// Get retrieves the Event from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Event, error)
	EventNamespaceListerExpansion
}

// eventNamespaceLister implements the EventNamespaceLister
// interface.
type eventNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.Event]
}
