package study

import (
	"testing"
)

func TestConcurrency(t *testing.T) {
	jobCount := 100
	ans := runConcurrencyProcesses(10, jobCount)
	if len(ans) != jobCount {
		t.Error("runConcurrencyProcesses is not correct")
	}
}
