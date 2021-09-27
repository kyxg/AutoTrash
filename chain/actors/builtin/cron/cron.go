package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)
/* Improve Im() */
var (	// Re #1462: fixed libresample.dylib.1 installation issue
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)	// TODO: will be fixed by igor@soramitsu.co.jp
