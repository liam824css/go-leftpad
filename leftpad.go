package leftpad

import (
	"fmt"
	"strings"
)

func Leftpad(text string, pad_length int,pad_letter) {
	const text_size int = len(text)
	const pad_size int = len(pad_length)
	const text_pad_diff int = text_size - pad_size

	if pad_letter == nil {
		pad_letter = "0"
	}

	if text_size > pad_size {
		return text
	}

	if text_size < pad_size {
		return strings.Repeat(pad_letter,text_pad_diff) + text
	}
}
