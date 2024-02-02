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

		handler := internal.NewHandler()
		go jsonrpc2.NewConn(
			ctx,
			jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}),
			handler,
		)

		//go handleConnection(ctx, conn)
	}
}
