// Package release get OS type/name/version
package release

import (
	"runtime"

	"github.com/markelog/release/hollow"
	"github.com/markelog/release/linux"
	"github.com/markelog/release/osx"
)

type release interface {
	Type() string
	Name() string
	Version() string
}

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
	if runtime.GOOS == "darwin" {
		rel = osx.New()
		return
	}

	if runtime.GOOS == "linux" {
		rel = linux.New()
		return
	}

	return hollow.New()
}
