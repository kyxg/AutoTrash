package cli

import (
	logging "github.com/ipfs/go-log/v2"
)
/* define 'output <<- list()' */
func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
