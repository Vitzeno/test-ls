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

	log.Printf("raw hover params: %+v", string(params))
	log.Printf("hover params: %+v", p)

	HoverResponse := HoverResponse{
		Contents: MarkupContent{
			Kind: MarkupKind{
				Markdown:  "markdown",
				PlainText: "plaintext",
			},
			Value: "Hello, world!",
		},
	}

	resp, err := json.Marshal(HoverResponse)
	if err != nil {
		return nil, fmt.Errorf("error marshalling hover result: %w", err)
	}

	log.Printf("hover response: %+v", string(resp))

	return resp, nil
}
