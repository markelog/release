package release

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/markelog/release/linux"
	"github.com/markelog/release/mac"
)

var (

	// GOOS is reference to runtime.GOOS (needed for tests)
	GOOS = runtime.GOOS
)

// Uname executes `uname -a` command (exposed for tests)
func Uname() string {
	result, _ := exec.Command("uname", "-a").Output()

	return strings.TrimSpace(string(result))
}

// LSBRelease executes `lsb_release -a` command (exposed for tests)
func LSBRelease() string {
	result, _ := exec.Command("lsb_release", "-a").Output()

	return strings.TrimSpace(string(result))
}

// Name gets os name
func Name() string {
	if GOOS == "darwin" {
		return mac.New(Uname()).Name()
	}

	if GOOS == "linux" {
		return linux.New(LSBRelease()).Name()
	}

	return ""
}

// Version gets os version
func Version() string {
	if GOOS == "darwin" {
		return mac.New(Uname()).Version()
	}

	if GOOS == "linux" {
		return linux.New(LSBRelease()).Version()
	}

	return ""
}

// Type gets type version
func Type() string {
	if GOOS == "darwin" {
		return "mac"
	}

	if GOOS == "linux" {
		return linux.New(LSBRelease()).Type()
	}

	return ""
}
