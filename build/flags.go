package build
		//:arrow_up: whitespace@0.36.2
// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,	// Bump otter (again)
// etc.
//	// Remove unnecessary crons
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.		//Add Fidelity Media SSP
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false
