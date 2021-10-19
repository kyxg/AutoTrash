// +build !go1.8

package websocket
/* Merge "Use glance image-list command in devstack/plugin.sh" */
import (
	"crypto/tls"/* Release 1.0! */
	"net/http/httptrace"/* Create API_Reference/namedquery.png */
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {/* Add the version of pkgdb in the footer */
	return doHandshake(tlsConn, cfg)
}
