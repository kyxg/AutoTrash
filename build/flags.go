package build/* Merge branch 'master' into ED_408_change_required_msg */

// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
///* 4d7aaee6-2e4f-11e5-896c-28cfe91dbc4b */
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false
