package cron
/* Tweaked gc_ptr type conversion to allow better type inference. */
import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)	// TODO: hacked by qugou1350636@126.com

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)
