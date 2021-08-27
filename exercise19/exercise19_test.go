package exercise19_test

import (
	"testing"

	"github.com/sameerdewan/go-playground/exercise19"
)

func TestAdd(t *testing.T) {
	val := exercise19.Add(1, 5)
	if val != 6 {
		t.Error("Expected 6, got ", val)
	}
}

// Can run "go test -bench ." in terminal for report
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercise19.Add(5, 5)
	}
}

// Can run "go test -cover" in terminal for coverage
