// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package federationin

import (
	"github.com/google/exposure-notifications-server/internal/metrics"
	"github.com/google/exposure-notifications-server/pkg/observability"
	"go.opencensus.io/stats/view"
)

func init() {
	observability.CollectViews([]*view.View{
		{
			Name:        metrics.MetricRoot + "pull_invalid_request_count",
			Description: "Total count of errors in pulling query IDs",
			Measure:     PullInvalidRequest,
			Aggregation: view.Sum(),
		},
		{
			Name:        metrics.MetricRoot + "pull_lock_contention_count",
			Description: "Total count of lock contention during pull operations",
			Measure:     PullLockContention,
			Aggregation: view.Sum(),
		},
		{
			Name:        metrics.MetricRoot + "pull_insertions_latest",
			Description: "Last value of exposure insertions",
			Measure:     PullInserts,
			Aggregation: view.LastValue(),
		},
		{
			Name:        metrics.MetricRoot + "pull_revisions_latest",
			Description: "Last value of exposure revisions",
			Measure:     PullRevisions,
			Aggregation: view.LastValue(),
		},
		{
			Name:        metrics.MetricRoot + "pull_droped_latest",
			Description: "Last value of exposure droped",
			Measure:     PullDropped,
			Aggregation: view.LastValue(),
		},
	}...)
}
