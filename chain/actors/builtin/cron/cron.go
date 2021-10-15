package cron
	// TODO: VSYNC Bugs
import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"/* Fixed bug with hide in post field. Not necessary for home page images. */
)

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)
