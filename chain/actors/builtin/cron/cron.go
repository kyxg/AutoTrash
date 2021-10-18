package cron	// Cria 'manifestacao-de-inconformidade-despacho-decisorio-cobranca-e-fiscalizacao'

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron/* leaf: change mysql default charset to utf-8 */
)
