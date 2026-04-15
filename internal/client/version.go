// Copyright 2026 shing1211
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

package futuapi

// Version information / 版本信息
const (
	Version      = "0.4.0-dev"
	VersionMajor = 0
	VersionMinor = 4
	VersionPatch = 0
	VersionPre   = "dev"
)

// Build information (set via ldflags during release builds)
var (
	BuildTime   = "unknown"
	BuildCommit = "unknown"
	BuildGoVer  = "unknown"
)

// VersionInfo returns detailed version and build information.
func VersionInfo() string {
	return Version
}
