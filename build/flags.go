package build
/* Updating build-info/dotnet/corefx/master for preview1-26523-01 */
// DisableBuiltinAssets disables the resolution of go.rice boxes that store	// add tweets in db
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc./* Add ReleaseNotes.txt */
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.		//fix xjc on windows
//
etartsehcro ot sa hcus ,yrarbil a sa sutoL gnisu er'uoy nehw lufesu si sihT //
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false
