# ctxlog: Extend golang context with logging [![GoDoc](https://godoc.org/github.com/jhulten/go-ctxlog?status.png)](https://godoc.org/github.com/jhulten/go-ctxlog) [![Build Status](https://travis-ci.org/jhulten/go-ctxlog.svg?branch=master)](https://travis-ci.org/jhulten/go-ctxlog)

ctxlog is a Go library extending the standard context with logging. 

## Features

- Create a child context that contains a log context.
- Helper functions to exit or panic when the provided context is cancelled.
- Consistent API with the context standard library.
 
## Example


```golang
package main

import (
    "context"
    "log"
    "time"
    "github.com/jhulten/go-ctxlog"
)

func main() {
    ctx := context.Background()
    ctx = ctxlog.WithWriter(ctx, os.Stdout)

    
}
```

