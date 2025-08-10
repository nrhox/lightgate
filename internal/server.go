// Copyright 2025 Jalu Nugroho
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func TryListenPort(port int) (net.Listener, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("port %d is unavailable", port)
	}
	return ln, nil
}

func hasExt(p string) bool {
	return filepath.Ext(p) != ""
}

func RunServer(option FlagOption) {
	dir, err := filepath.Abs(*option.Dir)
	if err != nil {
		log.Fatalf("invalid dir: %v", err)
	}

	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		log.Fatalf("directory does not exist or is not directory: %s", dir)
	}

	var redirects []RedirectRule
	if *option.RedirectPath != "" {
		rpathAbs, _ := filepath.Abs(*option.RedirectPath)
		rules, err := ParseRedirect(rpathAbs)
		if err != nil {
			log.Printf("warning: cannot load redirects file: %v", err)
		} else {
			redirects = rules
			log.Printf("loaded %d redirect rules from %s", len(redirects), rpathAbs)
		}
	}

	fileExists := func(p string) bool {
		st, err := os.Stat(p)
		if err != nil {
			return false
		}
		return !st.IsDir()
	}

	indexPath := filepath.Join(dir, *option.IndexFile)
	hasIndex := fileExists(indexPath)
	custom404Path := filepath.Join(dir, *option.File404)
	has404 := fileExists(custom404Path)

	// prebuild a map for redirects (exact-match)
	redirectMap := map[string]RedirectRule{}
	for _, r := range redirects {
		redirectMap[r.From] = r
	}

	mux := http.NewServeMux()

	// main handle
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if *option.verbose {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		}

		// exact redirect match
		if rr, ok := redirectMap[r.URL.Path]; ok {
			http.Redirect(w, r, rr.Target, rr.Status)
			return
		}

		// Build safe path
		reqPath := r.URL.Path
		target := SafeJoin(dir, reqPath)

		// If path is directory, try index.html inside it
		if fi, err := os.Stat(target); err == nil && fi.IsDir() {
			idx := filepath.Join(target, *option.IndexFile)
			if fileExists(idx) {
				http.ServeFile(w, r, idx)
				return
			}
			// Continue to fallback rules if no index in dir
		}

		// If file exists, serve it
		if fileExists(target) {
			http.ServeFile(w, r, target)
			return
		}

		// If not found:
		// - If request likely expects HTML (no extension OR Accept contains text/html) and index.html exists -> serve index.html (SPA)
		acceptsHTML := strings.Contains(r.Header.Get("Accept"), "text/html")
		if (!hasExt(r.URL.Path) || acceptsHTML) && hasIndex {
			http.ServeFile(w, r, indexPath)
			return
		}

		// Otherwise 404
		w.WriteHeader(http.StatusNotFound)
		if has404 {
			http.ServeFile(w, r, custom404Path)
			return
		}
		// fallback simple text
		_, _ = w.Write([]byte("404 not found"))
	})

	// try to listen port
	ln, err := TryListenPort(*option.Port)
	if err != nil {
		log.Fatalf("failed to bind: %v", err)
	}
	defer ln.Close()

	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  READ_TIMEOUT,
		WriteTimeout: WRITE_TIMEOUT,
		IdleTimeout:  IDLE_TIMOUT,
	}

	go func() {
		log.Printf("Serving directory %s", dir)
		log.Printf("SPA fallback index.html: %v | custom 404: %v", hasIndex, has404)
		log.Printf("Listening on http://localhost:%d", *option.Port)
		if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	GracefulShutdown(sig, server)
}

// graceful shutdown
func GracefulShutdown(s chan os.Signal, server *http.Server) {
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	<-s
	log.Printf("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	} else {
		log.Printf("server stopped")
	}
}
