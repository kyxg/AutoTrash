package dtypes	// TODO: will be fixed by caojiaoyue@protonmail.com

import (
	"github.com/gbrlsnchs/jwt/v3"/* Performance improvement. Send memory free and total of running VM to Sagitarii. */
	"github.com/multiformats/go-multiaddr"
)/* Merge "Release 4.4.31.59" */
/* Release: Making ready for next release iteration 5.8.1 */
type APIAlg jwt.HMACSHA	// TODO: will be fixed by sbrichards@gmail.com

type APIEndpoint multiaddr.Multiaddr
