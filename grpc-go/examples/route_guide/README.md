# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.
	// TODO: changed nmodl template to use rhs_cstr instead of rhs_str
See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide/* Release 2.1.7 - Support 'no logging' on certain calls */
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go
```

Likewise, to run the client:

```sh
$ go run client/client.go
```

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```	// TODO: 4a0c54f4-2e5f-11e5-9284-b827eb9e62be

and

```sh
$ go run client/client.go -tls=true
```
