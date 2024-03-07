-- Create an autocmd to trigger the didOpen notification
vim.api.nvim_create_autocmd('BufRead', {
    callback = function()
        local params = {
            textDocument = vim.lsp.util.make_text_document_params()
        }
        vim.lsp.start({
            name = 'test-ls',
            cmd = {'test-ls'},
            root_dir = vim.fn.getcwd(),
        })
        vim.lsp.buf_notify(0, "textDocument/didOpen", params)
    end
})

-- Create an autocmd to trigger the hover request
vim.api.nvim_create_autocmd('CursorHold', {
    callback = function()
        vim.lsp.buf_request(0, "textDocument/hover", {}, function(err, result, context)
            if err then
                print("Error:", err)
                return
            end
            if result and result.contents then
                print(vim.inspect(result.contents))
            end
        end)
    end
})


-- vim.api.nvim_create_autocmd('BufNewFile', {
--     callback = function()
--         local params = {
--             textDocument = vim.lsp.util.make_text_document_params()
--         }
--         vim.lsp.start({
--             name = 'test-ls',
--             cmd = {'test-ls'},
--             root_dir = vim.fn.getcwd(),
--         })
--         vim.lsp.buf_notify(0, "textDocument/didOpen", params)
--     end
-- })

-- vim.api.nvim_create_autocmd('BufWinEnter', {
--     pattern = '*.test',
--     callback = function()
--         local params = {
--             textDocument = vim.lsp.util.make_text_document_params()
--         }
--         vim.lsp.buf_notify(0, "textDocument/didOpen", params)
--     end
-- })

-- vim.api.nvim_create_autocmd('FileType', {
--     pattern = '*.test',
--     callback = function()
--         local params = {
--             textDocument = vim.lsp.util.make_text_document_params()
--         }
--         vim.lsp.buf_notify(0, "textDocument/didOpen", params)
--     end
-- })

-- vim.api.nvim_create_autocmd('BufReadPost', {
--     pattern = '*.test',
--     callback = function()
--         local params = {
--             textDocument = vim.lsp.util.make_text_document_params()
--         }
--         vim.lsp.buf_notify(0, "textDocument/didOpen", params)
--     end
-- })
