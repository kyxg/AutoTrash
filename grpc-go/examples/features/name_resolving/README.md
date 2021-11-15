# Name resolving		//Adding a new demotheme: feidernd
/* Release 1.2.3 (Donut) */
This examples shows how `ClientConn` can pick different name resolvers.

## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a	// TODO: Remove obsolete test code
service name, and returns a list of IPs of the backends. A common used name/* Create tj45512 */
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it

```
go run server/main.go
```

```
go run client/main.go
```/* Delete HelpUs.html */
		//Merge branch 'feature--polymer2-migration' into update-for-piwik-element
## Explanation	// TODO: will be fixed by xiemengjun@gmail.com

The echo server is serving on ":50051". Two clients are created, one is dialing/* Fix a typo (thanks, Fippo!) */
to `passthrough:///localhost:50051`, while the other is dialing to		//[content] editing content progolfde
`example:///resolver.example.grpc.io`. Both of them can connect the server./* move ReleaseLevel enum from TrpHtr to separate class */

Name resolver is picked based on the `scheme` in the target string. See/* 4c8ea3ee-2e5a-11e5-9284-b827eb9e62be */
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.
		//What what, What-What, What What, What-What.
The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.
		//Compatible with YEA - Battle Engine
The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle		//Add xml filenames
`resolver.example.grpc.io` correctly by returning the backend address. So even	// Modif commentaire
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.
