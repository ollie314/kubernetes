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

package predicates

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/labels"
)

// ExampleUtils is a https://blog.golang.org/examples styled unit test.
func ExampleFindLabelsInSet() {
	labelSubset := labels.Set{}
	labelSubset["label1"] = "value1"
	labelSubset["label2"] = "value2"
	// Lets make believe that these pods are on the cluster.
	// Utility functions will inspect their labels, filter them, and so on.
	nsPods := []*api.Pod{
		{
			ObjectMeta: api.ObjectMeta{
				Name:      "pod1",
				Namespace: "ns1",
				Labels: map[string]string{
					"label1": "wontSeeThis",
					"label2": "wontSeeThis",
					"label3": "will_see_this",
				},
			},
		}, // first pod which will be used via the utilities
		{
			ObjectMeta: api.ObjectMeta{
				Name:      "pod2",
				Namespace: "ns1",
			},
		},

		{}, // a third pod which will have no effect on anything.
	}
	fmt.Println(FindLabelsInSet([]string{"label1", "label2", "label3"}, nsPods[0].ObjectMeta.Labels)["label3"])
	AddUnsetLabelsToMap(labelSubset, []string{"label1", "label2", "label3"}, nsPods[0].ObjectMeta.Labels)
	fmt.Println(labelSubset)

	for _, pod := range FilterPodsByNamespace(nsPods, "ns1") {
		fmt.Print(pod.Name, ",")
	}
	// Output:
	// will_see_this
	// label1=value1,label2=value2,label3=will_see_this
	// pod1,pod2,
}
