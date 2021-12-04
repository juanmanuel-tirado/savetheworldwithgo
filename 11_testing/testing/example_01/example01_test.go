package example_01

import "testing"

func TestMe(t *testing.T) {
    r := 2 + 2
    if r != 4 {
        t.Error("expected 2 got", r)
    }
}
