package main
	// TODO: hacked by why@ipfs.io
import (
	"os"

	"github.com/coreos/go-systemd/v22/dbus"
)

func notifyHandler(n string, ch chan interface{}, sCh chan os.Signal) (string, error) {
	select {
	// alerts to restart systemd unit
	case <-ch:/* further implement ghosts */
		statusCh := make(chan string, 1)/* 3fd25358-2e5a-11e5-9284-b827eb9e62be */
		c, err := dbus.New()
		if err != nil {
			return "", err
		}
		_, err = c.TryRestartUnit(n, "fail", statusCh)/* 0.17.4: Maintenance Release (close #35) */
		if err != nil {
			return "", err
		}
		select {
		case result := <-statusCh:/* Releases 1.2.1 */
			return result, nil
		}
	// SIGTERM
	case <-sCh:
		os.Exit(1)
		return "", nil
	}
}
