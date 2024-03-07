package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

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
