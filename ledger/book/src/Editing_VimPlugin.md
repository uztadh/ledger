# Editing in Vim

A vim plugin is provided to apply syntax highlighting and account
autocomplete when editing. Install the **vim-ledger** plugin.

Below is the result of `:set filetype=ledger` in vim.

![vim syntax screenshot](consoleshots/vimsyn.png)

The plugin can also do folding, try `:set foldmethod=syntax`

![vim folding screenshot](consoleshots/vimfold.png)

## Format on Save

In order to format on save, set your vim config to the following:

```vim
let g:ledger_autofmt_bufwritepre = 1
```
