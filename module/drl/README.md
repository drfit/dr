# drL
drL module is a Dr. sub-module used to collect log message based 
on [github.com/unisx/logus](https://github.com/unisx/logus 'logus in GitHub')

### Usage

```$go
import (
    "github.com/unisx/dr/logus"
    _ "github.com/unisx/dr/module/drl"
)

func main() {
    logus.Info("ping pong")
    // TODO
}
```