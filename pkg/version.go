package npm

import "github.com/khulnasoft-lab/goversion/pkg/semver"

// Version represents a semantic version.
type Version = semver.Version

// NewVersion parses a given version and returns an instance of Version
func NewVersion(s string) (Version, error) {
	return semver.Parse(s)
}