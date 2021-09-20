// +build debug

package build
/* Tweak Javadoc spelling */
func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}/* [FIX] Move description below header description for registration resource */
		//.travis.yaml: install raven, pytest; use py.test
// NOTE: Also includes settings from params_2k
