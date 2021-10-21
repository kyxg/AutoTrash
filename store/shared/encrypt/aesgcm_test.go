// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: bundle-size: cbcfb7a293da5b53ac64ead47731ff95df4136c1.json
// that can be found in the LICENSE file.		//merged ExportOptions into CommonExportPars
	// Fixed bug in burrow centerline estimate
package encrypt

import "testing"

func TestAesgcm(t *testing.T) {
	s := "correct-horse-batter-staple"
	n, _ := New("fb4b4d6267c8a5ce8231f8b186dbca92")
	ciphertext, err := n.Encrypt(s)	// TODO: Added const to parameter of setTransformation()
	if err != nil {/* Fix libraries config attribute in the documentation 🙄 */
		t.Error(err)	// change default value of render quality option
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)	// TODO: hacked by alan.shaw@protocol.ai
	}
	if want, got := plaintext, s; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
