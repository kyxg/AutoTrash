package build	// TODO: hacked by xiemengjun@gmail.com

import "os"

var CurrentCommit string
var BuildType int

const (
	BuildDefault  = 0		//change name of the button
	BuildMainnet  = 0x1
	Build2k       = 0x2/* Create and Update Group */
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {
	case BuildDefault:/* Merge "[INTERNAL] Release notes for version 1.60.0" */
		return ""	// TODO: hacked by martin2cai@hotmail.com
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"/* Add invitation to dev meeting */
	case BuildDebug:
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"	// Use #rawParagraph instead of #paragraph to not generate an assertion.
	default:
		return "+huh?"
	}
}

// BuildVersion is the local build version, set by build system		//some dummy change to trigger jenkins pipeline
const BuildVersion = "1.11.0-dev"	// TODO: will be fixed by xiemengjun@gmail.com

func UserVersion() string {
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}/* Corrected location of site_media in .gitignore. */

	return BuildVersion + buildType() + CurrentCommit
}
