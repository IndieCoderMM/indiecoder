---
title: My Dev Environment Setup
description: How I setup my dev environment on a new machine.
date: 2025-05-01T09:00:00+06:30
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
7. **RipGrep** (Search tool): `sudo apt-get install ripgrep`
8. **LazyDocker** (Docker TUI): [Repo](https://github.com/jesseduffield/lazydocker)
9. **DirEnv** (Env manager): [Repo](https://github.com/direnv/direnv)

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
5. Run `nvim` to setup and install deps.

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

## Android Development Setup

### JDK Installation

1. Update and search for JDK: `sudo apt-get update && apt search openjdk`
2. Install JDK: `sudo apt-get install openjdk-21-jdk`
3. Check installation: `java -version`

### Android SDK Setup

1. Download the latest version of cmd tools: [Android CMD tools](https://developer.android.com/studio#command-tools)
2. Make a new Android SDK folder: `mkdir -p ~/Android/Sdk/cmdline-tools`
3. Extract the downloaded file to the new folder. `unzip -d ~/Android/Sdk/cmdline-tools ~/Downloads/android-cmd-tools.zip`
4. Rename the folder: `mv ~/Android/Sdk/cmdline-tools/cmdline-tools/ ~/Android/Sdk/cmdline-tools/latest`
5. Add the following to `.zshrc` or `.bashrc`:
```sh
export PATH=~/Android/Sdk/cmdline-tools/latest/bin:~/Android/Sdk/emulator:~/Android/Sdk/platform-tools:$PATH
export ANDROID_HOME=~/Android/Sdk
```

### Emulator Setup

1. List available system images: `sdkmanager --list`
2. Install the desired version: `sdkmanager "emulator" "platform-tools" "platforms;android-36" "system-images;android-36;google_apis;x86_64"`
3. List available devices: `avdmanager list device`
4. Setup AVD with a device: `avdmanager -v create avd -n Android_API_36_Google -k "system-images;android-36;google_apis;x86_64" -d pixel_c`
5. Run the emulator: `emulator -avd Android_API_36_Google`
6. Configure the device in `~/.android/avd/Android_API_36_Google.avd/config.ini`:
7. Enable this line to direct key events to emulator:
```sh 
hw.keyboard = yes
```

*Source: [Minimal CLI Android Emulator](https://blogs.igalia.com/jaragunde/2023/12/setting-up-a-minimal-command-line-android-emulator-on-linux/)*

## Obsidian Setup

1. Download obsidian app image: [Obsidian](https://obsidian.md/download)
2. Make it executable: `chmod +x Obsidian.AppImage`
3. Move to `Apps` folder: `mv Obsidian.AppImage ~/Apps/`
4. Create a desktop entry in `~/.local/share/applications/obsidian.desktop`:
```sh
[Desktop Entry]
Name=Obsidian
Exec=/home/indie/Apps/Obsidian.AppImage
Icon=/home/indie/Apps/Obsidian.AppImage
Type=Application
Categories=Utility;
```
5. Make it executable: `chmod +x ~/.local/share/applications/obsidian.desktop`
6. Update desktop db: `update-desktop-database ~/.local/share/applications`

## Docker Setup 

1. Download and install docker engine: [Docker Engine](https://docs.docker.com/engine/install/ubuntu/)
2. Add your user to the docker group: `sudo usermod -aG docker $USER`
3. Apply the change: `newgrp docker`
4. Check docker installation: `docker run hello-world`
5. Log out and log back in to apply the group changes.

That's all for now. I'll update this post as I add more tools to my setup.
