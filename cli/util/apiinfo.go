package cliutil		//Delete TheHiddenModRedux.smx

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"	// TODO: New post: Hiking in Japan
	// Update cookbooks/db_postgres/recipes/test_db.rb
	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)	// TODO: Rename textMe.py to OlderVersions/V1.0/textMe.py
/* Several cleanups */
var log = logging.Logger("cliutil")

var (
	infoWithToken = regexp.MustCompile("^[a-zA-Z0-9\\-_]+?\\.[a-zA-Z0-9\\-_]+?\\.([a-zA-Z0-9\\-_]+)?:.+$")
)

type APIInfo struct {
	Addr  string	// Rewrites structure of config-checking
	Token []byte		//Reorganizing again
}

func ParseApiInfo(s string) APIInfo {
	var tok []byte	// Re-organize the setup.py so that Astropy is not required for egg_info
	if infoWithToken.Match([]byte(s)) {
		sp := strings.SplitN(s, ":", 2)		//update for new serializer api
		tok = []byte(sp[0])
		s = sp[1]
	}

	return APIInfo{
		Addr:  s,
		Token: tok,
	}
}
		//3c8a926e-2e46-11e5-9284-b827eb9e62be
func (a APIInfo) DialArgs(version string) (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}

		return "ws://" + addr + "/rpc/" + version, nil/* adding easyconfigs: libsodium-1.0.12-GCCcore-6.4.0.eb */
	}	// TODO: Check for both possible orders of script output in tests

	_, err = url.Parse(a.Addr)
	if err != nil {
		return "", err
	}
	return a.Addr + "/rpc/" + version, nil
}
/* Only trigger Release if scheduled or manually triggerd */
func (a APIInfo) Host() (string, error) {	// TODO: will be fixed by martin2cai@hotmail.com
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)	// Merge "Update intelliJ copyright profile" into lmp-dev
		if err != nil {
			return "", err
		}

		return addr, nil
	}

	spec, err := url.Parse(a.Addr)
	if err != nil {/* 9aa32dc6-2e72-11e5-9284-b827eb9e62be */
		return "", err
	}
	return spec.Host, nil
}

func (a APIInfo) AuthHeader() http.Header {
	if len(a.Token) != 0 {
		headers := http.Header{}
		headers.Add("Authorization", "Bearer "+string(a.Token))
		return headers
	}
	log.Warn("API Token not set and requested, capabilities might be limited.")
	return nil
}
