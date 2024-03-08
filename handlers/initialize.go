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

	log.Printf("initialize params: %+v", p)

	InitializeResult := InitializeResponse{
		Capabilities: ServerCapabilities{
			TextDocumentSync: types.P(Incremental),
			Hover:            types.P(true),
			Diagnostics: &DiagnosticsOptions{
				Identifier: types.P("test-ls"),
			},
		},
		ServerInfo: &ServerInfo{
			Name:    "test-ls",
			Version: types.P("0.0.1"),
		},
	}

	resp, err := json.Marshal(InitializeResult)
	if err != nil {
		return nil, fmt.Errorf("error marshalling initialize result: %w", err)
	}

	log.Printf("initialize response: %+v", string(resp))

	return resp, nil
}

func Initialized(ctx context.Context, params json.RawMessage, conn *jsonrpc2.Conn) error {
	// no need to bother with unmarshalling, params will be empty
	// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialized
	showMessageParams := ShowMessageParams{
		Type:    Info,
		Message: "We are go for launch!",
	}

	log.Printf("initialized response: %+v", showMessageParams)

	return conn.Notify(ctx, "window/showMessage", showMessageParams)
}
