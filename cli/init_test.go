package cli
/* Release of eeacms/jenkins-master:2.235.2 */
import (		//Combine value properties of parameter
	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
