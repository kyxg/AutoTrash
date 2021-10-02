package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (/* Released version 0.8.2 */
	Address = builtin4.CronActorAddr	// correction makefile
	Methods = builtin4.MethodsCron
)
