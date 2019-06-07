package errors

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func BenchmarkWithError(b *testing.B) {
	err := Wrap(errors.New("vkg"))
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", err.Error())
	}
}

func BenchmarkWithoutError(b *testing.B) {
	err := Wrap(errors.New("vkg"))
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", err)
	}
}

func BenchmarkSoul(b *testing.B) {
	err := Wrap(errors.New("vkg"))
	os.Setenv("TERM_PROGRAM", "iTerm.app")
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%+v", err)
	}
}
