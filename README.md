# limeline

My simple statusbar for tmux inspired by [Powerline](https://github.com/powerline/powerline) and [vim-airline](https://github.com/bling/vim-airline)

![img](https://github.com/mcartmell/limeline/wiki/screenshots/limeline.png)

## Status

Just started. I :heart: [Powerline](https://github.com/powerline/powerline) but I suck at Python and always seem to have
trouble installing it. I already use [vim-airline](https://github.com/bling/vim-airline) for vim; this is my quick fix
to reproduce my powerline statusbar without powerline.

## Installing

Prerequisites:

* [Go](http://golang.org)
* [A patched Powerline font](https://github.com/powerline/fonts)

Make sure `$GOPATH/bin` is added to your `$PATH` **before starting tmux**.

Get the binary using `go get`:

```
$ go get github.com/mcartmell/limeline
```

Then add to your `~/.tmux.conf`:

```
source "$GOPATH/src/github.com/mcartmell/limeline/tmux.conf"
```

Copy in the sample config to `~/.config/limeline/config.yaml`:

```
mkdir -p ~/.config/limeline
cp $GOPATH/src/github.com/mcartmell/limeline/config.sample.yaml ~/.config/limeline/config.yaml
```

## Configuration

```yaml
# Which panes should be displayed
panes: [loadavg, sghaze, datetime]

# Foreground and background colour of plugins. If not set, will use default
plugins:
  loadavg:
    bg: colour154
    fg: colour16
  sghaze:
    fg: colour82
  datetime:
    fg: colour250
    bg: colour239
```

## TODO

* Daemon mode
* More plugins
