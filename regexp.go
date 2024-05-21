/*
Package emojiregexp provides a regular expression that matches all Unicode emojis.
*/
package emojiregexp

import (
	_ "embed"
	"regexp"
)

//go:embed re2.txt
var re2 string

var EmojiRegexp = regexp.MustCompile(re2)
