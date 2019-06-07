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
