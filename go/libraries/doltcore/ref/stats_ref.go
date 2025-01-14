// Copyright 2023 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ref

// StatsRefName is a dummy name, and there cannot be more than one stats ref.
const StatsRefName = "stats"

type StatsRef struct {
	stats string
}

var _ DoltRef = StatsRef{}

// NewStatsRef creates a reference to a statses list. There cannot be more than one statsRef.
func NewStatsRef() StatsRef {
	return StatsRef{StatsRefName}
}

// GetType will return StatsRefType
func (br StatsRef) GetType() RefType {
	return StatsRefType
}

// GetPath returns the name of the tag
func (br StatsRef) GetPath() string {
	return br.stats
}

// String returns the fully qualified reference name e.g. refs/heads/main
func (br StatsRef) String() string {
	return String(br)
}
