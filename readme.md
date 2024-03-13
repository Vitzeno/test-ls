# test-ls

## JSON RPC Responses

The JSON RPC spec consists of a header and a content section. The sectinos are separated by a `\r\n`. Each header field is also terminated by `\r\n` and at least one header is required. Supported headers include:

-   Content-Type
-   Content-Length

An example request would therefore look like

```json
Content-Length: 183\r\n
\r\n
{
	"jsonrpc": "2.0",
	"id": 1,
	"method": "textDocument/didSave",
	"params": {
		"TextDocument":{
            "URI":"file:///Users/Documents/Projects/test-ls/testdata/main.test"
        }
	}
}
```

Whilst running in TCP mode it is possible to test the responses by sending a curl request:

```bash
curl --http0.9 --location-trusted 'localhost:8080' --header "Content-Type: application/json" \
--data '{
    "jsonrpc": "2.0",
    "method": "initialize",
    "id": 1,
    "params": {
        ...
    }
}'
```

## Editor Setup

Neovim setup is the easist and is therefore used for most of the testing.

### Neovim setup

Simply add the following snippet to a new `init.lua` config file (not the main one stored in `.config/nvim/`). This will start the lsp (as long as the binary is present in `$PATH`).

```lua
vim.api.nvim_create_autocmd({ "BufRead", "BufReadPost", "BufNewFile" }, {
    callback = function()
        local params = {
            textDocument = vim.lsp.util.make_text_document_params()
        }
        vim.lsp.start({
            name = 'test-ls',
            cmd = {'test-ls'},
            root_dir = vim.fn.getcwd(),
            capabilities = capabilities,
            on_attach = on_attach,
        })
        vim.lsp.buf_notify(0, "textDocument/didOpen", params)
    end
})

...

```

Full `init.lua` file can be found [here](testdata/init.lua)

Note this starts the LSP on neovim startup rather than on a per file type basis.

Now run neovim in clean mode to avoid conflicts with any other plugins and use the new `init.lua` config file.

```bash
nvim --clean -u init.lua
```

For additional info on LSP setup with neovim consult this [guide](https://neovim.io/doc/user/lsp.html).

### VSCode

Setting up the sever to function in VSCode is abit more complex as it invloves creeating a client extension which would first invloke the server and then interface with VSCode.

### Example messages

#### Initialize Request

```json
{
	"jsonrpc": "2.0",
	"method": "initialize",
	"id": 1,
	"params": {
		"ProcessId": 53969,
		"ClientInfo": {
			"Name": "neovim",
			"Version": "0.9.5"
		},
		"WorkspaceFolders": [
			{
				"Uri": "file:///Users/mohamed/Documents/Projects/test-ls/testdata",
				"Name": "/Users/mohamed/Documents/Projects/test-ls/testdata"
			}
		]
	}
}
```

#### Initialize Response

```json
{
	"jsonrpc": "2.0",
	"method": "initialize",
	"id": 1,
	"params": {
		"capabilities": {
			"textDocumentSyncKind": 2,
			"hoverProvider": true,
			"diagnostics": {
				"identifier": "test-ls"
			}
		},
		"serverInfo": {
			"name": "test-ls",
			"version": "0.0.1"
		}
	}
}
```
