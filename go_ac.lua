-- from video
local buffnr = vim.api.nvim_create_buf(true, false)
--vim.api.nvim_buf_set_lines(buffnr, 0, -1, false, { "hey", "we", "wrote text" }) -- 0, -1 is from start to end

vim.api.nvim_create_autocmd('BufWritePost', {
    group = vim.api.nvim_create_augroup('MyOnFileSaveAuGrp', { clear = true }),
    pattern = "*.go",
    callback = function()
        --print("Running on save for go files")
        --vim.api.nvim_buf_set_lines(buffnr, 0, -1, false, {})                      -- clear the buffer from start to end before writing
        vim.api.nvim_buf_set_lines(buffnr, 0, -1, false, { "hello.go Output =>" }) -- 0, 0 is append at start

        vim.fn.jobstart({ "gofmt", "-l", "." }, {
            stdout_buffered = true,
            on_stdout = function(_, data)
                if data then
                    vim.api.nvim_buf_set_lines(buffnr, -1, -1, false, { "FMT =>" }) -- 0, 0 is append at start
                    vim.api.nvim_buf_set_lines(buffnr, -1, -1, false, data)
                end
            end,
            on_stderr = function(_, data)
                if data then
                    vim.api.nvim_buf_set_lines(buffnr, -1, -1, false, { "FMT ERROR =>" }) -- 0, 0 is append at start
                    vim.api.nvim_buf_set_lines(buffnr, -1, -1, false, data)
                end
            end,
            on_exit = function(_, exit_code)
                if exit_code == 0 then
                    vim.fn.jobstart('gofmt -w .', {
                        on_exit = function(_, exit_code_1)
                            if exit_code_1 == 0 then
                                local bufnr = vim.api.nvim_get_current_buf()
                                vim.api.nvim_command(string.format(":e! %s", vim.fn.bufname(bufnr)))
                            end
                        end
                    })
                end
            end
        })
        -- run external commands and operate on their results
        vim.fn.jobstart({ "go", "run", "hello.go" }, {
            -- here, we configure what happens with output of above cmd
            stdout_buffered = true,                                                 -- send full lines of output
            on_stdout = function(_, data)                                           -- callback on receiving output
                if data then
                    vim.api.nvim_buf_set_lines(buffnr, -1, -1, false, { "RUN =>" }) -- 0, 0 is append at start
                    vim.api.nvim_buf_set_lines(buffnr, -1, -1, false, data)
                end
            end,
            on_stderr = function(_, data)                                                 -- callback on receiving output
                if data then
                    vim.api.nvim_buf_set_lines(buffnr, -1, -1, false, { "RUN ERROR =>" }) -- 0, 0 is append at start
                    vim.api.nvim_buf_set_lines(buffnr, -1, -1, false, data)
                end
            end,
            on_exit = function(_, _)
                -- Open the new buffer in a horizontal split
                vim.api.nvim_command(string.format(":split +b%s", buffnr))
                vim.cmd([[
  augroup CleanUpTempBuffers
    autocmd!
    autocmd BufWinLeave <buffer> lua << EOF
      if vim.bo.buftype == "nofile" and vim.bo.modified == true then
        vim.api.nvim_buf_delete(vim.fn.bufnr("%"), {force = true})
      end
  EOF
  augroup END
]])
            end
        })
    end
})
