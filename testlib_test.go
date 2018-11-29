package testlib

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	h := Reverse("Hello")
	r := "olleH"
	if strings.Compare(h, r) != 0 {
		t.Errorf("Reverse failed to reverse, got: %s, wanted: %s", h, r)
	}
}
