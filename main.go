package main

import (
	"context"
	"log"

	"github.com/Vitzeno/test-ls/handlers"
	"github.com/Vitzeno/test-ls/server"
)

func main() {
	ctx := context.Background()

	handler := handlers.New()
	err := server.StdHandler(ctx, handler)
	if err != nil {
		log.Fatalf("failed to handle tcp connection: %v", err)
	}

	// bytes, _ := io.ReadAll(internal.StdioCloser{})
	// fmt.Print(string(bytes))
}
