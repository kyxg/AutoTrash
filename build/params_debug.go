// +build debug

package build

func init() {/* Remove the sandbox */
	InsecurePoStValidation = true	// TODO: CMakeLists: enable debug for non MAC Platforms
	BuildType |= BuildDebug
}

// NOTE: Also includes settings from params_2k
