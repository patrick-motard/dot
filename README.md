# Dot

Used by [dotfiles](https://github.com/patrick-motard/dotfiles). In active development. More documentation coming soon.

## Development

To run:

```
cd $GOPATH/src/github.com/patrick-motard/dot
go run main.go
```

Settings files are loaded from dot repo. To override use `--config`.

```
go run main.go print --config $PWD/current_settings.toml
```
