// Package osx gets osx type/name/version
package osx

import (
	"regexp"
)

var (
	versionPattern = "([0-9]+\\.[0-9])+(\\.[0-9])?"
)

// OSX struct
type OSX struct {
	uname string
}

// New creates OSX instance
func New(uname string) *OSX {
	return &OSX{
		uname: uname,
	}
}

// Name gets osx name
func (osx OSX) Name() string {
	return versions[osx.info()]["name"]
}

// Version gets osx version
func (osx OSX) Version() string {
	return versions[osx.info()]["version"]
}

// Type gets osx version
func (osx OSX) Type() string {
	return "osx"
}

func (osx OSX) info() string {
	rVersion := regexp.MustCompile(versionPattern)
	return rVersion.FindAllStringSubmatch(osx.uname, 1)[0][0]
}
