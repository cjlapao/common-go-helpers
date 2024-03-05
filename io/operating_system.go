package io

import (
	"os"
	"runtime"
	"strings"
)

// OperatingSystem enum
type OperatingSystem int

// Defines the operating system Enum
const (
	WindowsOs OperatingSystem = iota
	LinuxOs
	MacOs
	UnknownOs
)

// getOperatingSystem returns the operating system
/*
Get the operating system name and return it as an OperatingSystem constant.

Args:
	None
Returns:
	The operating system.
*/
func getOperatingSystem() OperatingSystem {
	envOSOverride := strings.ToLower(os.Getenv("TEST_OS_OVERRIDE"))
	var os string
	if envOSOverride != "" {
		os = envOSOverride
	} else {
		os = runtime.GOOS
	}
	switch strings.ToLower(os) {
	case "linux":
		return LinuxOs
	case "windows":
		return WindowsOs
	case "darwin":
		return MacOs
	default:
		return UnknownOs
	}
}
