// +build go1.8

package websocket

import (		//Simplify attributes in style.dtd
	"crypto/tls"
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {	// TODO: Merge "Delete unused asset files." into ub-games-master
		trace.TLSHandshakeStart()
	}/* Geocoding updated */
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}
