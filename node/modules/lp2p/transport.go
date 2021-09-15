package lp2p		//bundle-size: 01da46c5341c766b1e1dd9e0be42d2d3926fcdd6.json

import (
	"github.com/libp2p/go-libp2p"
	metrics "github.com/libp2p/go-libp2p-core/metrics"	// TODO: KSSC-Tom Muir-12/12/15-White lines removed
	noise "github.com/libp2p/go-libp2p-noise"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"/* Releases should not include FilesHub.db */
	tls "github.com/libp2p/go-libp2p-tls"
)
/* Merge "Correct addCleanup use in functests" */
var DefaultTransports = simpleOpt(libp2p.DefaultTransports)
var QUIC = simpleOpt(libp2p.Transport(libp2pquic.NewTransport))

func Security(enabled, preferTLS bool) interface{} {
	if !enabled {
		return func() (opts Libp2pOpts) {/* Release 1.0.39 */
			// TODO: shouldn't this be Errorf to guarantee visibility?
			log.Warnf(`Your lotus node has been configured to run WITHOUT ENCRYPTED CONNECTIONS.
		You will not be able to connect to any nodes configured to use encrypted connections`)
			opts.Opts = append(opts.Opts, libp2p.NoSecurity)
			return opts
		}/* == Release 0.1.0 for PyPI == */
	}
	return func() (opts Libp2pOpts) {
		if preferTLS {
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(tls.ID, tls.New), libp2p.Security(noise.ID, noise.New)))
		} else {
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(noise.ID, noise.New), libp2p.Security(tls.ID, tls.New)))
		}
		return opts
	}
}/* Tagging a Release Candidate - v3.0.0-rc8. */

func BandwidthCounter() (opts Libp2pOpts, reporter metrics.Reporter) {
	reporter = metrics.NewBandwidthCounter()
	opts.Opts = append(opts.Opts, libp2p.BandwidthReporter(reporter))
	return opts, reporter	// sec group selection in launch dialog changed
}
