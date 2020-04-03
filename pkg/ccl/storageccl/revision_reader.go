// Copyright 2016 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package storageccl

import (
	"context"
	"errors"

	"github.com/cockroachdb/cockroach/pkg/kv"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
)

// VersionedValues is similar to roachpb.KeyValue except instead of just the
// value at one time, it contains all the retrieved revisions of the value for
// the key, with the value timestamps set accordingly.
type VersionedValues struct {
	Key    roachpb.Key
	Values []roachpb.Value
}

// GetAllRevisions scans all keys between startKey and endKey getting all
// revisions between startTime and endTime.
// TODO(dt): if/when client gets a ScanRevisionsRequest or similar, use that.
func GetAllRevisions(
	ctx context.Context, db *kv.DB, startKey, endKey roachpb.Key, startTime, endTime hlc.Timestamp,
) ([]VersionedValues, error) {
	return nil, errors.New("unsupported")
}
