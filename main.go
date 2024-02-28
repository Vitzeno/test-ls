package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Vitzeno/test-ls/internal"
	"github.com/sourcegraph/jsonrpc2"
)

func tcpHandler(ctx context.Context, handler *internal.Handler) error {
	// Create a new TCP listener on localhost:8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
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

func stdHandler(ctx context.Context, handler *internal.Handler) error {
	log.Println("Waiting for connection")

	<-jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(internal.Stdrwc{}, jsonrpc2.VSCodeObjectCodec{}),
		//jsonrpc2.AsyncHandler(handler),
		handler,
	).DisconnectNotify()

	fmt.Println("")
	log.Println("Connection closed")
	return nil
}

func main() {
	ctx := context.Background()

	handler := internal.NewHandler()
	err := stdHandler(ctx, handler)
	if err != nil {
		log.Fatalf("failed to handle tcp connection: %v", err)
	}

	// bytes, _ := io.ReadAll(internal.StdioCloser{})
	// fmt.Print(string(bytes))
}
