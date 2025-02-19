// Copyright 2023 PingCAP, Inc.
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

package integrationtests

import (
	"sync/atomic"
	"testing"

	"github.com/pingcap/tidb/pkg/disttask/framework/proto"
	"github.com/pingcap/tidb/pkg/disttask/framework/testutil"
	"github.com/pingcap/tidb/pkg/testkit/testfailpoint"
	"github.com/stretchr/testify/require"
)

func TestFrameworkRollback(t *testing.T) {
	c := testutil.NewTestDXFContext(t, 2, 16, true)
	testutil.RegisterTaskTypeForRollback(t, c.MockCtrl, testutil.GetMockRollbackSchedulerExt(c.MockCtrl), c.TestContext)
	var counter atomic.Int32
	testfailpoint.EnableCall(t, "github.com/pingcap/tidb/pkg/disttask/framework/scheduler/afterRefreshTask",
		func(task *proto.Task) {
			if counter.Add(1) <= 2 {
				if task.State == proto.TaskStateRunning {
					require.NoError(t, c.TaskMgr.CancelTask(c.Ctx, task.ID))
				}
			}
		},
	)

	task := testutil.SubmitAndWaitTask(c.Ctx, t, "key1", "", 1)
	require.Equal(t, proto.TaskStateReverted, task.State)
}
