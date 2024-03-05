package io

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestGetOperatingSystem(t *testing.T) {
	t.Run("Override Environment Variable", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "linux")
		expectedOS := LinuxOs
		actualOS := getOperatingSystem()

		if actualOS != expectedOS {
			t.Errorf("Expected operating system to be %v, but got %v", expectedOS, actualOS)
		}
	})

	t.Run("Default Operating System", func(t *testing.T) {
		os.Setenv("TEST_OS_OVERRIDE", "")
		expectedOS := getExpectedDefaultOperatingSystem()
		actualOS := getOperatingSystem()

		if actualOS != expectedOS {
			t.Errorf("Expected operating system to be %v, but got %v", expectedOS, actualOS)
		}
	})
}

func getExpectedDefaultOperatingSystem() OperatingSystem {
	switch strings.ToLower(runtime.GOOS) {
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
