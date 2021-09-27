package cli

import (
	logging "github.com/ipfs/go-log/v2"/* getting rid of pyc binary */
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
