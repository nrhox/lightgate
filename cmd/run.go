// Copyright 2025 Jalu Nugroho
// SPDX-License-Identifier: MIT

package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/nrhox/lightgate/internal"
)

const version = "0.1.2"

func Execute() {
	flagsOption := internal.ReadFlags()

	// if no argument
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	if *flagsOption.ShowVersion {
		fmt.Println("lightgate versi", version)
		os.Exit(0)
	}

	internal.RunServer(flagsOption)
}
