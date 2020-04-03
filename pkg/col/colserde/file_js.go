// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !windows,!unix

package colserde

import (
    "syscall"
)

// NewFileDeserializerFromPath constructs a FileDeserializer by reading it from
// a file.
func NewFileDeserializerFromPath(path string) (*FileDeserializer, error) {
    return nil, syscall.ENOTSUP
}
