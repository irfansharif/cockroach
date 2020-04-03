// Copyright 2014 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package storage

import (
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/storage/enginepb"
	"github.com/cockroachdb/cockroach/pkg/util/protoutil"
)

func serializeMergeInputs(sources ...roachpb.InternalTimeSeriesData) ([][]byte, error) {
	// Wrap each proto in an inlined MVCC value, and marshal each wrapped value
	// to bytes. This is the format required by the engine.
	srcBytes := make([][]byte, 0, len(sources))
	var val roachpb.Value
	for _, src := range sources {
		if err := val.SetProto(&src); err != nil {
			return nil, err
		}
		bytes, err := protoutil.Marshal(&enginepb.MVCCMetadata{
			RawBytes: val.RawBytes,
		})
		if err != nil {
			return nil, err
		}
		srcBytes = append(srcBytes, bytes)
	}
	return srcBytes, nil
}

func deserializeMergeOutput(mergedBytes []byte) (roachpb.InternalTimeSeriesData, error) {
	// Unmarshal merged bytes and extract the time series value within.
	var meta enginepb.MVCCMetadata
	if err := protoutil.Unmarshal(mergedBytes, &meta); err != nil {
		return roachpb.InternalTimeSeriesData{}, err
	}
	mergedTS, err := MakeValue(meta).GetTimeseries()
	if err != nil {
		return roachpb.InternalTimeSeriesData{}, err
	}
	return mergedTS, nil
}

