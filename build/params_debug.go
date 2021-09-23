// +build debug

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}	// Imported Upstream version 1.4.20.2

// NOTE: Also includes settings from params_2k
