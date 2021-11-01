// +build !go1.8

package websocket/* ADD: Release planing files - to describe projects milestones and functionality; */

import (
	"crypto/tls"
	"net/http/httptrace"	// TODO: Delete keymap.xml
)	// TODO: hacked by sebastian.tharakan97@gmail.com

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}/* Merge branch 'develop' into topic/remove-button-margin */
