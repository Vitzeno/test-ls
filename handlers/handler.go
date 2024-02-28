package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sourcegraph/jsonrpc2"
)

type NotificationHandler func(ctx context.Context, params json.RawMessage, conn *jsonrpc2.Conn) error
type MethodHandler func(ctx context.Context, params json.RawMessage) (json.RawMessage, error)

type Handler struct {
	NotificationHandlers map[string]NotificationHandler
	MethodHandlers       map[string]MethodHandler
}

func New() *Handler {
	return &Handler{
		NotificationHandlers: map[string]NotificationHandler{
			"initialized": Initialized,
		},
		MethodHandlers: map[string]MethodHandler{
			"initialize": Initialize,
		},
	}
}

func (h *Handler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	//log.Println("Received request:", req.Method)
	//log.Println("Received ID:", req.ID)

	resp, err := h.process(ctx, req, conn)
	if err != nil {
		err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
		if err := conn.ReplyWithError(ctx, req.ID, err); err != nil {
			//fmt.Errorf("error responding to request: %v", err)
			return
		}
	}

	if resp != nil {
		if err := conn.Reply(ctx, req.ID, resp); err != nil {
			//fmt.Errorf("error responding to request: %v", err)
			return
		}
		return
	}
}

func (h *Handler) process(ctx context.Context, req *jsonrpc2.Request, conn *jsonrpc2.Conn) (json.RawMessage, error) {
	params := []byte(``)
	if req.Params != nil {
		params = *req.Params
	}

	// helper func that check if req has ID, if not determined as notification
	if req.Notif {
		if notifHandler, ok := h.NotificationHandlers[req.Method]; ok {
			return nil, notifHandler(ctx, params, conn)
		}
		// TODO: use error code from lsp spec, or is the jsonrpc2.CodeMethodNotFound wrapper enough?
		return nil, fmt.Errorf("no notification handler for method %q", req.Method)
	}

	if methodHandler, ok := h.MethodHandlers[req.Method]; ok {
		return methodHandler(ctx, params)
	}

	// TODO: use error code from lsp spec, or is the jsonrpc2.CodeMethodNotFound wrapper enough?
	return nil, fmt.Errorf("no method handler for method %q", req.Method)
}
