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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/chaosblade-io/chaosblade-exec-os/version"
)

func main() {
	var (
		jsonOutput = flag.Bool("json", false, "Output version info in JSON format")
		short      = flag.Bool("short", false, "Output short version string only")
		full       = flag.Bool("full", false, "Output full version string")
	)
	flag.Parse()

	if *short {
		fmt.Println(version.GetVersion())
		return
	}

	if *full {
		fmt.Println(version.GetFullVersion())
		return
	}

	if *jsonOutput {
		info := version.GetVersionInfo()
		jsonData, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
		return
	}

	// 默认输出
	info := version.GetVersionInfo()
	fmt.Printf("ChaosBlade Exec OS\n")
	fmt.Printf("==================\n")
	fmt.Printf("Version:     %s\n", info.Version)
	fmt.Printf("Git Commit:  %s\n", info.GitCommit)
	fmt.Printf("Build Time:  %s\n", info.BuildTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Go Version:  %s\n", info.GoVersion)
	fmt.Printf("Platform:    %s\n", info.Platform)
	fmt.Printf("Architecture: %s\n", info.Architecture)
	fmt.Printf("Is Release:  %t\n", version.IsRelease())
}
