package handlers

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeParams
type InitializeParams struct {
	/*
		The process Id of the parent process that started the server.
		If the parent process is not alive then the server should
		exit (see exit notification).
	*/
	ProcessId             int                `json:"processId"`
	Capabilities          ClientCapabilities `json:"capabilities,omitempty"`
	ClientInfo            *ClientInfo        `json:"clientInfo,omitempty"`
	Locale                *string            `json:"locale,omitempty"`
	InitializationOptions *interface{}       `json:"initializationOptions,omitempty"`
	Trace                 *string            `json:"trace,omitempty"`
	WorkspaceFolders      []WorkspaceFolder  `json:"workspaceFolders,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspaceFolder
type WorkspaceFolder struct {
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeResult
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *ServerInfo        `json:"serverInfo,omitempty"`
}

type ClientInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#clientCapabilities
type ClientCapabilities struct {
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentClientCapabilities
type TextDocumentClientCapabilities struct {
}

type ServerInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#serverCapabilities
type ServerCapabilities struct {
	/*
		Defines how text documents are synced.
		Is either a detailed structure defining each notification
		or for backwards compatibility the TextDocumentSyncKind number.
	*/
	TextDocumentSync *TextDocumentSyncKind `json:"textDocumentSyncKind,omitempty"`
	Hover            *bool                 `json:"hoverProvider,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentSyncKind
type TextDocumentSyncKind int // 0: None, 1: Full, 2: Incremental
const (
	None        TextDocumentSyncKind = 0
	Full        TextDocumentSyncKind = 1
	Incremental TextDocumentSyncKind = 2
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#showMessageRequestParams
type ShowMessageRequestParams struct {
	Type    MessageType          `json:"type"`
	Message string               `json:"message"`
	Actions []*MessageActionItem `json:"actions,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_showMessage
type ShowMessageParams struct {
	Type    MessageType `json:"type"`
	Message string      `json:"message"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#messageType
type MessageType int

const (
	Error   MessageType = 1
	Warning MessageType = 2
	Info    MessageType = 3
	Log     MessageType = 4
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#messageActionItem
type MessageActionItem struct {
	Title string `json:"title"`
}
