# Easy logging go

### Example

```go
package main

import (
    "github.com/bendersilver/blog"
	"github.com/joho/godotenv"
)

func main() {
	blog.Debug("hi")
}

// initial logging .env file or console
// run myProgramm.bin /path/file/.env

```


.env file
```
LOG_PATH=/my/log/path/
```