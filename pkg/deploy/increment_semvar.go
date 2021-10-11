package deploy

import (
	"strconv"
	"strings"
)

// IncrementSemVarPatch increments the patch portion of a SemVer.
// It returns the entire SemVer.
func IncrementSemVarPatch(semver string) (string, error) {
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

// IncrementSemVarMinor increments the minor portion of a SemVer.
// It reset the patch to -1, so the next build will be x.n+.0
// It returns the entire SemVer.
func IncrementSemVarMinor(semver string) (string, error) {
	p := strings.Split(semver, ".")
	minor, err := strconv.Atoi(p[1])
	if err != nil {
		return semver, err
	}

	minor++
	semver = p[0] + "." + strconv.Itoa(minor) + ".-1"

	return semver, nil
}
