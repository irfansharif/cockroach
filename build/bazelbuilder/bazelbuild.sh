#!/usr/bin/env bash

set -xeuo pipefail

bazel build //pkg/cmd/cockroach-short --remote_cache='http://localhost:9090'
bazel test //pkg:small_tests --remote_cache='http://localhost:9090'
