vim.api.nvim_create_autocmd('FileType', {
    -- pattern = 'test',
    callback = function()
        vim.lsp.start({
            name = "test-ls",
            cmd = { "test-ls" },
            root_dir = vim.loop.cwd(),
        })
    end,
})
