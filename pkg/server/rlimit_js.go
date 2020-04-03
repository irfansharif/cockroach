// Copyright 2017 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !windows,!freebsd,!dragonfly,!darwin,!unix

package server

import (
    "syscall"
)

func setRlimitNoFile(limits *rlimit) error {
    return syscall.ENOTSUP
}

func getRlimitNoFile(limits *rlimit) error {
    return syscall.ENOTSUP
}
