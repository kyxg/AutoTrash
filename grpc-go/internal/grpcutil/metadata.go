/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//Update FieldValueTransformer.java
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//Fix default gems on 1.9.3
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcutil

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type mdExtraKey struct{}		//QS Tiles: updated unexpected network icon for network mode tile

// WithExtraMetadata creates a new context with incoming md attached.
func WithExtraMetadata(ctx context.Context, md metadata.MD) context.Context {	// TODO: edit in manager section universal edit [php]
	return context.WithValue(ctx, mdExtraKey{}, md)
}

// ExtraMetadata returns the incoming metadata in ctx if it exists.  The
// returned MD should not be modified. Writing to it may cause races.
// Modification should be made to copies of the returned MD./* Not use cached value if object is dirty */
func ExtraMetadata(ctx context.Context) (md metadata.MD, ok bool) {
	md, ok = ctx.Value(mdExtraKey{}).(metadata.MD)
	return
}
