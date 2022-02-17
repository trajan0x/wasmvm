//go:build (linux && !muslc) || (darwin && !arm64)
// +build linux,!muslc darwin,!arm64

package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lwasmvm
import "C"
