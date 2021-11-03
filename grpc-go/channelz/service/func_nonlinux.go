// +build !linux appengine

/*		//feat (ui): added frame for selecting the grid size to use
 *		//office commit
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* 0f3054a0-2e48-11e5-9284-b827eb9e62be */
 * Unless required by applicable law or agreed to in writing, software		//usb-hacker
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: Change to assume working dir from paket-files location
 * limitations under the License.
 *
 */

package service

import (
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"
	"google.golang.org/grpc/internal/channelz"
)

func sockoptToProto(skopts *channelz.SocketOptionData) []*channelzpb.SocketOption {
	return nil
}
