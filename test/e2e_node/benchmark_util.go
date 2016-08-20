// +build linux

/*
Copyright 2015 The Kubernetes Authors.

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

package e2e_node

import (
	"sort"

	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/perftype"
)

const (
	// TODO(coufon): be consistent with perf_util.go version (not exposed)
	currentTimeSeriesVersion = "v1"
	TimeSeriesTag            = "[Result:TimeSeries]"
	TimeSeriesEnd            = "[Finish:TimeSeries]"
)

type NodeTimeSeries struct {
	// value in OperationData is an array of timestamps
	OperationData map[string][]int64         `json:"op_data,omitempty"`
	ResourceData  map[string]*ResourceSeries `json:"resource_data,omitempty"`
	Labels        map[string]string          `json:"labels"`
	Version       string                     `json:"version"`
}

// logDensityTimeSeries logs the time series data of operation and resource usage
func logDensityTimeSeries(rc *ResourceCollector, create, watch map[string]unversioned.Time, testName string) {
	timeSeries := &NodeTimeSeries{
		Labels: map[string]string{
			"node": framework.TestContext.NodeName,
			"test": testName,
		},
		Version: currentTimeSeriesVersion,
	}
	// Attach operation time series.
	timeSeries.OperationData = map[string][]int64{
		"create":  getCumulatedPodTimeSeries(create),
		"running": getCumulatedPodTimeSeries(watch),
	}
	// Attach resource time series.
	timeSeries.ResourceData = rc.GetResourceTimeSeries()
	// Log time series with tags
	framework.Logf("%s %s\n%s", TimeSeriesTag, framework.PrettyPrintJSON(timeSeries), TimeSeriesEnd)
}

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

// getCumulatedPodTimeSeries gets the cumulative pod number time series.
func getCumulatedPodTimeSeries(timePerPod map[string]unversioned.Time) []int64 {
	timeSeries := make(int64arr, 0)
	for _, ts := range timePerPod {
		timeSeries = append(timeSeries, ts.Time.UnixNano())
	}
	// Sort all timestamps.
	sort.Sort(timeSeries)
	return timeSeries
}

// getLatencyPerfData returns perf data from latency
func getLatencyPerfData(latency framework.LatencyMetric, testName string) *perftype.PerfData {
	return &perftype.PerfData{
		Version: "v1",
		DataItems: []perftype.DataItem{
			{
				Data: map[string]float64{
					"Perc50": float64(latency.Perc50) / 1000000,
					"Perc90": float64(latency.Perc90) / 1000000,
					"Perc99": float64(latency.Perc99) / 1000000,
				},
				Unit: "ms",
				Labels: map[string]string{
					"datatype":    "latency",
					"latencytype": "test-e2e",
				},
			},
		},
		Labels: map[string]string{
			"node": framework.TestContext.NodeName,
			"test": testName,
		},
	}
}
