package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"/* Release version 0.22. */
)

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron/* Make format consistent. */
)
