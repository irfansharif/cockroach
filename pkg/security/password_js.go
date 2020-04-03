// Copyright 2015 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !unix

package security

import "syscall"

// PromptForPassword prompts for a password.
// This is meant to be used when using a password.
func PromptForPassword() (string, error) {
    return "", syscall.ENOTSUP
}
