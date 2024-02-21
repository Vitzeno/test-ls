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
		log.Println("Waiting for connection")
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}
		log.Println("Accepted connection")

		handler := internal.NewHandler()

		// Start a new goroutine to handle multiple connections concurrently
		go func(conn net.Conn) {
			<-jsonrpc2.NewConn(
				ctx,
				jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}),
				handler,
			).DisconnectNotify()

			log.Println("Connection closed")
			// possible that curl cannot handle the \n\n after content-length header causing it to hang
		}(conn)
	}
}
