package sleep

import (
	"testing"
	"time"
)

// TestTmpExecuatable test time sleep
func TestTmpExecuatable(t *testing.T) {
	time.Sleep(time.Second * 15)
	t.Log("Test Done")
}
