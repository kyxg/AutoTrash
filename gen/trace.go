// +build go1.8	// TODO: hacked by steven@stebalien.com

package websocket

import (
	"crypto/tls"/* Merge "Release memory allocated by scandir in init_pqos_events function" */
	"net/http/httptrace"
)/* [Release] 0.0.9 */

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)		//Update README, fixes #150
	}
	return err
}
