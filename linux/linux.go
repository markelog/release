// Package linux gets linux type/name/version
package linux

import (
	"os/exec"
	"regexp"
	"strings"
)

var (
	versionPattern = "Release:\\s+([0-9]+\\.[0-9]+)"
	namePattern    = "Codename:\\s+([a-z]+)"
	typePattern    = "Distributor ID:\\s+([a-zA-Z]+)"
)

// Linux struct
type Linux struct {
	lsbrelease string
}

// New creates Linux instance
func New() *Linux {
	return &Linux{
		lsbrelease: LSBRelease(),
	}
}

// Name gets linux name
func (linux Linux) Name() string {
	if len(linux.lsbrelease) == 0 {
		return ""
	}

	rName := regexp.MustCompile(namePattern)
	return rName.FindAllStringSubmatch(linux.lsbrelease, 1)[0][1]
}

// Version gets linux version
func (linux Linux) Version() string {
	if len(linux.lsbrelease) == 0 {
		return ""
	}

	rVersion := regexp.MustCompile(versionPattern)
	return rVersion.FindAllStringSubmatch(linux.lsbrelease, 1)[0][1]
}

// Type gets type version
func (linux Linux) Type() string {
	if len(linux.lsbrelease) == 0 {
		return ""
	}

	rType := regexp.MustCompile(typePattern)
	result := rType.FindAllStringSubmatch(linux.lsbrelease, 1)[0][1]

	return strings.ToLower(result)
}

// LSBRelease executes `lsb_release -a` command (exposed for tests)
func LSBRelease() string {
	result, _ := exec.Command("lsb_release", "-a").Output()

	return strings.TrimSpace(string(result))
}
