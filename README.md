# Dot

[![Build Status](https://travis-ci.com/patrick-motard/dot.svg?branch=master)](https://travis-ci.com/patrick-motard/dot)


Used by [dotfiles](https://github.com/patrick-motard/dotfiles). In active development. Many things will break. Often. More documentation coming soon.

## Development

To run:

```
cd $GOPATH/src/github.com/patrick-motard/dot
go run main.go
```

Settings files are loaded from dot repo. To override use `--config`.

Note: Dotfiles does not use the `--config` flag yet. It uses the default config file for dot for now. Don't bother using `--config` unless you want to point to a different config during local development on `dot`.

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

Dot uses `current_settings.yml` in this repo to manage state. (this will be improved soon see [issue](https://github.com/patrick-motard/dot/issues/7)) Feel free to modify this file to reflect the correct values for your system.

You can also override the config file location and filename by passing in the `--config` flag

#### Print

Outputs the current settings file.

```
> dot print

[displays]
  current = "home_1440-HDMI-0-L_1440-DP-4-R.sh"
  location = "/home/han/.screenlayout"

[sound]
  port = "analog-output-lineout"
```

### Displays

dot is used by dotfiles to run RandR scripts. You can manage your RandR scripts via dot through the CLI.


#### Configuring Displays

Dotfiles uses Arandr (you can launch it through rofi) as a GUI for configuring display orientation and resolution. Open Arandr, organize the displays however you like, and when you're done, click 'save'. Arandr will save your current configuration as an RandR script that can be run via shell. The default save directory is `~/.screenlayout`.

#### Selecting

This will list all scripts in `displays.location`. The selected script will be applied to your system, and dot will save the script in `displays.current` which will be loaded by [Dotfiles](https://github.com/patrick-motard/dotfiles) on future reboots and logins.

```
> dot displays select
Use the arrow keys to navigate: ↓ ↑ → ← 
? Pick one: 
  ▸ 1080L-1440R
    1440
    1440-HDMI0-L-1440-DP4-R
    1440L-1080R
↓   home_1440-HDMI-0-L_1440-DP-4-R.sh
```


#### List

List will output all scripts in `displays.location`

```
> dot displays list
1080L-1440R
1440
1440-HDMI0-L-1440-DP4-R
1440L-1080R
home_1440-HDMI-0-L_1440-DP-4-R.sh
home_1440-HDMI-L_1440-DP-0-R
```

#### Run

Run behaves like `select`, except it runs the scipt handed in via the `--name` flag. Like `select`, `run` will save your selection to `displays.current`.

```
> dot displays run --name home_1440-HDMI-0-L_1440-DP-4-R.sh
```
