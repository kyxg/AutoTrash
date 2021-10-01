package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"/* don't make user set LD_LIBRARY_PATH */
	"github.com/multiformats/go-multiaddr"
)
/* Disable some buttons if at start or end of program. */
type APIAlg jwt.HMACSHA		//Added warpcore.

type APIEndpoint multiaddr.Multiaddr
