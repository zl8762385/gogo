package string
import (
	"testing"
)

func Test( t *testing.T ) {
	var tests = []struct {
		s, want string
	}{
		{"backward", "drawbcaB"},
		{"hello,世界", "界世，olleh"},
		{"",""},
	}

	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf( "Reverse(%q) == %q, what %q", c.s, got, c.want);
		}
	}
}