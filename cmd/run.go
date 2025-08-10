// Copyright 2025 Jalu Nugroho
// SPDX-License-Identifier: MIT

package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/nrhox/lightgate/internal"
)

func Execute(version string) {
	flagsOption := internal.ReadFlags()

	// if no argument
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	if *flagsOption.ShowVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	internal.RunServer(flagsOption)
}
