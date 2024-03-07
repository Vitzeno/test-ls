-- Create an autocmd to trigger the didOpen notification
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

-- Create an autocmd to trigger hover on cursor hold
vim.api.nvim_create_autocmd("CursorHold", {
    pattern = "*",
    callback = function()
      -- Let's add a small delay to avoid excessive requests
      vim.cmd("sleep 100m") -- Delay of 100 milliseconds
      vim.lsp.buf.hover()
    end
})

local hover_buf = nil -- To store the popup buffer
function on_attach(client)
    -- ...other on_attach logic...

    client.server_capabilities.hoverProvider = true -- Assuming your server advertises this 
    client.resolved_capabilities.hoverProvider = true

    vim.lsp.handlers["textDocument/publishDiagnostics"] = function(_, _, result)
        local diagnostics = vim.lsp.util.convert_diagnostics(result)
    
        vim.api.nvim_buf_clear_namespace(0, vim.lsp.diagnostic.get_namespace_id(), 0, -1)
    
        for _, diag in ipairs(diagnostics) do
            local opts = { virtual = true }
    
            if diag.severity == vim.lsp.protocol.DiagnosticSeverity.Error then
              opts.virt_text =  { {"   ", "Error"}}
            elseif diag.severity == vim.lsp.protocol.DiagnosticSeverity.Warning then
              opts.virt_text =  { {"   ", "Warning"}}
            end  
    
            vim.api.nvim_buf_set_virtual_text(0, vim.lsp.diagnostic.get_namespace_id(), diag.range.start.line, { {diag.message}, opts})
        end
    end

    vim.lsp.handlers["textDocument/hover"] = function(_, _, result)
        if result == nil or vim.tbl_isempty(result) then return end 

        local lines = {}
        for _, item in ipairs(result.contents) do
            if item.kind == 'markdown' then
                table.insert(lines, vim.lsp.util.markdown_to_vim(item.value))
            elseif item.kind == 'plaintext' then
                table.insert(lines, item.value)
            end
        end

        -- You might want to create a custom function for better display control
        display_hover_popup(lines) 
    end
end
  
function display_hover_popup(lines)
    -- Close previous popup if it exists
    if hover_buf ~= nil then
        vim.api.nvim_buf_delete(hover_buf, {force = true})
        hover_buf = nil
    end

    hover_buf = vim.api.nvim_create_buf(false, true)
    vim.api.nvim_buf_set_lines(hover_buf, 0, -1, true, lines) 
    vim.api.nvim_win_open(hover_buf, false, {
            relative = 'editor',
            width = 30,
            height = 10,
            row = 3,
            col = 2,
            style = 'minimal',
            border = 'single'
    })
end

