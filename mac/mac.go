// Package mac gets mac type/name/version
package mac

import (
	"regexp"
)

var (
	versionPattern = "([0-9]+\\.[0-9])+(\\.[0-9])?"
)

// Mac struct
type Mac struct {
	uname string
}

// New creates Mac instance
func New(uname string) *Mac {
	return &Mac{
		uname: uname,
	}
}

// Name gets mac name
func (mac Mac) Name() string {
	return versions[mac.info()]["name"]
}

// Version gets mac version
func (mac Mac) Version() string {
	return versions[mac.info()]["version"]
}

// Type gets mac version
func (mac Mac) Type() string {
	return "mac"
}

func (mac Mac) info() string {
	rVersion := regexp.MustCompile(versionPattern)
	return rVersion.FindAllStringSubmatch(mac.uname, 1)[0][0]
}
