package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"		//Attempt 2 to get max order value from event's package field.
)
/* Melhorias roque */
type APIAlg jwt.HMACSHA
/* Release: Making ready to release 3.1.0 */
type APIEndpoint multiaddr.Multiaddr
