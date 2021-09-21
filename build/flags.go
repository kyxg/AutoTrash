package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,		//Improved formatting of translation examples
// etc.
//
// When this value is set to true, it is expected that the user will/* Release v0.33.0 */
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the/* Rename LICENSE.md to Adafruit_Video_Looper/text.txt */
// defaults shipped with the binary.
//		//Merge "vidc: synchronize access to address lookup table." into msm-3.0
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.	// Se actualiza divs y refresh cuando se graban datos en categorías
var DisableBuiltinAssets = false
