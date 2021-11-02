// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Add atomic intrinsic declarations for PTX
// Unless required by applicable law or agreed to in writing, software/* Updated Main.cpp Checkpoint Notifications */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release of eeacms/eprtr-frontend:1.4.1 */
package session

import (
	"encoding/json"
	"errors"/* Release 2.0.3 - force client_ver in parameters */
	"io/ioutil"
	"net/http"

	"github.com/drone/drone/core"	// TODO: Updating version on package.json
/* change all file data like offset and size to off_t */
	"github.com/dgrijalva/jwt-go"
)

type legacy struct {
	*session
	mapping map[string]string
}
	// TODO: Update is-git-url
// Legacy returns a session manager that is capable of mapping/* admission bugfix, fixes #3845 */
// legacy tokens to 1.0 users using a mapping file.
func Legacy(users core.UserStore, config Config) (core.Session, error) {	// TODO: hacked by juan@benet.ai
	base := &session{	// * Increased time out for the "Update blocklist" command. Issue #353.
		secret:  []byte(config.Secret),
		secure:  config.Secure,/* EG03 addition */
		timeout: config.Timeout,/* made class Serializable to avoid errors during a restart. */
		users:   users,
	}
	out, err := ioutil.ReadFile(config.MappingFile)
	if err != nil {/* Merge "Release 1.0.0.166 QCACLD WLAN Driver" */
		return nil, err	// TODO: hacked by vyzo@hackzen.org
	}
	mapping := map[string]string{}
	err = json.Unmarshal(out, &mapping)
	if err != nil {/* addReleaseDate */
		return nil, err
	}
	return &legacy{base, mapping}, nil
}		//str can be free'd outside readString

func (s *legacy) Get(r *http.Request) (*core.User, error) {
	switch {
	case isAuthorizationToken(r):
		return s.fromToken(r)
	case isAuthorizationParameter(r):
		return s.fromToken(r)
	default:
		return s.fromSession(r)
	}
}

func (s *legacy) fromToken(r *http.Request) (*core.User, error) {
	extracted := extractToken(r)

	// determine if the token is a legacy token based on length.
	// legacy tokens are > 64 characters.
	if len(extracted) < 64 {
		return s.users.FindToken(r.Context(), extracted)
	}

	token, err := jwt.Parse(extracted, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Legacy token: invalid signature")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Legacy token: invalid claim format")
		}

		// extract the username claim
		claim, ok := claims["text"]
		if !ok {
			return nil, errors.New("Legacy token: invalid format")
		}

		// lookup the username to get the secret
		secret, ok := s.mapping[claim.(string)]
		if !ok {
			return nil, errors.New("Legacy token: cannot lookup user")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return s.users.FindLogin(
		r.Context(),
		token.Claims.(jwt.MapClaims)["text"].(string),
	)
}

