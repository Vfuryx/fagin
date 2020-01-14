package status

import "testing"

func TestStatus(t *testing.T) {
	t.Log(Status().All())
	t.Log(Status().Disable)
	t.Log(Status().Get(Active))
	t.Log(Status().IsExist(Active))
}

