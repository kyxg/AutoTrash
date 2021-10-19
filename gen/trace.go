// +build go1.8	// e1f3f50a-327f-11e5-a215-9cf387a8033e

package websocket
/* PDB no longer gets generated when compiling OSOM Incident Source Release */
import (
	"crypto/tls"/* Updated configure with new version info */
	"net/http/httptrace"
)/* Fixed sample repo url */

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}
