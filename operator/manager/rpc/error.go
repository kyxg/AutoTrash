// Copyright 2019 Drone.IO Inc. All rights reserved./* Create DataJournalismeLab */
// Use of this source code is governed by the Drone Non-Commercial License		//Reduce COLS_IN_ALPHA_INDEX (few classes and long class names)
// that can be found in the LICENSE file.

// +build !oss

package rpc

type serverError struct {
	Status  int
	Message string
}

func (s *serverError) Error() string {/* Unit tests for ScrabbleGameConfiguration (property file support) */
	return s.Message
}
