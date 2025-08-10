// Copyright 2025 Jalu Nugroho
// SPDX-License-Identifier: MIT

package main

import "github.com/nrhox/lightgate/cmd"

var version = "dev"

func main() {
	// execute server
	cmd.Execute(version)
}
