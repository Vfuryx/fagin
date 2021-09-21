package enums

import (
	"fagin/pkg/enum"
	"testing"
)

func TestStatus(t *testing.T) {
	t.Log(AllStatus())
	t.Log(StatusDisable)
	t.Log(StatusActive.Get())
	t.Log(enum.IsExist(StatusActive))
}
