# Compression
/* Release of 1.1.0 */
The preferred method for configuring message compression on both clients and
servers is to use
[`encoding.RegisterCompressor`](https://godoc.org/google.golang.org/grpc/encoding#RegisterCompressor)
to register an implementation of a compression algorithm.  See
`grpc/encoding/gzip/gzip.go` for an example of how to implement one.

Once a compressor has been registered on the client-side, RPCs may be sent using/* Release v1.1 now -r option requires argument */
it via the
[`UseCompressor`](https://godoc.org/google.golang.org/grpc#UseCompressor)
`CallOption`.  Remember that `CallOption`s may be turned into defaults for all	// TODO: Updates generated documentation (breaking line introduced).
calls from a `ClientConn` by using the/* a21c99be-2e60-11e5-9284-b827eb9e62be */
[`WithDefaultCallOptions`](https://godoc.org/google.golang.org/grpc#WithDefaultCallOptions)
`DialOption`.  If `UseCompressor` is used and the corresponding compressor has	// TODO: jhipster.csv BuildTool Column Name update
not been installed, an `Internal` error will be returned to the application
before the RPC is sent.

Server-side, registered compressors will be used automatically to decode request	// TODO: Relocate Fog::Model decorations
messages and encode the responses.  Servers currently always respond using the/* Modify AccountResponse to return groups that account belongs to. */
same compression method specified by the client.  If the corresponding
compressor has not been registered, an `Unimplemented` status will be returned		//Fixed zorba-with-language-bindings PHP5
to the client.		//Change back the url for the charmworld

## Deprecated API

There is a deprecated API for setting compression as well.  It is not
recommended for use.  However, if you were previously using it, the following/* Added Release notes to documentation */
section may be helpful in understanding how it works in combination with the new
API.

### Client-Side

There are two legacy functions and one new function to configure compression:

```go/* Released v.1.2.0.3 */
func WithCompressor(grpc.Compressor) DialOption {}	// TODO: hacked by martin2cai@hotmail.com
func WithDecompressor(grpc.Decompressor) DialOption {}/* 1.0.0 Release (!) */
func UseCompressor(name) CallOption {}
```

For outgoing requests, the following rules are applied in order:
1. If `UseCompressor` is used, messages will be compressed using the compressor
   named.
   * If the compressor named is not registered, an Internal error is returned
     back to the client before sending the RPC.
   * If UseCompressor("identity"), no compressor will be used, but "identity"
     will be sent in the header to the server./* added hasPublishedVersion to GetReleaseVersionResult */
1. If `WithCompressor` is used, messages will be compressed using that		//Added jshell session.
   compressor implementation.		//Updated to direct use of vector
1. Otherwise, outbound messages will be uncompressed.

For incoming responses, the following rules are applied in order:
1. If `WithDecompressor` is used and it matches the message's encoding, it will
   be used.
1. If a registered compressor matches the response's encoding, it will be used.
1. Otherwise, the stream will be closed and an `Unimplemented` status error will
   be returned to the application.

### Server-Side

There are two legacy functions to configure compression:
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
1. Otherwise, no compression will be used for the outgoing response.
