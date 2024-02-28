package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Vitzeno/test-ls/types"
	"github.com/sourcegraph/jsonrpc2"
)

func Initialize(ctx context.Context, params json.RawMessage) (json.RawMessage, error) {
	var p InitializeParams
	err := json.Unmarshal(params, &p)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling initialize params: %w", err)
	}

	InitializeResult := InitializeResult{
		Capabilities: ServerCapabilities{
			TextDocumentSync: types.P(Full),
			Hover:            types.P(false),
		},
		ServerInfo: &ServerInfo{
			Name:    "test-ls",
			Version: types.P("0.0.1"),
		},
	}

	log.Printf("initialize result: %v", InitializeResult)

	resp, err := json.Marshal(InitializeResult)
	if err != nil {
		return nil, fmt.Errorf("error marshalling initialize result: %w", err)
	}

	return resp, nil
}

func Initialized(ctx context.Context, params json.RawMessage, conn *jsonrpc2.Conn) error {
	// no need to bother with unmarshalling, params will be empty
	// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialized
	ShowMessageParams := ShowMessageParams{
		Type:    Info,
		Message: "test-ls initialized",
	}

	log.Printf("initialized: %v", ShowMessageParams)

	return conn.Notify(ctx, "window/showMessage", ShowMessageParams)
}
