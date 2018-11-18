# Dot

Used by [dotfiles](https://github.com/patrick-motard/dotfiles). In active development. Many things will break. Often. More documentation coming soon.

## Development

To run:

```
cd $GOPATH/src/github.com/patrick-motard/dot
go run main.go
```

Settings files are loaded from dot repo. To override use `--config`.

```
go run main.go print
```

## Install

`go install`


## Usage

`dot print`

`dot --help`

`dot {command} --help`

### Config File

Dot uses `current_settings.toml` in this repo to manage state. (this will be improved soon see [issue](https://github.com/patrick-motard/dot/issues/7)) Feel free to modify this file to reflect the correct values for your system.


### Displays

dot is used by dotfiles to run RandR scripts. You can manage your RandR scripts via dot through the CLI.


#### Configuring Displays

Dotfiles uses Arandr (you can launch it through rofi) as a GUI for configuring display orientation and resolution. Open Arandr, organize the displays however you like, and when you're done, click 'save'. Arandr will save your current configuration as an RandR script that can be run via shell. The default save directory is `~/.screenlayout`.

#### Selecting (and applying) a Display Script

This will list all scripts in `displays.location`. The selected script will be applied to your system, and dot will save the script in `displays.current` which will be loaded by [Dotfiles](https://github.com/patrick-motard/dotfiles) on future reboots and logins.

```
 dot displays select
Use the arrow keys to navigate: ↓ ↑ → ← 
? Pick one: 
  ▸ 1080L-1440R
    1440
    1440-HDMI0-L-1440-DP4-R
    1440L-1080R
↓   home_1440-HDMI-0-L_1440-DP-4-R.sh
```
