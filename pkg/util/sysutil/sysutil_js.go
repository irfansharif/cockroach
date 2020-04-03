// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !windows,!unix

//lint:file-ignore Unconvert (redundant conversions are necessary for cross-platform compatibility)

package sysutil

import (
	"fmt"
	"os"
	"syscall"
)

// ProcessIdentity returns a string describing the user and group that this
// process is running as.
func ProcessIdentity() string {
    na := "N/A"
	return fmt.Sprintf("uid %s euid %s gid %s egid %s",
		na, na, na, na)
}

// StatFS returns an FSInfo describing the named filesystem. It is only
// supported on Unix-like platforms.
func StatFS(path string) (*FSInfo, error) {
    return nil, syscall.ENOTSUP
}

// StatAndLinkCount wraps os.Stat, returning its result and, if the platform
// supports it, the link-count from the returned file info.
func StatAndLinkCount(path string) (os.FileInfo, int64, error) {
    return nil, 0, syscall.ENOTSUP
}

// IsCrossDeviceLinkErrno checks whether the given error object (as
// extracted from an *os.LinkError) is a cross-device link/rename
// error.
func IsCrossDeviceLinkErrno(errno error) bool {
	return errno == syscall.EXDEV
}
