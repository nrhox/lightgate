// Copyright 2025 Jalu Nugroho
// SPDX-License-Identifier: MIT

package internal

import (
	"path/filepath"
	"strings"
)

// filter illegal query path
func SafeJoin(base, reqPath string) string {
	clearReq := filepath.Clean("/" + strings.TrimPrefix(reqPath, "/"))
	full := filepath.Join(base, clearReq)
	rel, err := filepath.Rel(base, full)
	if err != nil {
		return base
	}
	if strings.HasPrefix(rel, "..") {
		return base
	}

	return full
}
