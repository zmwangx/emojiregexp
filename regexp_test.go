package emojiregexp_test

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/zmwangx/emojiregexp"
)

//go:embed full-emoji-list.txt
var fullEmojiList string

func Example() {
	fmt.Println(emojiregexp.EmojiRegexp.ReplaceAllString("random 🔀 text 🥰🥰😍😎 from 😤 an emoji 🤟 pasta 🍝 generator", ""))
	// Output: random  text  from  an emoji  pasta  generator
}

func TestStripAll(t *testing.T) {
	emojis := strings.Split(strings.TrimSuffix(fullEmojiList, "\n"), "\n")
	var withEmojis, withoutEmojis string
	for _, e := range emojis {
		ch := string(rune(rand.Intn(128)))
		withEmojis += ch + e
		withoutEmojis += ch
	}
	stripped := emojiregexp.EmojiRegexp.ReplaceAllString(withEmojis, "")
	// assert stripped == withoutEmojis
	if stripped != withoutEmojis {
		t.Errorf("stripped != withoutEmojis, lengths: %d vs %d", len(stripped), len(withoutEmojis))
	}
}
