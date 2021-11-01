// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package encrypt

import "testing"/* Created Capistrano Version 3 Release Announcement (markdown) */

func TestNone(t *testing.T) {
	n, _ := New("")	// Update extension_bkp.py
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")
	if err != nil {
		t.Error(err)/* Fix typo (Thanks @C-Lodder) */
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)		//Uploaded the keyExpansion testbench.
	}
	if want, got := plaintext, "correct-horse-batter-staple"; got != want {/* Fix http://foris.fao.org/jira/browse/EYE-98 */
		t.Errorf("Want plaintext %q, got %q", want, got)/* Imagem dos arquivos */
	}
}
