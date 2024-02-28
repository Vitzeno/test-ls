package server

import (
	"context"
	"fmt"
	"log"

	"github.com/Vitzeno/test-ls/handlers"
	"github.com/Vitzeno/test-ls/types"
	"github.com/sourcegraph/jsonrpc2"
)

// StdioServer is a server that listens for JSON-RPC messages on stdin and stdout.
type StdioServer struct {
	handler *handlers.Handler
}

func NewStdioServer(handler *handlers.Handler) *StdioServer {
	return &StdioServer{
		handler: handler,
	}
}

func (s *StdioServer) Serve(ctx context.Context) error {
	log.Println("Waiting for connection")

	<-jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(types.Stdrwc{}, jsonrpc2.VSCodeObjectCodec{}),
		//jsonrpc2.AsyncHandler(handler),
		s.handler,
	).DisconnectNotify()

	fmt.Println("")
	log.Println("Connection closed")
	return nil
}
