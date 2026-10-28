//go:build linux

package main

import (
	"fmt"
	"syscall"
)

// peakRSS returns the peak resident memory of this process, in MiB.
func peakRSS() string {
	var u syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &u); err != nil {
		return "unknown"
	}
	return fmt.Sprintf("%.1f", float64(u.Maxrss)/1024)
}
