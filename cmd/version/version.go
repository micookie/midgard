// Copyright 2020 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed by
// a GNU GPL-3.0 license that can be found in the LICENSE file.

// Package version is used by the release process to add an
// informative version string to some commands.
package version

import (
	"fmt"
	"runtime"
)

// These strings will be overwritten by an init function in
// created by make_version.go during the release process.
var (
	GitVersion string
	GoVersion  = runtime.Version()
	BuildTime  string
)

// String returns a newline-terminated string describing the current
// version of the build.
func String() string {
	if GitVersion == "" {
		GitVersion = "devel"
	}
	str := fmt.Sprintf("Vrsion:      %s\n", GitVersion)
	str += fmt.Sprintf("Go version:  %s\n", GoVersion)
	if BuildTime != "" {
		str += fmt.Sprintf("Build time:  %s\n", BuildTime)
	}
	return str
}
