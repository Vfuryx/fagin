package errorw_test

import (
	"errors"
	"fagin/pkg/errorw"
	"testing"
)

func BenchmarkErrorW(b *testing.B) {
	e1 := errors.New("1")

	err := errorw.New("2")

	for i := 0; i < b.N; i++ {
		err = errorw.Wrap(e1, err)

		_ = err.Error()
	}
}
