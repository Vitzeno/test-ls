package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sourcegraph/jsonrpc2"
)

func DidOpen(ctx context.Context, params json.RawMessage, conn *jsonrpc2.Conn) error {
	var p DidOpenTextDocumentParams
	err := json.Unmarshal(params, &p)
	if err != nil {
		return fmt.Errorf("error unmarshalling DidOpenTextDocumentParams params: %w", err)
	}

	log.Printf("didOpen params: %+v", p)

	return nil
}

func DidChange(ctx context.Context, params json.RawMessage, conn *jsonrpc2.Conn) error {
	var p DidChangeTextDocumentParams
	err := json.Unmarshal(params, &p)
	if err != nil {
		return fmt.Errorf("error unmarshalling DidChangeTextDocumentParams params: %w", err)
	}

	log.Printf("didChange params: %+v", p)

	return nil
}

func DidClose(ctx context.Context, params json.RawMessage, conn *jsonrpc2.Conn) error {
	var p DidCloseTextDocumentParams
	err := json.Unmarshal(params, &p)
	if err != nil {
		return fmt.Errorf("error unmarshalling DidCloseTextDocumentParams params: %w", err)
	}

	log.Printf("didClose params: %+v", p)

	return nil
}

func DidSave(ctx context.Context, params json.RawMessage, conn *jsonrpc2.Conn) error {
	var p DidSaveTextDocumentParams
	err := json.Unmarshal(params, &p)
	if err != nil {
		return fmt.Errorf("error unmarshalling DidSaveTextDocumentParams params: %w", err)
	}

	log.Printf("didSave params: %+v", p)

	return nil
}
