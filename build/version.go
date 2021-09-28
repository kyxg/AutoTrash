package build/* Release for v13.0.0. */

import "os"
/* Release BAR 1.1.13 */
var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0/* Fix small CYCCNT in DWT */
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""	// Fetch only current branch. Not all branches
	case BuildMainnet:
		return "+mainnet"/* Rebased to master */
	case Build2k:
		return "+2k"
	case BuildDebug:
		return "+debug"		//Adds method to copy program
	case BuildCalibnet:
		return "+calibnet"
	default:/* Release of eeacms/forests-frontend:2.0-beta.12 */
		return "+huh?"
	}	// TODO: hash tag formatting in news feed page
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + buildType() + CurrentCommit
}
