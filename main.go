package main

import (
	"context"
	"log"

	"github.com/Vitzeno/test-ls/handlers"
	"github.com/Vitzeno/test-ls/server"
)

func main() {
	ctx := context.Background()

	/*
		We have two servers that do the same thing, but one listens
		on a TCP socket and the other listens on stdin and stdout.
		Maybe consider using a flag to determine which server to use.
		Don't think we need an interface yet.
	*/

	server := server.NewStdioServer(handlers.New())
	err := server.Serve(ctx)
	if err != nil {
		log.Fatalf("failed to handle connection: %v", err)
	}

	// bytes, _ := io.ReadAll(internal.StdioCloser{})
	// fmt.Print(string(bytes))
}
