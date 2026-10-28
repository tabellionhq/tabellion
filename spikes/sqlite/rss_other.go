//go:build !linux

package main

// peakRSS is measured on Linux only.
func peakRSS() string { return "unknown" }
