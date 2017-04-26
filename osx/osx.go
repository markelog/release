// Package osx gets osx type/name/version
package osx

import (
	"io/ioutil"
	"os/exec"
	"regexp"
)

var (
	versionPattern  = "System Version:[[:alpha:]\\s]+([0-9]+\\.[0-9]+\\.[0-9]+)"
	codenamePattern = "SOFTWARE LICENSE AGREEMENT FOR ([\\w\\s]+)"
)

// OSX struct
type OSX struct{}

// New creates OSX instance
func New() *OSX {
	return &OSX{}
}

// Name gets osx name
func (osx OSX) Name() (result string) {
	path := "/System/Library/CoreServices/Setup Assistant.app/Contents/Resources/en.lproj/OSXSoftwareLicense.rtf"
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return ""
	}

	data := string(bytes)
	rCodename := regexp.MustCompile(codenamePattern)

	subs := rCodename.FindAllStringSubmatch(data, 1)
	if len(subs) == 0 {
		return
	}

	find := subs[0]
	if len(find) < 1 {
		return
	}

	return string(find[1])
}

// Version gets osx version
func (osx OSX) Version() (result string) {
	data := Execute("system_profiler", "SPSoftwareDataType")

	rVersion := regexp.MustCompile(versionPattern)

	subs := rVersion.FindAllStringSubmatch(data, 1)
	if len(subs) == 0 {
		return
	}

	find := subs[0]
	if len(find) < 1 {
		return
	}

	return find[1]
}

// Type gets osx version
func (osx OSX) Type() string {
	return "osx"
}

func Execute(name, arg string) string {
	output, _ := exec.Command(name, arg).Output()
	return string(output)
}
