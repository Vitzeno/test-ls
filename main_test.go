package main

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/Vitzeno/test-ls/internal"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/stretchr/testify/require"
)

func TestJSONRPC(t *testing.T) {
	for _, tc := range []struct {
		name string
		id   int
		req  string
		res  string
	}{
		{
			name: "initialize",
			id:   0,
			req:  `initialize`,
			res:  `{"id":%d,"result":{"foo":"bar"},"jsonrpc":"2.0"}`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			serverConn, clientConn := net.Pipe()

			handler := internal.NewHandler()

			// Unmarshal the expectedRes byte slice into a jsonrpc2.Response
			var expectedRes jsonrpc2.Response
			require.NoError(t, expectedRes.UnmarshalJSON([]byte(fmt.Sprintf(tc.res, tc.id))))

			server := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(serverConn), handler)
			defer server.Close()

			client := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(clientConn), nil)
			defer client.Close()

			var marshalledRequest jsonrpc2.Request
			require.NoError(t, marshalledRequest.UnmarshalJSON([]byte(fmt.Sprintf(`{"id":%d,"method":"%s","jsonrpc":"2.0"}`, tc.id, tc.req))))

			var actualResponse jsonrpc2.Response
			err := client.Call(ctx, marshalledRequest.Method, nil, &actualResponse)
			require.NoError(t, err)

			require.Equal(t, expectedRes, actualResponse)
		})
	}
}
