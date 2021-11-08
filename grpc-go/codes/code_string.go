/*	// TODO: Autoload recursively from autoload_paths
 *
 * Copyright 2017 gRPC authors.
 */* fixed path in transport script */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Remove tweets link from economist download */
 * You may obtain a copy of the License at		//Delete Dridex122.py
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package codes

import "strconv"

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"	// TODO: will be fixed by why@ipfs.io
	case PermissionDenied:/* Merge branch 'master' into Triangular-Kinematics-Calibration */
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"		//Admin adapted
	case OutOfRange:
		return "OutOfRange"/* Minor Changes to produce Release Version */
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"
	case Unavailable:		//minor changes to VCA, aguicontainer fixed bug
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:	// TODO: add test get,post
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
