package deploy

import (
	"strconv"
	"strings"
)

// IncrementSemVar increments the patch portion of a SemVer.
// It returns the entires SemVer.
func IncrementSemVar(semver string) (string, error) {
	if semver == "0.0.-1" || semver == "0.0.0" {
		return "0.0.1", nil
	}

	p := strings.Split(semver, ".")
	patch, err := strconv.Atoi(p[2])
	if err != nil {
		return semver, err
	}

	patch++
	semver = p[0] + "." + p[1] + "." + strconv.Itoa(patch)

	return semver, nil
}
