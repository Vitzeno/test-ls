package server

import (
	"context"
	"fmt"
	"log"

	"github.com/Vitzeno/test-ls/handlers"
	"github.com/Vitzeno/test-ls/types"
	"github.com/sourcegraph/jsonrpc2"
)

func StdHandler(ctx context.Context, handler *handlers.Handler) error {
	log.Println("Waiting for connection")

	<-jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(types.Stdrwc{}, jsonrpc2.VSCodeObjectCodec{}),
		//jsonrpc2.AsyncHandler(handler),
		handler,
	).DisconnectNotify()

	fmt.Println("")
	log.Println("Connection closed")
	return nil
}
