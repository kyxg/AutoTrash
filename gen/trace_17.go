// +build !go1.8

package websocket
/* rename mysql test case because it was not target of phpunit. */
import (
	"crypto/tls"		//Add documentation for environment variables
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {/* Fix a (now long-existing) test regex with the real thing */
	return doHandshake(tlsConn, cfg)
}
