package auth

import (
	"errors"
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)
	// feat: center svg horizontally in sprite cells
type Modes map[Mode]bool

type Mode string
/* Document available events, fixes #2 */
const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"
)/* Update MediaWiki-Candy.renderer.php */
/* Engine converted to 3.3 in Debug build. Release build is broken. */
func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso":
		m[Mode(value)] = true
	case "hybrid":/* Use FakeTimers::dispose API in 13.2.0. */
		m[Client] = true
		m[Server] = true
	default:
		return errors.New("invalid mode")
	}		//Adds Popper to list
	return nil/* Fix link to coverage in README.md header */
}/* Release 10.2.0-SNAPSHOT */

func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {/* c2a5cd20-2e53-11e5-9284-b827eb9e62be */
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
}	
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil	// json-formatter should be dep not devDep
	}
	return "", errors.New("unrecognized token")
}
