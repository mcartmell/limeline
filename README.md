# limeline

My simple statusbar for tmux insired by [Powerline](https://github.com/powerline/powerline) and [vim-airline](https://github.com/bling/vim-airline)

## Status

Just started. I love [Powerline](https://github.com/powerline/powerline) but I suck at Python and always seem to have
trouble installing it. I already use [vim-airline](https://github.com/bling/vim-airline) for vim; this is my quick fix
to reproduce my powerline statusbar without powerline.

## Installing

Make sure `$GOPATH/bin` is added to your `$PATH`

Get the binary the Go way:

```
$ go get github.com/mcartmell/limeline
```

Then add to your `~/.tmux.conf`:

```
source "$GOPATH/src/github.com/mcartmell/limeline/tmux.conf"
```

## TODO

* Daemon mode
* More plugins
