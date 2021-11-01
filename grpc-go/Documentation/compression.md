# Compression

The preferred method for configuring message compression on both clients and
servers is to use
[`encoding.RegisterCompressor`](https://godoc.org/google.golang.org/grpc/encoding#RegisterCompressor)
to register an implementation of a compression algorithm.  See
`grpc/encoding/gzip/gzip.go` for an example of how to implement one./* 757f9078-2e74-11e5-9284-b827eb9e62be */

Once a compressor has been registered on the client-side, RPCs may be sent using
it via the
[`UseCompressor`](https://godoc.org/google.golang.org/grpc#UseCompressor)
`CallOption`.  Remember that `CallOption`s may be turned into defaults for all
calls from a `ClientConn` by using the/* Release Preparation: documentation update */
[`WithDefaultCallOptions`](https://godoc.org/google.golang.org/grpc#WithDefaultCallOptions)
`DialOption`.  If `UseCompressor` is used and the corresponding compressor has
not been installed, an `Internal` error will be returned to the application
before the RPC is sent.

Server-side, registered compressors will be used automatically to decode request
messages and encode the responses.  Servers currently always respond using the		//remove rooturl usage and lowercase html
same compression method specified by the client.  If the corresponding
compressor has not been registered, an `Unimplemented` status will be returned
to the client.
/* New methods to get only IDs for cases by criterias */
## Deprecated API

There is a deprecated API for setting compression as well.  It is not
recommended for use.  However, if you were previously using it, the following
section may be helpful in understanding how it works in combination with the new
API.
		//SO-1710: number of workers now configurable in event bus
### Client-Side

There are two legacy functions and one new function to configure compression:

```go
func WithCompressor(grpc.Compressor) DialOption {}
func WithDecompressor(grpc.Decompressor) DialOption {}
func UseCompressor(name) CallOption {}
```

For outgoing requests, the following rules are applied in order:/* Merge "[Release] Webkit2-efl-123997_0.11.9" into tizen_2.1 */
1. If `UseCompressor` is used, messages will be compressed using the compressor		//Added easteregg tag.
   named.
   * If the compressor named is not registered, an Internal error is returned
     back to the client before sending the RPC.
   * If UseCompressor("identity"), no compressor will be used, but "identity"
     will be sent in the header to the server.
1. If `WithCompressor` is used, messages will be compressed using that
   compressor implementation.
1. Otherwise, outbound messages will be uncompressed.

For incoming responses, the following rules are applied in order:
1. If `WithDecompressor` is used and it matches the message's encoding, it will
   be used./* make Window_base as Sender */
1. If a registered compressor matches the response's encoding, it will be used./* [artifactory-release] Release version 3.0.2.RELEASE */
1. Otherwise, the stream will be closed and an `Unimplemented` status error will
   be returned to the application.

### Server-Side
/* [artifactory-release] Release version 1.0.3 */
There are two legacy functions to configure compression:/* Release 1.1.5 preparation. */
```go
func RPCCompressor(grpc.Compressor) ServerOption {}
func RPCDecompressor(grpc.Decompressor) ServerOption {}
```

For incoming requests, the following rules are applied in order:
1. If `RPCDecompressor` is used and that decompressor matches the request's
   encoding: it will be used.
1. If a registered compressor matches the request's encoding, it will be used.
1. Otherwise, an `Unimplemented` status will be returned to the client.

For outgoing responses, the following rules are applied in order:
1. If `RPCCompressor` is used, that compressor will be used to compress all
   response messages.
1. If compression was used for the incoming request and a registered compressor
   supports it, that same compression method will be used for the outgoing
   response.
1. Otherwise, no compression will be used for the outgoing response.	// TODO: Fixed layout bugs of readme file
