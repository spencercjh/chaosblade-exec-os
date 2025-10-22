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

import "fmt"

type cgroupSubsysFormatInvalidError struct {
	line string
}

type mountPointFormatInvalidError struct {
	line string
}

type pathNotExposedFromMountPointError struct {
	mountPoint string
	root       string
	path       string
}

func (err cgroupSubsysFormatInvalidError) Error() string {
	return fmt.Sprintf("invalid format for CGroupSubsys: %q", err.line)
}

func (err mountPointFormatInvalidError) Error() string {
	return fmt.Sprintf("invalid format for MountPoint: %q", err.line)
}

func (err pathNotExposedFromMountPointError) Error() string {
	return fmt.Sprintf("path %q is not a descendant of mount point root %q and cannot be exposed from %q", err.path, err.root, err.mountPoint)
}
