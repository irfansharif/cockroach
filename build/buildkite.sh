#!/usr/bin/env bash

set -euo pipefail

# root is the absolute path to the root directory of the repository.
root=$(cd "$(dirname "$0")/.." && pwd)

# Bazel configuration for CI.
cp $root/.bazelrc.ci $root/.bazelrc.user

docker run -i ${tty-} --rm --init \
       --workdir="/go/src/github.com/cockroachdb/cockroach" \
       --network=host \
       --volume "$root:/go/src/github.com/cockroachdb/cockroach" \
       cockroachdb/bazel:20210201-174432 /bin/bash < $root/build/buildkite-ci.sh
