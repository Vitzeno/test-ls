package handlers

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

func New() *Handler {
	return &Handler{
		NotificationHandlers: map[string]NotificationHandler{},
		MethodHandlers: map[string]MethodHandler{
			"initialize": Initialize,
		},
	}
}

func (h *Handler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	// Implement handling of different LSP requests here
	log.Println("Received request:", req.Method)
	log.Println("Received ID:", req.ID)

	resp, err := h.process(req)
	if err != nil {
		err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
		if err := conn.ReplyWithError(ctx, req.ID, err); err != nil {
			log.Println(err)
			return
		}
	}

	if resp != nil {
		if err := conn.Reply(ctx, req.ID, resp); err != nil {
			log.Println("Error responding to request:", err)
		}
		return
	}
}

func (h *Handler) process(req *jsonrpc2.Request) (json.RawMessage, error) {
	params := []byte(``)
	if req.Params != nil {
		params = *req.Params
	}

	// helper func that check if req has ID, if not determined as notification
	if req.Notif {
		if notifHandler, ok := h.NotificationHandlers[req.Method]; ok {
			return nil, notifHandler(params)
		}
		// TODO: use error code from lsp spec, or is the jsonrpc2.CodeMethodNotFound wrapper enough?
		return nil, fmt.Errorf("no notification handler for method %q", req.Method)
	}

	if methodHandler, ok := h.MethodHandlers[req.Method]; ok {
		return methodHandler(params)
	}

	// TODO: use error code from lsp spec, or is the jsonrpc2.CodeMethodNotFound wrapper enough?
	return nil, fmt.Errorf("no method handler for method %q", req.Method)
}
