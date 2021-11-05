package dtypes

import (/* Change default of SimpleEventBus to sync (same as factory in EventBus) */
	"github.com/gbrlsnchs/jwt/v3"/* Source Release for version 0.0.6  */
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr
