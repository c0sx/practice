package leetcode966

import "testing"

func TestSpellchecker(t *testing.T) {
	tests := []struct {
		wordlist []string
		queries  []string
		expected []string
	}{
		{
			wordlist: []string{"KiTe", "kite", "hare", "Hare"},
			queries:  []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			expected: []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
		{
			wordlist: []string{"yellow"},
			queries:  []string{"YellOw"},
			expected: []string{"yellow"},
		},
	}

	for _, test := range tests {
		result := spellchecker(test.wordlist, test.queries)
		if len(result) != len(test.expected) {
			t.Errorf("spellchecker(%v, %v) = %v; want %v", test.wordlist, test.queries, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("spellchecker(%v, %v) = %v; want %v", test.wordlist, test.queries, result, test.expected)
				break
			}
		}
	}
}
