// This code has been modified from its original form by The Cockroach Authors.
// All modifications are Copyright 2024 The Cockroach Authors.
//
// Copyright 2019 The etcd Authors
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

package rafttest

import (
	"testing"

	"github.com/cockroachdb/datadriven"
)

func (env *InteractionEnv) handleTickElection(t *testing.T, d datadriven.TestData) error {
	idx := firstAsNodeIdx(t, d)
	return env.Tick(idx, env.Nodes[idx].Config.ElectionTick)
}

func (env *InteractionEnv) handleTickHeartbeat(t *testing.T, d datadriven.TestData) error {
	idx := firstAsNodeIdx(t, d)
	return env.Tick(idx, env.Nodes[idx].Config.HeartbeatTick)
}

// Tick the node at the given index the given number of times.
func (env *InteractionEnv) Tick(idx int, num int64) error {
	for i := int64(0); i < num; i++ {
		env.Nodes[idx].Tick()
	}
	return nil
}
