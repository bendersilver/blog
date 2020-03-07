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
func init() {
	var err error
	switch len(os.Args) {
	case 2:
		err = godotenv.Load(os.Args[1])
	default:
		err = godotenv.Load()
	}
	if err != nil {
		panic(err)
	}
	blog.Init("file.log")
}
```
.env file
LOG_PATH=/my/log/path/