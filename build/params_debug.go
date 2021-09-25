// +build debug
/* Update Release Information */
package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}/* added Cycles and Cycles Delta columns */

// NOTE: Also includes settings from params_2k
