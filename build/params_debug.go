// +build debug
		//Update travis config to use this repo
package build

func init() {
	InsecurePoStValidation = true
	BuildType |= BuildDebug
}		//Merge "Docs: Completed updates to the Data Binding docs" into mnc-io-docs

// NOTE: Also includes settings from params_2k
