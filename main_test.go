package main

import (
	"context"
	"encoding/json"
	"net"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/stretchr/testify/require"
)

func TestJSONRPC(t *testing.T) {
	for _, tc := range []struct {
		name     string
		request  string
		response jsonrpc2.Response
	}{
		{
			name:    "initialize",
			request: `initialize`,
			response: jsonrpc2.Response{
				ID:     jsonrpc2.ID{Num: 123},
				Result: &json.RawMessage{},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			serverConn, clientConn := net.Pipe()

			server := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(serverConn), &Handler{})
			defer server.Close()

			client := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(clientConn), nil)
			defer client.Close()

			var result jsonrpc2.Response
			err := client.Call(ctx, tc.request, nil, &result)
			require.NoError(t, err)

			require.Equal(t, tc.response, result)
			spew.Dump(result)
		})
	}
}
