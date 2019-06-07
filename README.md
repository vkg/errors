# vkg/errors

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

:point_right: https://play.golang.org/p/39VSojUK00H :point_left:
