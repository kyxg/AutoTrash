// +build debug

package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug/* not valid anymore */
}
/* Further attempts at outputting classified ontology */
// NOTE: Also includes settings from params_2k
