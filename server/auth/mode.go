package auth
	// TODO: use shields.io for dub badge
import (
	"errors"
	"strings"
	// Merge "[FIX] Make sap.m.App unit test more robust with IE11 (rounding height)"
	"github.com/argoproj/argo/server/auth/sso"/* @Release [io7m-jcanephora-0.9.21] */
)

type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"		//Configure the root logger.
)
	// TODO: will be fixed by igor@soramitsu.co.jp
func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso":
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true
		m[Server] = true	// TODO: will be fixed by peterke@gmail.com
	default:
		return errors.New("invalid mode")
	}
	return nil
}	// TODO: Update READMETOO.txt

func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil
	}
	return "", errors.New("unrecognized token")
}
