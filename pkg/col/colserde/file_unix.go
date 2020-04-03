// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !windows,!js

package colserde

import (
    "github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgcode"
    "github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgerror"
    "github.com/edsrzf/mmap-go"
    "os"
)

// NewFileDeserializerFromPath constructs a FileDeserializer by reading it from
// a file.
func NewFileDeserializerFromPath(path string) (*FileDeserializer, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, pgerror.Wrapf(err, pgcode.Io, `opening %s`, path)
    }
    // TODO(dan): This is currently using copy on write semantics because we store
    // the nulls differently in-mem than arrow does and there's an in-place
    // conversion. If we used the same format that arrow does, this could be
    // switched to mmap.RDONLY (it's easy to check, the test fails with a SIGBUS
    // right now with mmap.RDONLY).
    buf, err := mmap.Map(f, mmap.COPY, 0 /* flags */)
    if err != nil {
        return nil, pgerror.Wrapf(err, pgcode.Io, `mmaping %s`, path)
    }
    return newFileDeserializer(buf, buf.Unmap)
}
