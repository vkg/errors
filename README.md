# vkg/errors

[![Build Status](https://travis-ci.com/vkg/errors.svg?branch=master)](https://travis-ci.com/vkg/errors)
[![GoDoc](https://godoc.org/github.com/vkg/errors?status.svg)](https://godoc.org/github.com/vkg/errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/vkg/errors)](https://goreportcard.com/report/github.com/vkg/errors)

New generation Go error handling library, highly respecting [vkgtaro](http://github.com/vkgtaro)

## Synopsis

```
package main

 import (
	stderrors "errors"
	"fmt"

 	"github.com/vkg/errors"
)

 func main() {
	err := errors.Wrap(stderrors.New("new error handling library"))
	fmt.Printf("%+v", err)
}
```

### Try now

:point_right: https://play.golang.org/p/QIXoOrtLsEU :point_left:
