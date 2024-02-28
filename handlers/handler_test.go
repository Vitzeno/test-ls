package handlers

import (
	"context"
	"net"
	"testing"

	"github.com/Vitzeno/test-ls/types"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	for _, tc := range []struct {
		name      string
		reqMethod string
		resp      string
		params    InitializeParams
	}{
		{
			name:      "initialize",
			reqMethod: `initialize`,
			resp:      `{"capabilities":{"textDocumentSyncKind":1},"serverInfo":{"name":"test-ls","version":"0.0.1"}}`,
			params: InitializeParams{
				ProcessId: 1,
				ClientInfo: &ClientInfo{
					Name:    "test",
					Version: types.P("0.0.1"),
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			serverConn, clientConn := net.Pipe()

			handler := New()

			server := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(serverConn), handler)
			defer server.Close()

			client := jsonrpc2.NewConn(ctx, jsonrpc2.NewPlainObjectStream(clientConn), nil)
			defer client.Close()

			var actualResponse jsonrpc2.Response
			err := client.Call(ctx, tc.reqMethod, tc.params, &actualResponse)
			require.NoError(t, err)

			var expectedRes jsonrpc2.Response
			require.NoError(t, expectedRes.UnmarshalJSON([]byte(tc.resp)))
			require.Equal(t, expectedRes, actualResponse)
		})
	}
}
