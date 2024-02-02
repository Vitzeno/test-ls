package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sourcegraph/jsonrpc2"
)

type NotificationHandler func(params json.RawMessage) error
type MethodHandler func(params json.RawMessage) (json.RawMessage, error)

type Handler struct {
	NotificationHandlers map[string]NotificationHandler
	MethodHandlers       map[string]MethodHandler
}

func NewHandler() *Handler {
	return &Handler{
		NotificationHandlers: map[string]NotificationHandler{},
		MethodHandlers: map[string]MethodHandler{
			"initialize": Initialize,
		},
	}
}

func (h *Handler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	// Implement handling of different LSP requests here
	// log.Println("Received request:", req.Method)
	// log.Println("Received params:", req.Params)

	res, err := h.process(req)
	if err != nil {
		err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
		if err := conn.ReplyWithError(ctx, req.ID, err); err != nil {
			log.Println(err)
			return
		}
	}

	if res != nil {
		var result jsonrpc2.Response
		result.ID = req.ID
		result.Result = &res

		if err := conn.Reply(ctx, req.ID, result); err != nil {
			log.Println("Error replying to request:", err)
		}
	}
}

func (h *Handler) process(req *jsonrpc2.Request) (json.RawMessage, error) {
	params := []byte(``)
	if req.Params != nil {
		params = *req.Params
	}

	if req.Notif {
		if nh, ok := h.NotificationHandlers[req.Method]; ok {
			return nil, nh(params)
		}
		// TODO: use error code from lsp spec, or is the jsonrpc2.CodeMethodNotFound wrapper enough?
		return nil, fmt.Errorf("no notification handler for method %q", req.Method)
	}

	if mh, ok := h.MethodHandlers[req.Method]; ok {
		return mh(params)
	}

	// TODO: use error code from lsp spec, or is the jsonrpc2.CodeMethodNotFound wrapper enough?
	return nil, fmt.Errorf("no method handler for method %q", req.Method)
}
