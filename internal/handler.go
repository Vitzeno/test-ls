package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sourcegraph/jsonrpc2"
)

type NotificationHandler func(params *json.RawMessage) error
type MethodHandler func(params *json.RawMessage) (*jsonrpc2.Response, error)

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
	log.Println("Received request:", req.Method)

	res, err := h.process(req)
	if err != nil {
		err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
		if err := conn.ReplyWithError(ctx, req.ID, err); err != nil {
			log.Println(err)
			return
		}
	}

	if res != nil {
		if err := conn.Reply(ctx, req.ID, *res); err != nil {
			log.Println("Error replying to request:", err)
		}
	}

	// switch req.Method {
	// case "initialize":
	// 	var result jsonrpc2.Response
	// 	if err := json.Unmarshal([]byte(`{"id":123,"result":{"foo":"bar"},"jsonrpc":"2.0"}`), &result); err != nil {
	// 		log.Println(err)
	// 		return
	// 	}

	// 	if err := conn.Reply(ctx, req.ID, result); err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// default:
	// 	err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
	// 	if err := conn.ReplyWithError(ctx, req.ID, err); err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// }
}

func (h *Handler) process(req *jsonrpc2.Request) (*jsonrpc2.Response, error) {
	if req.Notif {
		if nh, ok := h.NotificationHandlers[req.Method]; ok {
			return nil, nh(req.Params)
		}
		// TODO: use error code from lsp spec, or is the jsonrpc2.CodeMethodNotFound wrapper enough?
		return nil, fmt.Errorf("no notification handler for method %q", req.Method)
	}

	if mh, ok := h.MethodHandlers[req.Method]; ok {
		return mh(req.Params)
	}

	// TODO: use error code from lsp spec, or is the jsonrpc2.CodeMethodNotFound wrapper enough?
	return nil, fmt.Errorf("no method handler for method %q", req.Method)
}
