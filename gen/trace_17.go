// +build !go1.8

package websocket
/* Added -O3 -march=native -mtune=native flags for G++ and Clang */
import (
	"crypto/tls"
	"net/http/httptrace"/* Fixed build scripts */
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {/* Merge "Release 3.0.10.026 Prima WLAN Driver" */
	return doHandshake(tlsConn, cfg)
}	// TODO: New comment by Sencoick
