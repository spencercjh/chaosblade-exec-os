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

package cgroups

import "testing"

func Test_replaceCgroupFsPathForDaemonSetPod(t *testing.T) {
	type args struct {
		mountPointPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pass",
			args: args{
				mountPointPath: "/sys/fs/cgroup/cpu",
			},
			want: "/host-sys/fs/cgroup/cpu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceCgroupFsPathForDaemonSetPod(tt.args.mountPointPath, "/host-sys/fs/cgroup/"); got != tt.want {
				t.Errorf("replaceCgroupFsPathForDaemonSetPod() = %v, want %v", got, tt.want)
			}
		})
	}
}
