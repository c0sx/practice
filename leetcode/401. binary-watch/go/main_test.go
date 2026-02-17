package lc

import "testing"

func TestReadBinaryWatch(t *testing.T) {
	tests := []struct {
		turnedOn int
		want     []string
	}{
		{
			turnedOn: 1,
			want:     []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			turnedOn: 9,
			want:     []string{},
		},
	}

	for _, tt := range tests {
		got := readBinaryWatch(tt.turnedOn)
		if !equal(got, tt.want) {
			t.Errorf("readBinaryWatch(%d) = %v, want %v", tt.turnedOn, got, tt.want)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
