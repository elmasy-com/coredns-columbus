package columbus

import (
	"testing"

	"github.com/coredns/caddy"
)

func TestSetup(t *testing.T) {
	c := caddy.NewTestController("dns", `columbus`)
	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}

	c = caddy.NewTestController("dns", `columbus more`)
	if err := setup(c); err == nil {
		t.Fatalf("Expected errors, but got: %v", err)
	}
}
