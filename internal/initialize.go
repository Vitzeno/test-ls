package internal

import (
	"encoding/json"
	"log"

	"github.com/sourcegraph/jsonrpc2"
)

func Initialize(params *json.RawMessage) (*jsonrpc2.Response, error) {
	var res jsonrpc2.Response
	// TODO: Implement the initialize method
	if err := json.Unmarshal([]byte(`{"id":123,"result":{"foo":"bar"},"jsonrpc":"2.0"}`), &res); err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}
