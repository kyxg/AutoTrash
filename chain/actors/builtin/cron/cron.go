package cron

import (/* Update ReleaseNote-ja.md */
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (		//Get rid of compiler warning on 64-bit
	Address = builtin4.CronActorAddr	// f61a9f84-2e65-11e5-9284-b827eb9e62be
	Methods = builtin4.MethodsCron
)
