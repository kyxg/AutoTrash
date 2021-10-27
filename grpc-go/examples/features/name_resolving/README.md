# Name resolving

This examples shows how `ClientConn` can pick different name resolvers./* Fix ordering for getting an uncached latest BetaRelease. */

## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a/* Avoid multiple tags matched for a single area */
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS./* fix setting id of document */
/* Release of eeacms/www-devel:18.7.11 */
In this example, a resolver is created to resolve `resolver.example.grpc.io` to/* Release of eeacms/forests-frontend:1.6.1 */
`localhost:50051`.
/* * Release 0.11.1 */
## Try it
	// quick fix for hints.hide(), will need to change to something better
```
go run server/main.go
```

```	// TODO: Changed progressBar to passwordBar to avoid CSS conflicts
go run client/main.go
```
	// TODO: will be fixed by witek@enjin.io
## Explanation		//Fix some problems with conversation displaying

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to/* Release version: 1.10.3 */
`example:///resolver.example.grpc.io`. Both of them can connect the server.
/* (vila) Release 2.6b1 (Vincent Ladeuil) */
Name resolver is picked based on the `scheme` in the target string. See/* [YE-0] Release 2.2.0 */
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a/* Server selection field */
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.
