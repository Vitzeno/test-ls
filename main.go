package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/sourcegraph/jsonrpc2"
)

type Result struct {
	Label string `json:"label"`
}

type Handler struct{}

func (s *Handler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	// Implement handling of different LSP requests here
	log.Println("Received request:", req.Method)

	switch req.Method {
	case "initialize":
		var result jsonrpc2.Response
		if err := json.Unmarshal([]byte(`{"id":123,"result":{"foo":"bar"},"jsonrpc":"2.0"}`), &result); err != nil {
			log.Println(err)
			return
		}

		if err := conn.Reply(ctx, req.ID, result); err != nil {
			log.Println(err)
			return
		}
	default:
		err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
		if err := conn.ReplyWithError(ctx, req.ID, err); err != nil {
			log.Println(err)
			return
		}
	}
}

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

		// b := make([]byte, 1024)
		// _, err = conn.Read(b)
		// if err != nil {
		// 	log.Printf("failed to read connection: %v", err)
		// 	continue
		// }

		// log.Println(string(b))

		/*
			we keep getting jsonrpc2 protocol errors because requests looks like:

			POST / HTTP/1.1
			Content-Type: application/json
			User-Agent: PostmanRuntime/7.36.1
			Accept: */ /*
			Postman-Token: 38ba7b6e-21f5-4f60-ab44-f4e48ad9a071
			Host: localhost:8080
			Accept-Encoding: gzip, deflate, br
			Connection: keep-alive
			Content-Length: 69

			{
				"jsonrpc": "2.0",
				"method": "initialize",
				"id": 1
			}

			note the headers before the json-rpc request body
		*/

		go jsonrpc2.NewConn(
			ctx,
			jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}),
			&Handler{},
		)

		//go handleConnection(ctx, conn)
	}
}

// func handleConnection(ctx context.Context, conn net.Conn) {
// 	rpcConn := jsonrpc2.NewConn(
// 		ctx,
// 		jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}),
// 		&Handler{},
// 	)

// 	log.Println("rpc connection established")

// 	defer rpcConn.Close()
// }
