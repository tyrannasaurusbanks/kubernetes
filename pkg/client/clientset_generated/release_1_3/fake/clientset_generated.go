/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apimachinery/registered"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3"
	v1autoscaling "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3/typed/autoscaling/v1"
	fakev1autoscaling "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3/typed/autoscaling/v1/fake"
	v1batch "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3/typed/batch/v1"
	fakev1batch "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3/typed/batch/v1/fake"
	v1core "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3/typed/core/v1"
	fakev1core "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3/typed/core/v1/fake"
	v1beta1extensions "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3/typed/extensions/v1beta1"
	fakev1beta1extensions "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_3/typed/extensions/v1beta1/fake"
	"k8s.io/kubernetes/pkg/client/testing/core"
	"k8s.io/kubernetes/pkg/client/typed/discovery"
	fakediscovery "k8s.io/kubernetes/pkg/client/typed/discovery/fake"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/watch"
)

// Clientset returns a clientset that will respond with the provided objects
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := core.NewObjects(api.Scheme, api.Codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := core.Fake{}
	fakePtr.AddReactor("*", "*", core.ObjectReaction(o, registered.RESTMapper()))

	fakePtr.AddWatchReactor("*", core.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	core.Fake
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return &fakediscovery.FakeDiscovery{Fake: &c.Fake}
}

var _ clientset.Interface = &Clientset{}

// Core retrieves the CoreClient
func (c *Clientset) Core() v1core.CoreInterface {
	return &fakev1core.FakeCore{Fake: &c.Fake}
}

// Extensions retrieves the ExtensionsClient
func (c *Clientset) Extensions() v1beta1extensions.ExtensionsInterface {
	return &fakev1beta1extensions.FakeExtensions{Fake: &c.Fake}
}

// Autoscaling retrieves the AutoscalingClient
func (c *Clientset) Autoscaling() v1autoscaling.AutoscalingInterface {
	return &fakev1autoscaling.FakeAutoscaling{Fake: &c.Fake}
}

// Batch retrieves the BatchClient
func (c *Clientset) Batch() v1batch.BatchInterface {
	return &fakev1batch.FakeBatch{Fake: &c.Fake}
}
