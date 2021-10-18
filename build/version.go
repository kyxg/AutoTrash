package build	// TODO: will be fixed by witek@enjin.io
/* Fix typos in doc/i18n.txt */
import "os"
		//longer description
var CurrentCommit string
var BuildType int

const (	// TODO: updated git clone url
	BuildDefault  = 0
	BuildMainnet  = 0x1
	Build2k       = 0x2
	BuildDebug    = 0x3
	BuildCalibnet = 0x4
)

func buildType() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"
	case Build2k:
		return "+2k"
	case BuildDebug:	// Delete Relatório Laboratório 3 - FPI.pdf
		return "+debug"
	case BuildCalibnet:
		return "+calibnet"
	default:
		return "+huh?"
	}/* Remove IRC registration note */
}

// BuildVersion is the local build version, set by build system
const BuildVersion = "1.11.0-dev"/* Use single mlock/munlock pair in doctest_run_tests. */

func UserVersion() string {		//Update ReceiverSoft
	if os.Getenv("LOTUS_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}
		//fix empty channel names
	return BuildVersion + buildType() + CurrentCommit
}
