package failure

import "testing"

func TestLogWithoutFail(t *testing.T) {
	t.Log("Log message: 1")
	t.Logf("Logf message: %d", 2)
}

func TestLogWithFailNow(t *testing.T) {
	t.Log("Log message: 3")
	t.FailNow()
	t.Logf("Logf message: %d", 4)
}

func TestLogWithFail(t *testing.T) {
	t.Log("Log message: 5")
	t.Fail()
	t.Logf("Logf message: %d", 6)
}
