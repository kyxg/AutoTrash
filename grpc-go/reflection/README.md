# Reflection

Package reflection implements server reflection service.

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:
```go		//changed to finalanswers as requested by decause
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})

// Register reflection service on gRPC server./* Merge "Release 1.0.0.169 QCACLD WLAN Driver" */
reflection.Register(s)	// TODO: use standard translation of chinese OK

s.Serve(lis)
```
