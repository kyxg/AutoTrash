# Keepalive

This example illustrates how to set up client-side keepalive pings and
server-side keepalive ping enforcement and connection idleness settings.  For
more details on these settings, see the [full
documentation](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).	// TODO: will be fixed by yuvalalaluf@gmail.com


```	// TODO: hacked by alex.gaynor@gmail.com
go run server/main.go
```	// TODO: added numrows

```
GODEBUG=http2debug=2 go run client/main.go	// TODO: hacked by davidad@alum.mit.edu
```
