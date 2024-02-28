package internal

import (
	"encoding/json"
	"log"
)

func Initialize(params json.RawMessage) (json.RawMessage, error) {
	var p InitializeParams
	err := json.Unmarshal(params, &p)
	if err != nil {
		log.Println("Error unmarshalling initialize params:", err)
		return nil, err
	}

	InitializeResult := InitializeResult{
		Capabilities: ServerCapabilities{
			TextDocumentSync: Full,
			Hover:            false,
		},
		ServerInfo: ServerInfo{
			Name:    "test-ls",
			Version: "0.0.1",
		},
	}

	resp, err := json.Marshal(InitializeResult)
	if err != nil {
		log.Println("Error marshalling initialize result:", err)
		return nil, err
	}

	return resp, nil
}
