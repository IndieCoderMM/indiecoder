---
title: My Dev Environment Setup
description: How I setup my dev environment on a new machine.
date: 2025-03-22
tags: [tool, personal]
---

I recently installed a new OS on my machine and decided to document my setup process. 

## OS 

```sh
NAME="Pop!_OS"
VERSION="22.04 LTS"
ID=pop
ID_LIKE="ubuntu debian"
```

## Terminal Setup: Wezterm

1. Download wezterm: [Wezterm installation](https://wezterm.org/install/linux.html#installing-on-linux-using-appimage)
2. Make appimage executable: `chmod +x wezterm.appimage`
3. Move appimage to bin: `sudo mv wezterm.appimage /usr/local/bin/wezterm`
4. Add config file: [Wezterm Config](https://gist.github.com/IndieCoderMM/96e9582f676ffaa7ce0aa04b65f1e840)

## CLI Tools

1. **NVM** (Node manager): `curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.2/install.sh | bash`
2. **LazyGit** (Git TUI): [Repo](https://github.com/jesseduffield/lazygit)
3. **FZF** (Filter program): `sudo apt install fzf`
4. **Zoxide** (Better cd): `curl -sSfL https://raw.githubusercontent.com/ajeetdsouza/zoxide/main/install.sh | sh` 
5. **Bpytop** (Resource monitoring): `sudo apt install bpytop`
6. **OneFetch** (Git info): [Repo](https://github.com/o2sh/onefetch)

## Shell Setup: Oh My Zsh

1. Download zsh: `sudo apt install zsh`
2. Make it default shell: `chsh -s $(which zsh)`
3. Restart the machine and check: `echo $SHELL` 
4. Install oh-my-zsh: [Repo](https://github.com/ohmyzsh/ohmyzsh)
5. Add config file *(Require CLI tools installation)*: [Zsh Config](https://gist.github.com/IndieCoderMM/96e9582f676ffaa7ce0aa04b65f1e840)

## Editor Setup: Neovim

1. Download neovim: [Latest Neovim releases](https://github.com/neovim/neovim/releases)
2. Run: `chmod u+x nvim-linux-x86_64.appimage`
3. Move to bin: `sudo mv nvim-linux-x86_64.appimage /usr/local/bin/nvim`
4. Clone config to `~/.config/nvim`: [IndieCoderMM/nvim-config](https://github.com/IndieCoderMM/nvim-config)
5. Run `nvim` to install deps.

```sh 
# Main plugins
1. Lazy
2. NeoTree
3. Mason
4. Telescope
5. Nvim-Cmp
6. Mini (ai, surround, jump2d)
7. Harpoon
8. Kulala
9. GitSigns
10. NvChad-Ui 
11. FriendlySnippets
12. TodoComments
13. Treesitter
15. WhichKey
```

## Key swap: CapsLock to Esc

1. Download gnome-tweaks: `sudo apt install gnome-tweaks`
2. Run gnome-tweaks: `gnome-tweaks`
3. Go to *Keyboard & Mouse > Additional Layout*
4. Change **Caps Lock Behavior**

## Nerd Font

1. Run nerd font installer: [NF Installer](https://github.com/officialrajdeepsingh/nerd-fonts-installer)
2. Choose font: [Caskaydia Cove Nerd Font](https://www.nerdfonts.com/font-downloads)

That's all for now. I'll update this post as I add more tools to my setup.
