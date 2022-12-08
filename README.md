# Glippy - Clipboard for Go

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/) [![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) ![Go Report](https://goreportcard.com/badge/github.com/F1bonacc1/glippy) [![GoDoc](https://godoc.org/github.com/f1bonacc1/glippy?status.svg)](http://godoc.org/github.com/f1bonacc1/glippy)

Couldn't find a go clipboard package which is both multi-platform and doesn't require additional installations. So this package is a combination of work of wonderful and talented people who created the packages mentioned in credits.

### Supported Platforms:

- OSX
- Windows
- Linux, Unix

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



### Credits

https://github.com/atotto/clipboard

https://github.com/aarzilli/nucular

https://github.com/jezek/xgb