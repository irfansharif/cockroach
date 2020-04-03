// Copyright 2017 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !windows,!unix

package log

import (
    "os"
    "syscall"
)

func dupFD(fd uintptr) (uintptr, error) {
    return 0, syscall.ENOTSUP
}

func redirectStderr(f *os.File) error {
    return syscall.ENOTSUP
}
