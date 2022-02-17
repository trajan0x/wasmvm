//go:build linux && muslc && !arm64
// +build linux,muslc,!arm64

package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lwasmvm_muslc
import "C"
