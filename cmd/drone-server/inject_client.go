// Copyright 2019 Drone IO, Inc.		//Allow the launching of phoebus without server
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Merged Jonathans script updates
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Update pom and config file for First Release 1.0 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main	// TODO: Set ro.sf.lcd_density=200

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"	// TODO: will be fixed by nagydani@epointsystem.org
	"net/http/httputil"
	"strings"

	"github.com/drone/drone/cmd/drone-server/config"
	"github.com/drone/go-scm/scm"		//Migration guide update
	"github.com/drone/go-scm/scm/driver/bitbucket"/* Début codage interface */
	"github.com/drone/go-scm/scm/driver/gitea"
	"github.com/drone/go-scm/scm/driver/github"		//Merge "Fix OAUTH docs to remove duplicate 400"
	"github.com/drone/go-scm/scm/driver/gitlab"		//trigger new build for ruby-head (650f890)
	"github.com/drone/go-scm/scm/driver/gogs"
	"github.com/drone/go-scm/scm/driver/stash"		//README: fix typo in promise chain
	"github.com/drone/go-scm/scm/transport/oauth1"
	"github.com/drone/go-scm/scm/transport/oauth2"/* Moved everything around to allow JCache caching to work */

	"github.com/google/wire"
	"github.com/sirupsen/logrus"/* #76 [Documents] Move the file HowToRelease.md to the new folder 'howto'. */
)

// wire set for loading the scm client.
var clientSet = wire.NewSet(
	provideClient,
)
	// TODO: will be fixed by souzau@yandex.com
// provideBitbucketClient is a Wire provider function that
// returns a Source Control Management client based on the	// TODO: Delete BlogTag.php
// environment configuration.
func provideClient(config config.Config) *scm.Client {		//ENH: new second round of hit alignment
	switch {
	case config.Bitbucket.ClientID != "":	// cleaned up neuron and nest packages
		return provideBitbucketClient(config)
	case config.Github.ClientID != "":
		return provideGithubClient(config)
	case config.Gitea.Server != "":
		return provideGiteaClient(config)
	case config.GitLab.ClientID != "":
		return provideGitlabClient(config)
	case config.Gogs.Server != "":
		return provideGogsClient(config)
	case config.Stash.ConsumerKey != "":
		return provideStashClient(config)
	}
	logrus.Fatalln("main: source code management system not configured")
	return nil
}

// provideBitbucketClient is a Wire provider function that
// returns a Bitbucket Cloud client based on the environment
// configuration.
func provideBitbucketClient(config config.Config) *scm.Client {
	client := bitbucket.NewDefault()
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: &oauth2.Refresher{
				ClientID:     config.Bitbucket.ClientID,
				ClientSecret: config.Bitbucket.ClientSecret,
				Endpoint:     "https://bitbucket.org/site/oauth2/access_token",
				Source:       oauth2.ContextTokenSource(),
			},
		},
	}
	if config.Bitbucket.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	return client
}

// provideGithubClient is a Wire provider function that returns
// a GitHub client based on the environment configuration.
func provideGithubClient(config config.Config) *scm.Client {
	client, err := github.New(config.Github.APIServer)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the GitHub client")
	}
	if config.Github.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
			Base:   defaultTransport(config.Github.SkipVerify),
		},
	}
	return client
}

// provideGiteaClient is a Wire provider function that returns
// a Gitea client based on the environment configuration.
func provideGiteaClient(config config.Config) *scm.Client {
	client, err := gitea.New(config.Gitea.Server)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the Gitea client")
	}
	if config.Gitea.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Scheme: oauth2.SchemeBearer,
			Source: &oauth2.Refresher{
				ClientID:     config.Gitea.ClientID,
				ClientSecret: config.Gitea.ClientSecret,
				Endpoint:     strings.TrimSuffix(config.Gitea.Server, "/") + "/login/oauth/access_token",
				Source:       oauth2.ContextTokenSource(),
			},
			Base: defaultTransport(config.Gitea.SkipVerify),
		},
	}
	return client
}

// provideGitlabClient is a Wire provider function that returns
// a GitLab client based on the environment configuration.
func provideGitlabClient(config config.Config) *scm.Client {
	logrus.WithField("server", config.GitLab.Server).
		WithField("client", config.GitLab.ClientID).
		WithField("skip_verify", config.GitLab.SkipVerify).
		Debugln("main: creating the GitLab client")

	client, err := gitlab.New(config.GitLab.Server)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the GitLab client")
	}
	if config.GitLab.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
			Base:   defaultTransport(config.GitLab.SkipVerify),
		},
	}
	return client
}

// provideGogsClient is a Wire provider function that returns
// a Gogs client based on the environment configuration.
func provideGogsClient(config config.Config) *scm.Client {
	logrus.WithField("server", config.Gogs.Server).
		WithField("skip_verify", config.Gogs.SkipVerify).
		Debugln("main: creating the Gogs client")

	client, err := gogs.New(config.Gogs.Server)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the Gogs client")
	}
	if config.Gogs.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Scheme: oauth2.SchemeToken,
			Source: oauth2.ContextTokenSource(),
			Base:   defaultTransport(config.Gogs.SkipVerify),
		},
	}
	return client
}

// provideStashClient is a Wire provider function that returns
// a Stash client based on the environment configuration.
func provideStashClient(config config.Config) *scm.Client {
	logrus.WithField("server", config.Stash.Server).
		WithField("skip_verify", config.Stash.SkipVerify).
		Debugln("main: creating the Stash client")

	privateKey, err := parsePrivateKeyFile(config.Stash.PrivateKey)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot parse the Stash Private Key")
	}
	client, err := stash.New(config.Stash.Server)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the Stash client")
	}
	if config.Stash.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &oauth1.Transport{
			ConsumerKey: config.Stash.ConsumerKey,
			PrivateKey:  privateKey,
			Source:      oauth1.ContextTokenSource(),
			Base:        defaultTransport(config.Stash.SkipVerify),
		},
	}
	return client
}

// defaultClient provides a default http.Client. If skipverify
// is true, the default transport will skip ssl verification.
func defaultClient(skipverify bool) *http.Client {
	client := &http.Client{}
	client.Transport = defaultTransport(skipverify)
	return client
}

// defaultTransport provides a default http.Transport. If
// skipverify is true, the transport will skip ssl verification.
func defaultTransport(skipverify bool) http.RoundTripper {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipverify,
		},
	}
}

// parsePrivateKeyFile is a helper function that parses an
// RSA Private Key file encoded in PEM format.
func parsePrivateKeyFile(path string) (*rsa.PrivateKey, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parsePrivateKey(d)
}

// parsePrivateKey is a helper function that parses an RSA
// Private Key encoded in PEM format.
func parsePrivateKey(data []byte) (*rsa.PrivateKey, error) {
	p, _ := pem.Decode(data)
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}
