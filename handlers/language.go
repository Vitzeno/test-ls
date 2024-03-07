package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"log"

	"github.com/Vitzeno/test-ls/types"
	"github.com/sourcegraph/jsonrpc2"
)

// TODO: actually read file and return diagnostics
func Diagnostics(ctx context.Context, params json.RawMessage, conn *jsonrpc2.Conn) error {
	// no need to bother with unmarshalling, params will be empty
	// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialized
	diag := PublishDiagnosticsParams{
		URI: "file:///path/to/file",
		Diagnostics: []Diagnostic{
			Diagnostic{
				Range: Range{
					Start: Position{
						Line:      0,
						Character: 0,
					},
					End: Position{
						Line:      0,
						Character: 12,
					},
				},
				Severity: types.P(DiagError),
				Message:  "Hello, world!",
			},
		},
		Version: types.P(1),
	}

	log.Printf("diagnostics: %+v", diag)

	return conn.Notify(ctx, "textDocument/publishDiagnostics", diag)

}

func Hover(ctx context.Context, params json.RawMessage) (json.RawMessage, error) {
	var p HoverParams
	err := json.Unmarshal(params, &p)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling hover params: %w", err)
	}

	log.Printf("hover params: %+v", p)

	HoverResponse := HoverResponse{
		Contents: MarkupContent{
			Kind: MarkupKind{
				PlainText: "plaintext",
			},
			Value: "Hello, world!",
		},
		Range: &Range{
			Start: Position{
				Line:      0,
				Character: 0,
			},
			End: Position{
				Line:      0,
				Character: 12,
			},
		},
	}

	resp, err := json.Marshal(HoverResponse)
	if err != nil {
		return nil, fmt.Errorf("error marshalling hover result: %w", err)
	}

	log.Printf("hover response: %+v", string(resp))

	return resp, nil
}
