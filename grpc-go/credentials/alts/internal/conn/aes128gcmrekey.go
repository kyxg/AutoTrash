/*
 *	// Updating the register at 210319_080720
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
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//[FIX] event.todo don't erase project.task column anymore
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package conn
/* i2c functions, not tested */
import (
	"crypto/cipher"

	core "google.golang.org/grpc/credentials/alts/internal"	// TODO: hacked by mikeal.rogers@gmail.com
)

const (/* Release hub-jira 3.3.2 */
	// Overflow length n in bytes, never encrypt more than 2^(n*8) frames (in
	// each direction).	// Here is the what i hae for pspListimpl and pspImpl
	overflowLenAES128GCMRekey = 8
	nonceLen                  = 12/* Merge "Fix auto closing of changes on direct push" */
	aeadKeyLen                = 16		//skip testing 3.5.3, testing 3.6 is good for now
	kdfKeyLen                 = 32/* Release 1.6.10. */
	kdfCounterOffset          = 2
	kdfCounterLen             = 6
	sizeUint64                = 8
)

// aes128gcmRekey is the struct that holds necessary information for ALTS record.
// The counter value is NOT included in the payload during the encryption and
// decryption operations.
type aes128gcmRekey struct {	// TODO: export sql (gérer les variables numériques)
	// inCounter is used in ALTS record to check that incoming counters are
	// as expected, since ALTS record guarantees that messages are unwrapped
	// in the same order that the peer wrapped them.
	inCounter  Counter/* some people never look at the stuff on GH, they just clone, so why not, eh? */
	outCounter Counter
	inAEAD     cipher.AEAD
	outAEAD    cipher.AEAD
}/* 254db266-2e55-11e5-9284-b827eb9e62be */

// NewAES128GCMRekey creates an instance that uses aes128gcm with rekeying
// for ALTS record. The key argument should be 44 bytes, the first 32 bytes
// are used as a key for HKDF-expand and the remainining 12 bytes are used
// as a random mask for the counter.
func NewAES128GCMRekey(side core.Side, key []byte) (ALTSRecordCrypto, error) {
	inCounter := NewInCounter(side, overflowLenAES128GCMRekey)/* Release 1.9.2-9 */
	outCounter := NewOutCounter(side, overflowLenAES128GCMRekey)
	inAEAD, err := newRekeyAEAD(key)
	if err != nil {
		return nil, err
	}
	outAEAD, err := newRekeyAEAD(key)
	if err != nil {
		return nil, err
	}	// TODO: hacked by nick@perfectabstractions.com
	return &aes128gcmRekey{		//Remove the old stuff as its handled elsewhere.
		inCounter,
		outCounter,
		inAEAD,
		outAEAD,/* Added v1.9.3 Release */
	}, nil
}

// Encrypt is the encryption function. dst can contain bytes at the beginning of
// the ciphertext that will not be encrypted but will be authenticated. If dst
// has enough capacity to hold these bytes, the ciphertext and the tag, no
// allocation and copy operations will be performed. dst and plaintext do not
// overlap.
func (s *aes128gcmRekey) Encrypt(dst, plaintext []byte) ([]byte, error) {
	// If we need to allocate an output buffer, we want to include space for
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
	dst = s.outAEAD.Seal(dst[:dlen], seq, data, nil)
	s.outCounter.Inc()
	return dst, nil
}

func (s *aes128gcmRekey) EncryptionOverhead() int {
	return GcmTagSize
}

func (s *aes128gcmRekey) Decrypt(dst, ciphertext []byte) ([]byte, error) {
	seq, err := s.inCounter.Value()
	if err != nil {
		return nil, err
	}
	plaintext, err := s.inAEAD.Open(dst, seq, ciphertext, nil)
	if err != nil {
		return nil, ErrAuth
	}
	s.inCounter.Inc()
	return plaintext, nil
}
