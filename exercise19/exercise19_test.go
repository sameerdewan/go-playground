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

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercise19.Add(5, 5)
	}
}
