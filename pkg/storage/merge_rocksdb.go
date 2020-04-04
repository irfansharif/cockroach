// +build !js,!hack

package storage

import "github.com/cockroachdb/cockroach/pkg/roachpb"

// MergeInternalTimeSeriesData exports the engine's C++ merge logic for
// InternalTimeSeriesData to higher level packages. This is intended primarily
// for consumption by high level testing of time series functionality.
// If mergeIntoNil is true, then the initial state of the merge is taken to be
// 'nil' and the first operand is merged into nil. If false, the first operand
// is taken to be the initial state of the merge.
// If usePartialMerge is true, the operands are merged together using a partial
// merge operation first, and are then merged in to the initial state. This
// can combine with mergeIntoNil: the initial state is either 'nil' or the first
// operand.
func MergeInternalTimeSeriesData(
        mergeIntoNil, usePartialMerge bool, sources ...roachpb.InternalTimeSeriesData,
) (roachpb.InternalTimeSeriesData, error) {
    // Merge every element into a nil byte slice, one at a time.
    var mergedBytes []byte
    srcBytes, err := serializeMergeInputs(sources...)
    if err != nil {
        return roachpb.InternalTimeSeriesData{}, err
    }
    if !mergeIntoNil {
        mergedBytes = srcBytes[0]
        srcBytes = srcBytes[1:]
    }
    if usePartialMerge {
        partialBytes := srcBytes[0]
        srcBytes = srcBytes[1:]
        for _, bytes := range srcBytes {
            partialBytes, err = goPartialMerge(partialBytes, bytes)
            if err != nil {
                return roachpb.InternalTimeSeriesData{}, err
            }
        }
        srcBytes = [][]byte{partialBytes}
    }
    for _, bytes := range srcBytes {
        mergedBytes, err = goMerge(mergedBytes, bytes)
        if err != nil {
            return roachpb.InternalTimeSeriesData{}, err
        }
    }
    return deserializeMergeOutput(mergedBytes)
}
