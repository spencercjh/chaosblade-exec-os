//go:build !linux

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

package runtime

import (
	"context"
)

// GetCPUQuotaToCPUCntByPidForCgroups1 is unsupported for non-linux systems because cgroups is a Linux kernel feature.
func GetCPUQuotaToCPUCntByPidForCgroups1(
	_ context.Context,
	_, _ string,
	_ int,
	_ func(v float64) int,
) (int, float64, CPUQuotaStatus, error) {
	return -1, 1.0, CPUQuotaUndefined, nil
}

// GetCPUQuotaToCPUCntByPidForCgroups2 is unsupported for non-linux systems because cgroups is a Linux kernel feature.
func GetCPUQuotaToCPUCntByPidForCgroups2(
	_ context.Context,
	_, _ string,
	_ int,
	_ func(v float64) int,
) (int, float64, CPUQuotaStatus, error) {
	return -1, 1.0, CPUQuotaUndefined, nil
}
