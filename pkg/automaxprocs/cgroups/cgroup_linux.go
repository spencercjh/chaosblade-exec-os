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
	"bufio"
	"context"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/chaosblade-io/chaosblade-spec-go/log"
)

// CGroup represents the data structure for a Linux control group.
type CGroup struct {
	path string
}

// NewCGroup returns a new *CGroup from a given path.
func NewCGroup(path string) *CGroup {
	return &CGroup{path: path}
}

// Path returns the path of the CGroup*.
func (cg *CGroup) Path() string {
	return cg.path
}

// ParamPath returns the path of the given cgroup param under itself.
func (cg *CGroup) ParamPath(param string) string {
	return filepath.Join(cg.path, param)
}

// readFirstLine reads the first line from a cgroup param file.
func (cg *CGroup) readFirstLine(param string) (string, error) {
	path := cg.ParamPath(param)
	paramFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	log.Infof(context.Background(), "opening file, path: %s", path)
	defer paramFile.Close()

	scanner := bufio.NewScanner(paramFile)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", io.ErrUnexpectedEOF
}

// readInt parses the first line from a cgroup param file as int.
func (cg *CGroup) readInt(param string) (int, error) {
	text, err := cg.readFirstLine(param)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(text)
}
