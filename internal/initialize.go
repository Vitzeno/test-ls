package internal

import (
	"encoding/json"
)

type InitializeParams struct {
	ProcessId             int                `json:"processId"`
	ClientInfo            ClientInfo         `json:"clientInfo"`
	Locale                string             `json:"locale"`
	RootPath              string             `json:"rootPath"`
	RootUri               string             `json:"rootUri"`
	InitializationOptions interface{}        `json:"initializationOptions"`
	Capabilities          ClientCapabilities `json:"capabilities"`
	Trace                 string             `json:"trace"`
	WorkspaceFolders      []string           `json:"workspaceFolders"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ClientCapabilities struct {
}

func Initialize(params json.RawMessage) (json.RawMessage, error) {
	// var p InitializeParams
	// err := json.Unmarshal(*params, &p)
	// if err != nil {
	// 	return nil, err
	// }

	return []byte(`{"foo":"bar"}`), nil
}

/*
interface InitializeParams extends WorkDoneProgressParams {

	 * The process Id of the parent process that started the server. Is null if
	 * the process has not been started by another process. If the parent
	 * process is not alive then the server should exit (see exit notification)
	 * its process.

	 processId: integer | null;


	  * Information about the client
	  *
	  * @since 3.15.0

	 clientInfo?: {

		  * The name of the client as defined by the client.

		 name: string;


		  * The client's version as defined by the client.

		 version?: string;
	 };


	  * The locale the client is currently showing the user interface
	  * in. This must not necessarily be the locale of the operating
	  * system.
	  *
	  * Uses IETF language tags as the value's syntax
	  * (See https://en.wikipedia.org/wiki/IETF_language_tag)
	  *
	  * @since 3.16.0

	 locale?: string;


	  * The rootPath of the workspace. Is null
	  * if no folder is open.
	  *
	  * @deprecated in favour of `rootUri`.

	 rootPath?: string | null;

	  * The rootUri of the workspace. Is null if no
	  * folder is open. If both `rootPath` and `rootUri` are set
	  * `rootUri` wins.
	  *
	  * @deprecated in favour of `workspaceFolders`

	 rootUri: DocumentUri | null;


	  * User provided initialization options.

	 initializationOptions?: LSPAny;


	  * The capabilities provided by the client (editor or tool)

	 capabilities: ClientCapabilities;


	  * The initial trace setting. If omitted trace is disabled ('off').

	 trace?: TraceValue;


	  * The workspace folders configured in the client when the server starts.
	  * This property is only available if the client supports workspace folders.
	  * It can be `null` if the client supports workspace folders but none are
	  * configured.
	  *
	  * @since 3.6.0

	 workspaceFolders?: WorkspaceFolder[] | null;
 }
*/
