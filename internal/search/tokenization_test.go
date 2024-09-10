package search

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expectedMap map[string]int
	}{
		{
			name:    "Single sentence",
			content: "Hello, World, I'm so happy to see you again, world!",
			expectedMap: map[string]int{
				"hello": 1,
				"world": 2,
				"happy": 1,
				"see":   1,
				"again": 1,
			},
		},
		{
			name:        "Stop words excluding",
			content:     "and, or, but, nor, so, yet, for, in, on, at, by, with, about, against, between, into, through, during, before, after, above, below, to, from, up, down, under, over, the, a, an, he, she, it, they, we, you, i, me, him, her, us, them, is, are, was, were, be, been, am, have, has, had, do, does, did, will, would, shall, should, can, could, may, might, must, this, that, these, those, my, your, his, its, our, their, of, if, then, there, here, when, where, why, how, which, no, not, neither, never, none, very, too, quite, rather, almost, just, only, m",
			expectedMap: map[string]int{},
		},
		{
			name:        "Empty content",
			content:     "",
			expectedMap: map[string]int{},
		},
		{
			name:    "Special characters and regular words",
			content: "Wait—what are you doing?... - she said. Stop! The price is $5.99 for these beautiful 2 items. She said, 'It's a beautiful day! email at example@test.com. The temperature dropped by 20%!",
			expectedMap: map[string]int{
				"wait":        1,
				"what":        1,
				"doing":       1,
				"said":        2,
				"stop":        1,
				"price":       1,
				"items":       1,
				"beautiful":   2,
				"day":         1,
				"email":       1,
				"temperature": 1,
				"com":         1,
				"dropped":     1,
				"example":     1,
				"test":        1,
			},
		},
		{
			name:        "Special characters only",
			content:     "!@#$%^&*()_+={}[]:;\"'<>,.?/`~ ",
			expectedMap: map[string]int{},
		},
		{
			name:    "Sentences with Non-ASCII Characters",
			content: "Hello, こんにちは, Привет, 안녕하세요! I'm happy 😊, sad 😢, and confused 🤔!",
			expectedMap: map[string]int{
				"hello":    1,
				"こんにちは":    1,
				"привет":   1,
				"안녕하세요":    1,
				"happy":    1,
				"sad":      1,
				"confused": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := tt.content

			words := tokenize(content)

			if len(words) != len(tt.expectedMap) {
				fmt.Println(words)
				t.Errorf("Expected %d tokens, got %d", len(tt.expectedMap), len(words))
			}

			for word := range words {
				fmt.Println(word)
				if words[word] != tt.expectedMap[word] {
					t.Errorf("Expected %d frequency, got %d", tt.expectedMap[word], words[word])
				}
			}
		})
	}
}
