package cmd

import (
	"testing"
)

func TestSanitize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"plain text", "Hello World", "Hello World"},
		{"minecraft color codes", "§aGreen §bAqua", "Green Aqua"},
		{"ansi escapes", "\x1b[31mRed\x1b[0m", "Red"},
		{"html tags", "<color=red>Red</color>", "Red"},
		{"quake color codes", "^1Red ^2Green", "Red Green"},
		{"control characters", "Hello\x00World", "HelloWorld"},
		{"whitespace collapse", "Hello    World", "Hello World"},
		{"leading/trailing whitespace", "  Hello  ", "Hello"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sanitize(tt.input)
			if got != tt.want {
				t.Errorf("sanitize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
