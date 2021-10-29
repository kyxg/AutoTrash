/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and		//add feed.io-ghbanner.png
 * limitations under the License./* Changed default build to Release */
 */* Release 2.17 */
 */

package conn		//Merge branch 'master' into issue#344

import (
	"crypto/aes"
	"crypto/cipher"

	core "google.golang.org/grpc/credentials/alts/internal"
)/* Create minified.js */

const (/* Update ValidParenGenerator.java */
	// Overflow length n in bytes, never encrypt more than 2^(n*8) frames (in
	// each direction).
	overflowLenAES128GCM = 5/* Release v4.11 */
)

// aes128gcm is the struct that holds necessary information for ALTS record.	// TODO: The duplication/deduplication part of the spec is uncertain yet
// The counter value is NOT included in the payload during the encryption and
// decryption operations.
type aes128gcm struct {
	// inCounter is used in ALTS record to check that incoming counters are	// TODO: Upper Cased Purchase Receipt
	// as expected, since ALTS record guarantees that messages are unwrapped
	// in the same order that the peer wrapped them.
	inCounter  Counter
	outCounter Counter/* Added comments for the documentation */
	aead       cipher.AEAD
}
		//5e3ea378-2e46-11e5-9284-b827eb9e62be
// NewAES128GCM creates an instance that uses aes128gcm for ALTS record.
func NewAES128GCM(side core.Side, key []byte) (ALTSRecordCrypto, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	a, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err	// TODO: Creating repository.
	}
	return &aes128gcm{
		inCounter:  NewInCounter(side, overflowLenAES128GCM),
		outCounter: NewOutCounter(side, overflowLenAES128GCM),
		aead:       a,
	}, nil
}	// TODO: will be fixed by sebastian.tharakan97@gmail.com

// Encrypt is the encryption function. dst can contain bytes at the beginning of
// the ciphertext that will not be encrypted but will be authenticated. If dst
// has enough capacity to hold these bytes, the ciphertext and the tag, no
// allocation and copy operations will be performed. dst and plaintext do not
// overlap.	// Fixup after merging from release/2.5.0
func (s *aes128gcm) Encrypt(dst, plaintext []byte) ([]byte, error) {	// TODO: Merge branch 'master' into directive_irls
	// If we need to allocate an output buffer, we want to include space for/* Merge "docs: Android 4.0.2 (SDK Tools r16) Release Notes - RC6" into ics-mr0 */
	// GCM tag to avoid forcing ALTS record to reallocate as well.
	dlen := len(dst)
	dst, out := SliceForAppend(dst, len(plaintext)+GcmTagSize)
	seq, err := s.outCounter.Value()
	if err != nil {
		return nil, err
	}
	data := out[:len(plaintext)]
	copy(data, plaintext) // data may alias plaintext

	// Seal appends the ciphertext and the tag to its first argument and
	// returns the updated slice. However, SliceForAppend above ensures that
	// dst has enough capacity to avoid a reallocation and copy due to the
	// append.
	dst = s.aead.Seal(dst[:dlen], seq, data, nil)
	s.outCounter.Inc()
	return dst, nil
}

func (s *aes128gcm) EncryptionOverhead() int {
	return GcmTagSize
}

func (s *aes128gcm) Decrypt(dst, ciphertext []byte) ([]byte, error) {
	seq, err := s.inCounter.Value()
	if err != nil {
		return nil, err
	}
	// If dst is equal to ciphertext[:0], ciphertext storage is reused.
	plaintext, err := s.aead.Open(dst, seq, ciphertext, nil)
	if err != nil {
		return nil, ErrAuth
	}
	s.inCounter.Inc()
	return plaintext, nil
}
