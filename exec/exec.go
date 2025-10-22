/*
 * Copyright 2025 The ChaosBlade Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package exec

import (
	"context"
	"fmt"
	"strings"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

// todo
var cl = channel.NewLocalChannel()

// stop hang process
func Destroy(ctx context.Context, c spec.Channel, action string) *spec.Response {
	suid := ctx.Value(spec.Uid)
	/* If suid is specified, it will be deleted exactly
	 * according to suid, otherwise it will be based on action. */
	if suid != nil && suid != spec.UnknownUid && suid != "" {
		ctx = context.WithValue(ctx, channel.ProcessKey, suid)
	} else {
		ctx = context.WithValue(ctx, channel.ProcessKey, action)
	}

	// Adapt to old versions.
	originalBin := ctx.Value("bin")
	pids := make([]string, 0)
	if originalBin != nil {
		originalPids, _ := cl.GetPidsByProcessName(originalBin.(string), ctx)
		pids = append(pids, originalPids...)
	}

	ps, _ := cl.GetPidsByProcessName(spec.ChaosOsBin, ctx)
	pids = append(ps, pids...)
	if len(pids) == 0 {
		// If no processes found, consider the destroy operation successful
		// This can happen when processes have already been cleaned up or never existed
		return spec.ReturnSuccess("no processes found to destroy")
	}
	return cl.Run(ctx, "kill", fmt.Sprintf(`-9 %s`, strings.Join(pids, " ")))
}

func CheckFilepathExists(ctx context.Context, cl spec.Channel, filepath string) bool {
	response := cl.Run(ctx, fmt.Sprintf("[ -e %s ] && echo true || echo false", filepath), "")
	if response.Success && strings.Contains(response.Result.(string), "true") {
		return true
	}
	return false
}
