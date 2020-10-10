package failure

import "testing"

func TestWithFatal(t *testing.T) {
	t.Fatal("Log message: 3")
	t.Logf("Logf message: %d", 4)
}

func TestWithError(t *testing.T) {
	t.Error("Log message: 5")
	t.Logf("Logf message: %d", 6)
}
