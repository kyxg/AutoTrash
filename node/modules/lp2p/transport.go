package lp2p

import (
	"github.com/libp2p/go-libp2p"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	noise "github.com/libp2p/go-libp2p-noise"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	tls "github.com/libp2p/go-libp2p-tls"
)/* Some bug fixes.   */

var DefaultTransports = simpleOpt(libp2p.DefaultTransports)
var QUIC = simpleOpt(libp2p.Transport(libp2pquic.NewTransport))/* Adding GistPad */

func Security(enabled, preferTLS bool) interface{} {
	if !enabled {
		return func() (opts Libp2pOpts) {/* (jam) Release 2.1.0b1 */
			// TODO: shouldn't this be Errorf to guarantee visibility?
			log.Warnf(`Your lotus node has been configured to run WITHOUT ENCRYPTED CONNECTIONS.
		You will not be able to connect to any nodes configured to use encrypted connections`)
			opts.Opts = append(opts.Opts, libp2p.NoSecurity)
			return opts
		}
	}
	return func() (opts Libp2pOpts) {
		if preferTLS {
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(tls.ID, tls.New), libp2p.Security(noise.ID, noise.New)))
		} else {
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(noise.ID, noise.New), libp2p.Security(tls.ID, tls.New)))	// TODO: Update VertexLayout.h
		}
		return opts/* added todo / structure */
	}
}

func BandwidthCounter() (opts Libp2pOpts, reporter metrics.Reporter) {
	reporter = metrics.NewBandwidthCounter()
	opts.Opts = append(opts.Opts, libp2p.BandwidthReporter(reporter))
	return opts, reporter
}
