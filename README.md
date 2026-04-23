# Glippy - Clipboard for Go

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/) [![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/f1bonacc1/glippy/graphs/commit-activity) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) ![Go Report](https://goreportcard.com/badge/github.com/F1bonacc1/glippy) [![GoDoc](https://godoc.org/github.com/f1bonacc1/glippy?status.svg)](http://godoc.org/github.com/f1bonacc1/glippy)

Couldn't find a go clipboard package which is both multi-platform and doesn't require additional installations. So this package is a combination of work of wonderful and talented people who created the packages mentioned in credits.

### Supported Platforms:

- OSX
- Windows
- Linux, Unix (X11 and Wayland)
- OSC 52 fallback — used automatically over SSH or in headless environments where neither X11 nor Wayland is reachable

### OSC 52 fallback

When the native clipboard is unavailable on Linux/FreeBSD (no X11 connection, or Wayland without `wl-clipboard`), `Set` falls back to writing an [OSC 52](https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-Operating-System-Commands) escape sequence to `/dev/tty`. Most modern terminals honor this: iTerm2, WezTerm, kitty, Alacritty, foot, and xterm (with `allowWindowOps` enabled).

- **tmux**: set `set-option -g set-clipboard on` in `~/.tmux.conf` so tmux forwards the sequence. Glippy wraps the sequence in tmux's DCS passthrough when `TMUX` is set.
- **Size limit**: OSC 52 payloads are capped at ~75 KB of plaintext; larger selections return an error rather than silently truncating.
- **Disable**: set `GLIPPY_DISABLE_OSC52=1` to skip this fallback entirely.
- **Windows**: OSC 52 is not used on Windows; the native clipboard path is always taken.

### Quick Start

```shell
go get github.com/f1bonacc1/glippy
```



```go
package main

import (
	"fmt"
	"github.com/f1bonacc1/glippy"
)

func main(){
  // set clipboard text
  glippy.Set("Hello, Glippy")
  
  // get clipboard text
  text, err := glippy.Get()
  if err != nil{
    panic(err)
  }
  fmt.Print(text) // Output: "Hello, Glippy"
}
```


```go
package main

import (
  "context"
  "fmt"
  "github.com/f1bonacc1/glippy"
)

func main(){
  // create clipboard watch channel
  ch := glippy.Watch(context.Background())
	
  for data := range ch {
    fmt.Println(data)
  }
}
```


### Credits

https://github.com/atotto/clipboard

https://github.com/aarzilli/nucular

https://github.com/jezek/xgb