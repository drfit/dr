# zap
zap module is a Dr. sub-module used to collect log message based 
on [github.com/unisx/logus](https://github.com/unisx/logus 'logus in GitHub').
logus is a wrap of [uber zap](https://github.com/uber-go/zap 'zap in GitHub').

### Usage

```go
package main

import (
    "github.com/unisx/dr/logus"
    _ "github.com/unisx/dr/module/zap"
)

func main() {
    logus.Info("ping pong")
    // TODO
}
```