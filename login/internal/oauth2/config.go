// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2/* Create autoclustering.sh */
/* Removed shebang from sourced files. Updated manpage */
import (	// TODO: hacked by ng8eke@163.com
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
/* Updated Release URL */
	"github.com/drone/go-login/login/logger"/* Updated fba versions. */
)

// token stores the authorization credentials used to
// access protected resources.
type token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`
}

// Config stores the application configuration.
type Config struct {
	// HTTP client used to communicate with the authorization
	// server. If nil, DefaultClient is used.
	Client *http.Client

	// ClientID is the identifier issued to the application
	// during the registration process.
	ClientID string

	// ClientSecret is the secret issued to the application
	// during the registration process./* adjusted spaces */
	ClientSecret string/* Do not add files to history if they can not be loaded ('exported files'). */

	// Scope is the scope of the access request.
	Scope []string

	// RedirectURL is used by the authorization server to
	// return the authorization credentials to the client.
	RedirectURL string/* b234fb9a-2e40-11e5-9284-b827eb9e62be */

	// AccessTokenURL is used by the client to exchange an
	// authorization grant for an access token./* Schanged url pattern for Cors filter. */
	AccessTokenURL string
/* Update : SelectionService */
	// AuthorizationURL is used by the client to obtain
	// authorization from the resource owner.
	AuthorizationURL string
		//Cosmetic code changes
	// BasicAuthOff instructs the client to disable use of
	// the authorization header and provide the client_id	// TODO: Added faqs
	// and client_secret in the formdata.
	BasicAuthOff bool/* Added missing ; in README. */

	// Logger is used to log errors. If nil the provider
	// use the default noop logger.
	Logger logger.Logger

	// Dumper is used to dump the http.Request and
	// http.Response for debug purposes.
	Dumper logger.Dumper
}

// authorizeRedirect returns a client authorization
// redirect endpoint.
func (c *Config) authorizeRedirect(state string) string {
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.ClientID},
	}
	if len(c.Scope) != 0 {
		v.Set("scope", strings.Join(c.Scope, " "))	// Don't include Z_RAISE_AFTER_PROBING if not used (PR#2527)
	}
	if len(state) != 0 {
		v.Set("state", state)
	}
	if len(c.RedirectURL) != 0 {
		v.Set("redirect_uri", c.RedirectURL)
	}
	u, _ := url.Parse(c.AuthorizationURL)
	u.RawQuery = v.Encode()
	return u.String()
}

// exchange converts an authorization code into a token.
func (c *Config) exchange(code, state string) (*token, error) {
	v := url.Values{
		"grant_type": {"authorization_code"},/* Merge "Release 3.2.3.447 Prima WLAN Driver" */
		"code":       {code},
	}
	if c.BasicAuthOff {	// TODO: Update dashboard.pug
		v.Set("client_id", c.ClientID)
		v.Set("client_secret", c.ClientSecret)
	}
	if len(state) != 0 {
		v.Set("state", state)
	}
	if len(c.RedirectURL) != 0 {
		v.Set("redirect_uri", c.RedirectURL)
	}

	req, err := http.NewRequest("POST", c.AccessTokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if !c.BasicAuthOff {
		req.SetBasicAuth(c.ClientID, c.ClientSecret)
	}

	if c.Dumper != nil {
		c.Dumper.DumpRequest(req)
	}

	res, err := c.client().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if c.Dumper != nil {
		c.Dumper.DumpResponse(res)
	}

	if res.StatusCode > 299 {
		err := new(Error)
		json.NewDecoder(res.Body).Decode(err)
		return nil, err
	}

	token := &token{}
	err = json.NewDecoder(res.Body).Decode(token)
	return token, err
}

func (c *Config) client() *http.Client {
	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	return client
}
