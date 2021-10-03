package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)/* Correct name and description */

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr
