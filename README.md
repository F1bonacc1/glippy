# Glippy - Clipboard for Go

#### Supported Platforms:

- OSX
- Windows
- Linux, Unix

### Quick Start

```go
package main

import (
	"fmt"
	"github.com/f1bonacc1/glippy"
)

func main(){
  // set clipboard text
  glippy.Set("Hello Glippy")
  
  text, err := glippy.Get()
  if err != nil{
    panic(err)
  }
  fmt.Print(text)
}
```



### Credits

https://github.com/atotto/clipboard

https://github.com/aarzilli/nucular