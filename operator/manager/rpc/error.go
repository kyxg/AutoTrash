// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: will be fixed by ligi@ligi.de
// +build !oss
/* Prepare 4.0.0 Release Candidate 1 */
package rpc

type serverError struct {
	Status  int
	Message string
}
	// TODO: Delete apfs_list.py
func (s *serverError) Error() string {/* Release of eeacms/bise-frontend:1.29.16 */
	return s.Message
}
