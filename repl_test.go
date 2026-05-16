package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "squirtle bulbasaur PIKACHU",
			expected: []string{"squirtle", "bulbasaur", "pikachu"},
		},
		{
			input:    "  Lugia Machamp MeWtWo",
			expected: []string{"lugia", "machamp", "mewtwo"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length does not match expected length")
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word does not match expected word")
				t.Fail()
			}
		}
	}
}
