//go:build linux

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

import (
	"context"
	"testing"
)

func TestDetectCGroupVersion(t *testing.T) {
	ctx := context.Background()

	// 测试默认路径
	version := DetectCGroupVersion(ctx, "")
	if version != CGroupV1 && version != CGroupV2 {
		t.Errorf("Expected CGroupV1 or CGroupV2, got %v", version)
	}

	// 测试指定路径
	version = DetectCGroupVersion(ctx, "/sys/fs/cgroup")
	if version != CGroupV1 && version != CGroupV2 {
		t.Errorf("Expected CGroupV1 or CGroupV2, got %v", version)
	}
}

func TestIsCGroupV2(t *testing.T) {
	ctx := context.Background()

	// 测试默认路径
	isV2 := IsCGroupV2(ctx, "")
	version := DetectCGroupVersion(ctx, "")
	expected := (version == CGroupV2)

	if isV2 != expected {
		t.Errorf("IsCGroupV2() = %v, expected %v", isV2, expected)
	}
}
