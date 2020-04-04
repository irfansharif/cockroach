// Copyright 2015 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// +build !js,!hack

package security

import (
    "fmt"
    "os"

    "golang.org/x/crypto/ssh/terminal"
)


// PromptForPassword prompts for a password.
// This is meant to be used when using a password.
func PromptForPassword() (string, error) {
    fmt.Print("Enter password: ")
    password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
    if err != nil {
        return "", err
    }
    // Make sure stdout moves on to the next line.
    fmt.Print("\n")

    return string(password), nil
}
