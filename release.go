// Package release get OS type/name/version
package release

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/markelog/release/hollow"
	"github.com/markelog/release/linux"
	"github.com/markelog/release/mac"
)

type release interface {
	Type() string
	Name() string
	Version() string
}

var (

	// GOOS is reference to runtime.GOOS (needed for tests)
	GOOS = runtime.GOOS
)

// Name gets os name
func Name() (name string) {
	return new().Name()
}

// Version gets os version
func Version() (version string) {
	return new().Version()
}

// Type gets type version
func Type() (typa string) {
	return new().Type()
}

// All gets all avaliable info
func All() (typa, name, version string) {
	instance := new()

	typa, name, version = instance.Type(), instance.Name(), instance.Version()

	return
}

func new() (rel release) {
	if GOOS == "darwin" {
		rel = mac.New(Uname())
		return
	}

	if GOOS == "linux" {
		rel = linux.New(LSBRelease())
		return
	}

	return hollow.New()
}

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
