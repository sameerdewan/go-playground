package exercise19

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
