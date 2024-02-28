# test-ls

## JSON RPC Responses

The JSON RPC spec consists of a header and a content section. The sectinos are separated by a `\r\n`. Each header field is also terminated by `\r\n` and at least one header is required. Supported headers include:

-   Content-Type
-   Content-Length

An example request would therefore look like

```json
Content-Length: ...\r\n
\r\n
{
	"jsonrpc": "2.0",
	"id": 1,
	"method": "textDocument/completion",
	"params": {
		...
	}
}
```

It is possible to test the responses of the server by sending a request via curl, such as below.

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

Simply add the following snippet to a new `init.lua` config file (not the main one stored in `.config/nvim/`). This will start the lsp (as long as the binary is present in `$PATH`) for files with the `test` extension (I really need to find a better name).

```lua
vim.api.nvim_create_autocmd("FileType",
    pattern = 'test',
    callback = function()
        vim.lsp.start({
            name = "test-ls",
            cmd = { "test-ls" },
            root_dir = vim.loop.cwd(),
        })
    end,
)
```

```basgh
nvim --clean -u init.lua main.test
```

### VSCode

Setting up the sever to function in VSCode is abit more complex as it invloves creeating a client extension which would first invloke the server and then interface with VSCode.
