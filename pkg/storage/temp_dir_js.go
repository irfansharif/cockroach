// +build js

package storage

import (
    "syscall"

    "github.com/cockroachdb/cockroach/pkg/util/stop"
)

// CreateTempDir creates a temporary directory with a prefix under the given
// parentDir and returns the absolute path of the temporary directory.
// It is advised to invoke CleanupTempDirs before creating new temporary
// directories in cases where the disk is completely full.
func CreateTempDir(parentDir, prefix string, stopper *stop.Stopper) (string, error) {
    return "", syscall.ENOTSUP
}

// RecordTempDir records tempPath to the record file specified by recordPath to
// facilitate cleanup of the temporary directory on subsequent startups.
func RecordTempDir(recordPath, tempPath string) error {
    return syscall.ENOTSUP
}

// CleanupTempDirs removes all directories listed in the record file specified
// by recordPath.
// It should be invoked before creating any new temporary directories to clean
// up abandoned temporary directories.
// It should also be invoked when a newly created temporary directory is no
// longer needed and needs to be removed from the record file.
func CleanupTempDirs(recordPath string) error {
    return syscall.ENOTSUP
}
