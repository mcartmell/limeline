# limeline

My simple statusbar for tmux, written in Go, inspired by [Powerline](https://github.com/powerline/powerline), [tmux-powerline](https://github.com/erikw/tmux-powerline) and [vim-airline](https://github.com/bling/vim-airline)

![img](https://github.com/mcartmell/limeline/wiki/screenshots/limeline2.png)

## Status

I :heart: [Powerline](https://github.com/powerline/powerline) but I always seem to have trouble setting it up on a new workstation. [tmux-powerline](https://github.com/erikw/tmux-powerline) is a good alternative but is now deprecated, and I wanted to write my own plugins. So I came up with this. It's lightweight, easy to install and I can write plugins for it.

I use it myself, so it should be usable if you're looking for a statusbar alternative.

## Features

* Client/daemon mode: no waiting for the status bar to redraw
* 5 plugins:
  * Load average
  * The [Singapore PSI reading](http://www.nea.gov.sg/anti-pollution-radiation-protection/air-pollution-control/psi/psi)
  * Current date and time
  * Current weather and temperature
  * Last track played on LastFM
* Configurable colours

## Installing

Prerequisites:

* [Go](http://golang.org)
* [A patched Powerline font](https://github.com/powerline/fonts)

Make sure `$GOPATH/bin` or `$GOBIN` is added to your `$PATH` **before starting tmux**.

Get both the client and daemon binaries using `go get`:

```
$ go get github.com/mcartmell/limeline/...
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

Then restart tmux and you're good to go.

**OPTIONAL**: If you want to customize the tmux settings, or some of limeline's settings aren't working for you, copy the `tmux.conf`, edit it and source your own copy instead.

## Configuration

See [the sample config](https://github.com/mcartmell/limeline/blob/master/config.sample.yaml) for an example.

## Alternatives

* [Powerline](https://github.com/powerline/powerline). The original, written in Python.
* [tmux-powerline](https://github.com/erikw/tmux-powerline). A lightweight powerline just for tmux, written in bash.

## TODO

* More plugins
