package main

import (
	"context"
	"log"
	"net"

	"github.com/Vitzeno/test-ls/internal"
	"github.com/sourcegraph/jsonrpc2"
)

func main() {
	ctx := context.Background()
	// Create a new TCP listener on localhost:8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	log.Println("Server listening on localhost:8080")

	// Accept incoming connections and handle them
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}

		handler := internal.NewHandler()
		go jsonrpc2.NewConn(
			ctx,
			jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}),
			handler,
		)
	}
}
