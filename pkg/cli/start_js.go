// Copyright 2017 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !windows,!unix

package cli

import (
    "os"

    "github.com/cockroachdb/cockroach/pkg/cli/cliflags"
)

// drainSignals are the signals that will cause the server to drain and exit.
//
// If two drain signals are seen, the second drain signal will be reraised
// without a signal handler. The default action of any signal listed here thus
// must terminate the process.
var drainSignals = []os.Signal{}

func handleSignalDuringShutdown(sig os.Signal) { }

var startBackground bool

func init() {
    for _, cmd := range StartCmds {
        BoolFlag(cmd.Flags(), &startBackground, cliflags.Background, false)
    }
}

func maybeRerunBackground() (bool, error) {
    return false, nil
}

func disableOtherPermissionBits() { }
