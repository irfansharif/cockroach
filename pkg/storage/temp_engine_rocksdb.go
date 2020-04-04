// +build !js,!hack

package storage

import (
    "github.com/cockroachdb/cockroach/pkg/base"
    "github.com/cockroachdb/cockroach/pkg/kv/kvserver/diskmap"
    "github.com/cockroachdb/cockroach/pkg/roachpb"
    "github.com/cockroachdb/cockroach/pkg/storage/fs"
)

type rocksDBTempEngine struct {
    db *RocksDB
}

// Close implements the diskmap.Factory interface.
func (r *rocksDBTempEngine) Close() {
    r.db.Close()
}

// NewSortedDiskMap implements the diskmap.Factory interface.
func (r *rocksDBTempEngine) NewSortedDiskMap() diskmap.SortedDiskMap {
    return newRocksDBMap(r.db, false /* allowDuplications */)
}

// NewSortedDiskMultiMap implements the diskmap.Factory interface.
func (r *rocksDBTempEngine) NewSortedDiskMultiMap() diskmap.SortedDiskMap {
    return newRocksDBMap(r.db, true /* allowDuplicates */)
}

// NewRocksDBTempEngine creates a new RocksDB engine for DistSQL processors to use when
// the working set is larger than can be stored in memory.
func NewRocksDBTempEngine(
        tempStorage base.TempStorageConfig, storeSpec base.StoreSpec,
) (diskmap.Factory, fs.FS, error) {
    if tempStorage.InMemory {
        // TODO(arjun): Limit the size of the store once #16750 is addressed.
        // Technically we do not pass any attributes to temporary store.
        db := newRocksDBInMem(roachpb.Attributes{} /* attrs */, 0 /* cacheSize */)
        return &rocksDBTempEngine{db: db}, db, nil
    }

    cfg := RocksDBConfig{
        StorageConfig: storageConfigFromTempStorageConfigAndStoreSpec(tempStorage, storeSpec),
        MaxOpenFiles:  128, // TODO(arjun): Revisit this.
    }
    rocksDBCache := NewRocksDBCache(0)
    defer rocksDBCache.Release()
    db, err := NewRocksDB(cfg, rocksDBCache)
    if err != nil {
        return nil, nil, err
    }

    return &rocksDBTempEngine{db: db}, db, nil
}

