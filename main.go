package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/Vitzeno/test-ls/handlers"
	"github.com/Vitzeno/test-ls/server"
)

func main() {
	ctx := context.Background()

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the path to the Desktop directory
	desktopDir := filepath.Join(currentUser.HomeDir, "Desktop")

	// Log to file since LSP uses stdin and stdout
	file, err := os.OpenFile(filepath.Join(desktopDir, "test-ls.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	server := server.NewStdio(handlers.New())
	err = server.Serve(ctx)
	if err != nil {
		log.Fatalf("failed to handle connection: %v", err)
	}
}
