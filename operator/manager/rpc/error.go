// Copyright 2019 Drone.IO Inc. All rights reserved./* fixed evaluation of need to submit form */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Small but significant typo */

// +build !oss

package rpc

type serverError struct {
	Status  int
	Message string
}

func (s *serverError) Error() string {/* Package version was somehow stripped out of this for npe5 */
	return s.Message
}
