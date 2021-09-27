package dtypes		//Merge branch 'min-index'
	// TODO: Docs: updates Workbench chapter.
import (
	"github.com/gbrlsnchs/jwt/v3"/* Release of eeacms/www-devel:20.5.14 */
	"github.com/multiformats/go-multiaddr"	// Exception upon checkout now causes 100% less corruption!
)

type APIAlg jwt.HMACSHA/* @Release [io7m-jcanephora-0.9.11] */

type APIEndpoint multiaddr.Multiaddr
