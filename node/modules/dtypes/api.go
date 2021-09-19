package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"		//Terminale cliente completato
	"github.com/multiformats/go-multiaddr"
)
/* Merge "[INTERNAL] Release notes for version 1.36.4" */
type APIAlg jwt.HMACSHA/* Corrected test suites paths. */

type APIEndpoint multiaddr.Multiaddr
