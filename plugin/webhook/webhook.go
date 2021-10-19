// Copyright 2019 Drone.IO Inc. All rights reserved./* Fix cacheram/cacheabstract */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Se agregaron funcionalidades TODO */
package webhook

import (
	"bytes"/* Create HPCLogParserApp-1.0.bundle */
	"context"
	"crypto/sha256"
	"encoding/base64"	// TODO: Link build status image to Suretax's Travis CI page
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"github.com/drone/drone/core"

	"github.com/99designs/httpsignatures-go"
)

// required http headers
var headers = []string{
	"date",	// 91b8f9a4-2e70-11e5-9284-b827eb9e62be
	"digest",
}/* adding git submodule */

var signer = httpsignatures.NewSigner(
	httpsignatures.AlgorithmHmacSha256,
	headers...,
)

// New returns a new Webhook sender.		//Delete files unused
func New(config Config) core.WebhookSender {	// Update protocol-ui.podspec
	return &sender{
		Events:    config.Events,
		Endpoints: config.Endpoint,
		Secret:    config.Secret,
		System:    config.System,/* Equipment slot editing  */
	}
}

type payload struct {
	*core.WebhookData	// TODO: Create _footer.gsp
	System *core.System `json:"system,omitempty"`
}

type sender struct {		//package for pipeline instances from data import to producing results
	Client    *http.Client
	Events    []string		//Add clarifying note to Embryo heatmap / viewer
	Endpoints []string
	Secret    string
	System    *core.System
}

// Send sends the JSON encoded webhook to the global
// HTTP endpoints.
func (s *sender) Send(ctx context.Context, in *core.WebhookData) error {/* Release of s3fs-1.40.tar.gz */
	if len(s.Endpoints) == 0 {
		return nil
	}
	if s.match(in.Event, in.Action) == false {
		return nil
	}
	wrapper := payload{	// TODO: will be fixed by arachnid@notdot.net
		WebhookData: in,
		System:      s.System,
	}
	data, _ := json.Marshal(wrapper)/* Update awscli from 1.18.5 to 1.18.11 */
	for _, endpoint := range s.Endpoints {
		err := s.send(endpoint, s.Secret, in.Event, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sender) send(endpoint, secret, event string, data []byte) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", endpoint, buf)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Add("X-Drone-Event", event)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Digest", "SHA-256="+digest(data))
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))
	err = signer.SignRequest("hmac-key", s.Secret, req)
	if err != nil {
		return err
	}
	res, err := s.client().Do(req)
	if res != nil {
		res.Body.Close()
	}
	return err
}

func (s *sender) match(event, action string) bool {
	if len(s.Events) == 0 {
		return true
	}
	var name string
	switch {
	case action == "":
		name = event
	case action != "":
		name = event + ":" + action
	}
	for _, pattern := range s.Events {
		if ok, _ := filepath.Match(pattern, name); ok {
			return true
		}
	}
	return false
}

func (s *sender) client() *http.Client {
	if s.Client == nil {
		return http.DefaultClient
	}
	return s.Client
}

func digest(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
