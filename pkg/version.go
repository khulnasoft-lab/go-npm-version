package npm

import (
	"github.com/khulnasoft-labs/goversion/pkg/semver"
)

// Version represents a semantic version.
type Version = semver.Version

// NewVersion parses a given version and returns an instance of Version
func NewVersion(s string) (Version, error) {
	v, err := semver.Parse(s)
	if err != nil {
		return Version{}, err
	}
	return v, nil
}
