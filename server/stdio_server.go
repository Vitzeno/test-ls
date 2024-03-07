package server

import (
	"context"

	"github.com/Vitzeno/test-ls/handlers"
	"github.com/Vitzeno/test-ls/types"
	"github.com/sourcegraph/jsonrpc2"
)

// Stdio is a server that listens for JSON-RPC messages on stdin and stdout.
type Stdio struct {
	handler *handlers.Handler
}

func NewStdio(handler *handlers.Handler) *Stdio {
	return &Stdio{
		handler: handler,
	}
}

func (s *Stdio) Serve(ctx context.Context) error {
	<-jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(types.Stdrwc{}, jsonrpc2.VSCodeObjectCodec{}),
		//jsonrpc2.AsyncHandler(handler),
		s.handler,
	).DisconnectNotify()

	return nil
}
