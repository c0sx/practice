package q2

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens []string
		want   int
	}{
		// {[]string{"2", "1", "+", "3", "*"}, 9},
		// {[]string{"4", "13", "5", "/", "+"}, 6},
		// {[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
		{[]string{"3", "-4", "+"}, -1},
	}

	for _, tt := range tests {
		if got := evalRPN(tt.tokens); got != tt.want {
			t.Errorf("evalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
		}
	}
}
