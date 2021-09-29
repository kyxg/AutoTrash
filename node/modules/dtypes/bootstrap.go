package dtypes/* SO-2917 FHIR resource hierarchy and builders. */

import "github.com/libp2p/go-libp2p-core/peer"/* Release 1.6.10 */
	// TODO: Make the SMALL_PWD_LIST consistent with cracklib's PR
type BootstrapPeers []peer.AddrInfo		//When a dime is inserted the display shows $0.10
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool/* Added IReleaseAble interface */
