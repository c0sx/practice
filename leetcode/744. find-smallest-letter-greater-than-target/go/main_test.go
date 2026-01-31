package lc

import (
	"testing"
)

func TestNextGreaterLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'a',
			want:    'c',
		},
		{
			letters: []byte{'c', 'f', 'j'},
			target:  'c',
			want:    'f',
		},
		{
			letters: []byte{'x', 'x', 'y', 'y'},
			target:  'z',
			want:    'x',
		},
	}

	for _, tt := range tests {
		res := nextGreatestLetter(tt.letters, tt.target)
		if res != tt.want {
			t.Errorf("nextGreatestLetter(%c, %c) = %c; want %c", tt.letters, tt.target, res, tt.want)
		}
	}
}
