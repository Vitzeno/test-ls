vim.lsp.start({
    name = 'test-ls',
    cmd = {'test-ls'},
    root_dir = vim.fs.dirname(vim.fs.find({'.test-ls'}, { upward = true })[1]),
})
