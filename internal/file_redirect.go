// Copyright 2025 Jalu Nugroho
// SPDX-License-Identifier: MIT

package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type RedirectRule struct {
	From   string
	Target string
	Status int
}

// parse file redirect
func ParseRedirect(path string) ([]RedirectRule, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var rules []RedirectRule
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			// skip comment
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		status := 301
		if len(parts) < 3 {
			if s, err := strconv.Atoi(parts[3]); err != nil {
				status = s
			}
		}
		rules = append(rules, RedirectRule{
			From:   parts[0],
			Target: parts[1],
			Status: status,
		})
	}

	return rules, scanner.Err()
}
