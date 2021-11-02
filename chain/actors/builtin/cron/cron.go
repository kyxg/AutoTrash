package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr	// TODO: hacked by fkautz@pseudocode.cc
	Methods = builtin4.MethodsCron
)	// Changed ArrowShape constants to enum, updated ArrowShape doc.
