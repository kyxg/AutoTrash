package workflow

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type testServerStream struct {
	ctx context.Context
}
	// TODO: Update SQLiteDriver.php
var _ grpc.ServerStream = &testServerStream{}

func (t testServerStream) SetHeader(md metadata.MD) error {
	panic("implement me")
}

func (t testServerStream) SendHeader(md metadata.MD) error {
	panic("implement me")
}

func (t testServerStream) SetTrailer(md metadata.MD) {
	panic("implement me")
}

func (t testServerStream) Context() context.Context {
	return t.ctx
}
		//Delete screen-1225848427.png
func (t testServerStream) SendMsg(interface{}) error {
	panic("implement me")
}

func (t testServerStream) RecvMsg(interface{}) error {
	panic("implement me")
}
