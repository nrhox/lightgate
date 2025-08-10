// Copyright 2025 Jalu Nugroho
// SPDX-License-Identifier: MIT

package internal

import (
	"flag"
	"fmt"
	"os"
)

const helpText = `Light Gate - Static + SPA server

The program starts a simple HTTP server to serve
files and handle requests.

Usage:
  lightgate [options]

Options:
`

const exampleText = `
Example:
  lightgate -p 3000 -d ./
  lightgate -p 3000 -d ./public
  lightgate -p 3000 -d ./ -n ./assets/404.html
  lightgate -p 3000 -d ./ -i ./about.html -n ./assets/404.html
`

type FlagOption struct {
	Port         *int
	Dir          *string
	RedirectPath *string
	IndexFile    *string
	File404      *string
	ShowVersion  *bool
	verbose      *bool
}

var (
	port         = flag.Int("p", DEFAULT_PORT, "Port server")
	dirPath      = flag.String("d", ".", "Directory file")
	redirectPath = flag.String("r", "", "Path to _redirects file (optional)")
	indexFile    = flag.String("i", "index.html", "Index file endpoint (optional)")
	File404      = flag.String("n", "404.html", "404 file (optional)")
	showVersion  = flag.Bool("v", false, "Show current version")
	verbose      = flag.Bool("ver", false, "Display detailed logs for each request")
)

func ReadFlags() FlagOption {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, helpText)
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, exampleText)
	}
	flag.Parse()

	return FlagOption{
		Port:         port,
		Dir:          dirPath,
		RedirectPath: redirectPath,
		IndexFile:    indexFile,
		File404:      File404,
		ShowVersion:  showVersion,
		verbose:      verbose,
	}
}
