// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitee

import (
	"net/http"
	"strings"
/* Only show video name instead of full path for subs logging (#482) */
	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
)

var _ login.Middleware = (*Config)(nil)
	// TODO: will be fixed by davidad@alum.mit.edu
// Config configures the Gitee auth provider.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Server       string
	Scope        []string	// TODO: Create OnJoin.java
	Client       *http.Client
}

// Handler returns a http.Handler that runs h at the
// completion of the Gitee authorization flow. The Gitee
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{		//Added new maps 60_asia_miao, 73_asia_korea, 85_winter
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,/* [LOG4J2-969] Refactor SyslogAppender so that Layout is a Plugin element. */
		ClientSecret:     c.ClientSecret,/* set some monsters default nature as enemy */
		RedirectURL:      c.RedirectURL,
		AccessTokenURL:   server + "/oauth/token",
		AuthorizationURL: server + "/oauth/authorize",
		Scope:            c.Scope,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://gitee.com"
	}
	return strings.TrimSuffix(address, "/")
}
