package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Vitzeno/test-ls/handlers"
	"github.com/sourcegraph/jsonrpc2"
)

// TcpServer is a server that listens for JSON-RPC messages on a TCP socket.
type TcpServer struct {
	handler *handlers.Handler
}

func NewTcpServer(handler *handlers.Handler) *TcpServer {
	return &TcpServer{
		handler: handler,
	}
}

func (t *TcpServer) Serve(ctx context.Context) error {
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
				t.handler,
			).DisconnectNotify()

			log.Println("Connection closed")
			// possible that curl cannot handle the \n\n after content-length header causing it to hang
		}(conn)
	}
}
