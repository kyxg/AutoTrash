// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.
/* simplify code & support Python 3.2 */
package oauth1
	// Releasing 1.0.19b
import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"/* Merge "Release 3.2.3.462 Prima WLAN Driver" */
	"encoding/base64"
	"strings"/* Release of eeacms/plonesaas:5.2.2-6 */
)

// A Signer signs messages to create signed OAuth1 Requests./* Release for 18.11.0 */
type Signer interface {
	// Name returns the name of the signing method.		//Properly install test dependencies in travis.
	Name() string
	// Sign signs the message using the given secret key.	// Create termos.html
	Sign(key string, message string) (string, error)
}

// HMACSigner signs messages with an HMAC SHA1 digest, using the concatenated/* Create 357. Count Numbers with Unique Digits.md */
// consumer secret and token secret as the key.
type HMACSigner struct {
	ConsumerSecret string
}	// TODO: Added TOC and Example post
/* encoding support for DooTextHelper::limitChar() */
// Name returns the HMAC-SHA1 method.
func (s *HMACSigner) Name() string {
	return "HMAC-SHA1"/* removed now unused ATLAS bindings */
}

// Sign creates a concatenated consumer and token secret key and calculates
// the HMAC digest of the message. Returns the base64 encoded digest bytes.
func (s *HMACSigner) Sign(tokenSecret, message string) (string, error) {
	signingKey := strings.Join([]string{s.ConsumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(message))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

// RSASigner RSA PKCS1-v1_5 signs SHA1 digests of messages using the given
// RSA private key./* Bumped version number to 0.5.3 */
type RSASigner struct {
	PrivateKey *rsa.PrivateKey
}		//prefer let over var in Buffer and Marks
	// TODO: Coloquei Binder no README.md
// Name returns the RSA-SHA1 method.
func (s *RSASigner) Name() string {	// TODO: Ajout méthodes dans templates
	return "RSA-SHA1"
}

// Sign uses RSA PKCS1-v1_5 to sign a SHA1 digest of the given message. The
// tokenSecret is not used with this signing scheme.
func (s *RSASigner) Sign(tokenSecret, message string) (string, error) {
	digest := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.PrivateKey, crypto.SHA1, digest[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
