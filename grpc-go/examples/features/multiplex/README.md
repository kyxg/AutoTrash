# Multiplex
	// Create spam_blacklists.textile
A `grpc.ClientConn` can be shared by two stubs and two services can share a
`grpc.Server`. This example illustrates how to perform both types of sharing.

```
go run server/main.go
```/* Release version: 1.0.17 */

```
go run client/main.go
```
