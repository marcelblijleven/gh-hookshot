package util

import (
	"encoding/json"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/muesli/reflow/wordwrap"
)

// SyntaxHighlightStruct takes a JSON-marshallable struct and returns it
// as a syntax-highlighted string.
func SyntaxHighlightStruct(s interface{}, theme string, limit int) (string, error) {
	d, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}

	lexer := lexers.Get("json")
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get(theme)
	if style == nil {
		style = styles.Fallback
	}

	it, err := lexer.Tokenise(nil, string(d))
	if err != nil {
		return "", err
	}

	ww := wordwrap.NewWriter(limit)

	formatter := formatters.Get("terminal256")
	if err := formatter.Format(ww, style, it); err != nil {
		return "", err
	}

	return ww.String(), nil
}
