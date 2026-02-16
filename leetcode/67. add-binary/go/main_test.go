package lc

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
	}

	for _, tt := range tests {
		got := addBinary(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("addBinary(%s, %s) = %s, want %s", tt.a, tt.b, got, tt.want)
		}
	}
}
