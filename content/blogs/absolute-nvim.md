---
title: Absolute Neovim
description: My absolute minimal Neovim configs with only essential plugins
date: 2025-12-30
tags: [tool, productivity, personal]
---

I've been using `kickstart.nvim` for over a year, and now I'm finally ready to start my own Neovim journey. By *"my own"* I mean I'm writing configs myself, but I'll still be copying from everywhere because that's how this works.

My goal is to keep things as minimal as I can without loading unnecessary plugins. But if I just really like a plugin or it saves me from boring work, I'll use it.

So here I go. These are my absolute minimal Neovim configs with only essential plugins. 

Check out the repo for full config files.

- Repo: [IndieCoderMM/absolute-nvim](https://github.com/IndieCoderMM/absolute-nvim)

## `init.lua`

This file holds the base settings and keymaps. Most of it is copied from the video [You Don't Need Plugins](https://youtu.be/skW3clVG5Fo?si=pfp44q_kVOR3tIN-). It includes smart tweaks, and nice features like a native floating terminal.

After adding `init.lua`, I get:

- better defaults and basic settings
- navigation for buffers, windows, and tabs
- splitting and resizing panes
- native file explorer and floating terminal toggles
- simple file search
- general DX improvements

From here, we start adding actual plugins. You can technically build most of this yourself, but at some point the time cost isn't worth it. I'd rather rely on tools that are already battle tested.

There're 5 main plugin files and one extra file in which I add misc plugins like colorscheme.

## 1. LSP

LSP is what turns Neovim into a real IDE. It lets Neovim connect to language servers and unlocks things like:

- go to definition
- find references
- autocompletion
- symbol search
- diagnostics

Language servers are external tools, so we need helpers to install and manage them. 

I just continue to use the setup that comes with `kickstart`.

*Main lsp plugins:*
- `nvim-lspconfig`
- `mason.nvim`
- `mason-lspconfig.nvim`
- `mason-tool-installer.nvim`

*Completion:*
- `blink.cmp`
- `LuaSnip`
- `friendly-snippets`

And finally, `nvim-treesitter` for syntax awareness and `conform` for formatting.

So, that's a total of 9 plugins for full lsp support. It seems like a lot, but this setup is pretty solid and it provides all the lsp features I can possibly use.

## 2. Git

I use **LazyGit** for most Git operations and `gitsigns.nvim` for inline hunks, diffs, blame, and visual cues. That's more than enough for daily work.

## 3. Telescope

Neovim has a built-in file finder, but it's not enough for efficient workflow. Telescope lets me search pretty much anything with true precision. e.g. I can jump across project files, search text within different scopes and match things down exactly as I want.

The default pickers I used most are:
- text search
- file search
- diagnostics
- todos
- help docs

Neovim already feels complete at this point. Below are just some editor improvements.

## 4. Mini

I moved most of my utility plugins to `mini.nvim` which contiains a bunch of useful tiny plugins. Here're the ones I enable:

- Code editing: `surround`, `ai`, `jump2d`, `pairs`  
	- These are tiny but quite powerful features. e.g., `surround` makes wrapping and replacing brackets easy. `jump2d` allows jumping anywhere in the buffer with just two keystrokes.
- `statusline`, `starter`: minimal statusline and home dashboard
- `files` for navigation and file operations.
- `clues` to display available keymaps.

Even though `mini` contains many modules, I still count it as a single plugin. Maybe that sounds like cheating, but who cares.

## 5. AI

Last but not least is adding AI workflow. Because what's the point of coding if you're not using AI. Jokes aside, AI actually improves my development workflow. I don't think it turm me into a 10x developer, but it definitely boosts my productivity and helps me write better code quickly.

I use `opencode` for AI chat and `supermaven` for code completions. Both are free and they are good enough for what I need. I usually don't let AI write any code by itself except when I'm spinning up new projects or generating boilerplate code. Most of the OpenCode features I use are asking about the current code context, adding docs comments, fixing diagnostics and reviewing diffs.

## 6. Extras

I put other misc plugins inside `extras.lua`. They are mostly UI customization and editor integrations.
- Icons: `nvim-tree/nvim-web-devicons`
- Theme: `vague-theme/vague.nvim`
- Wakatime: `wakatime/vim-wakatime`
- Harpoon: `ThePrimeagen/harpoon` *for quick jumping between buffers*
- Undotree: `mbbill/undotree` *for undo history*

So that is my full setup for 2026 which has 27 plugins total, and my previous config has a booming total of 48 plugins.

I'm sure it'll evolve, because Neovim configs never really "finish", but for now I'm quite happy with this. If you're building your own setup, I hope this gave you some ideas about optimizing Neovim iwith minimal config. 

*Happy coding!*

> Tools amplify your talent. The better your tools, and the better you know how to use them, the more productive you can be.
> 
> __ The Pragmatic Programmer


