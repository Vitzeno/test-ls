package main

import (
	"context"
	"net"
	"testing"

	"github.com/Vitzeno/test-ls/internal"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/stretchr/testify/require"
)

func TestJSONRPC(t *testing.T) {
	for _, tc := range []struct {
		name string
		req  string
		res  []byte
	}{
		{
			name: "initialize",
			req:  `initialize`,
			res:  []byte(`{"id":123,"result":{"foo":"bar"},"jsonrpc":"2.0"}`),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			serverConn, clientConn := net.Pipe()

			handler := internal.NewHandler()

			// Unmarshal the expectedRes byte slice into a jsonrpc2.Response
			var expectedRes jsonrpc2.Response
			require.NoError(t, expectedRes.UnmarshalJSON(tc.res))

			server := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(serverConn), handler)
			defer server.Close()

			client := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(clientConn), nil)
			defer client.Close()

			var acReq jsonrpc2.Request
			require.NoError(t, acReq.UnmarshalJSON([]byte(`{"id":123,"method":"`+tc.req+`","jsonrpc":"2.0"}`)))

			var actualRes jsonrpc2.Response
			err := client.Call(ctx, acReq.Method, nil, &actualRes)
			require.NoError(t, err)

			require.Equal(t, expectedRes, actualRes)
		})
	}
}
