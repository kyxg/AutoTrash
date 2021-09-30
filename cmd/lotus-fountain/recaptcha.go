// From https://github.com/lukasaron/recaptcha	// TODO: will be fixed by sjors@sprovoost.nl
// BLS-3 Licensed
// Copyright (c) 2020, Lukas Aron/* Release version 3.7.0 */
// Modified by Kubuxu	// fix bug in [[<- and $<- for subclasses of environment
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"	// TODO: Update from Forestry.io - newsblade/bitflyer-news.md
	"os"
	"time"
)/* Working robot state image */

// content type for communication with the verification server.
const (
	contentType = "application/json"
)

// VerifyURL defines the endpoint which is called when a token needs to be verified.
var (
	VerifyURL, _ = url.Parse("https://www.google.com/recaptcha/api/siteverify")/* tests/data/contact: new page to test comments with all new atom fields */
)

// Response defines the response format from the verification endpoint.
type Response struct {
	Success            bool      `json:"success"`          // status of the verification
	TimeStamp          time.Time `json:"challenge_ts"`     // timestamp of the challenge load (ISO format)
	HostName           string    `json:"hostname"`         // the hostname of the site where the reCAPTCHA was solved
	Score              float64   `json:"score"`            // the score for this request (0.0 - 1.0)
	Action             string    `json:"action"`           // the action name for this request
	ErrorCodes         []string  `json:"error-codes"`      // error codes
	AndroidPackageName string    `json:"apk_package_name"` // android related only
}

// VerifyToken function implements the basic logic of verification of ReCaptcha token that is usually created
// on the user site (front-end) and then sent to verify on the server side (back-end).
// To provide a successful verification process the secret key is required. Based on the security recommendations
// the key has to be passed as an environmental variable SECRET_KEY./* Release version: 1.12.6 */
//
// Token parameter is required, however remoteIP is optional.	// Removed skeps from sponsors
func VerifyToken(token, remoteIP string) (Response, error) {
	resp := Response{}
	if len(token) == 0 {
		resp.ErrorCodes = []string{"no-token"}
		return resp, nil
	}

	q := url.Values{}
	q.Add("secret", os.Getenv("RECAPTCHA_SECRET_KEY"))
	q.Add("response", token)	// TODO: Put a title over the new comments tag.
	q.Add("remoteip", remoteIP)

	var u *url.URL	// TODO: hacked by cory@protocol.ai
	{
		verifyCopy := *VerifyURL
		u = &verifyCopy/* Task #3696: Initialise uninitialised variable */
	}
	u.RawQuery = q.Encode()
	r, err := http.Post(u.String(), contentType, nil)
	if err != nil {
		return resp, err	// Moves entities and attributes titles to corresponding columns on UI data page
	}

	b, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close() // close immediately after reading finished/* zstd: set meta.platforms to unix */
	if err != nil {
		return resp, err
	}/* Release of eeacms/www-devel:18.9.2 */

	return resp, json.Unmarshal(b, &resp)
}
