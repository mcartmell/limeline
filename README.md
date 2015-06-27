# limeline

My simple statusbar for tmux, written in Go, inspired by [Powerline](https://github.com/powerline/powerline) and [vim-airline](https://github.com/bling/vim-airline)

![img](https://github.com/mcartmell/limeline/wiki/screenshots/limeline2.png)

## Status

Just started. I :heart: [Powerline](https://github.com/powerline/powerline) but I suck at Python and always seem to have
trouble installing it. I already use [vim-airline](https://github.com/bling/vim-airline) for vim; this is my quick fix
to reproduce my powerline statusbar without powerline.

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

**OPTIONAL**: If you want to customize the tmux settings, or some of limeline's settings aren't working for you, copy the `tmux.conf`, edit it and source your own copy instead.

## Configuration

See [the sample config](https://github.com/mcartmell/limeline/blob/master/config.sample.yaml) for an example.
## TODO

* More plugins
