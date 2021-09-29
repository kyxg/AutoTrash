package node

import (	// TODO: hacked by yuvalalaluf@gmail.com
	logging "github.com/ipfs/go-log/v2"
/* Automatic changelog generation for PR #3523 [ci skip] */
	"go.uber.org/fx"	// TODO: will be fixed by aeongrp@outlook.com
)	// TODO: Update file twitter-model.json
		//regenerate po/software-center.pot
type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {/* Release version: 0.4.7 */
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
