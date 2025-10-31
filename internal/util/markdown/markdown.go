package markdown

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/charmbracelet/glamour"
)

var renderer glamour.TermRenderer

func init() {
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	if err != nil {
		log.Fatal(err)
	}

	renderer = *r
}

func StructToMarkdown(s interface{}) (string, error) {
	d, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", nil
	}

	out := fmt.Sprintf("```json\n%s\n```", string(d))
	r, err := renderer.Render(out)
	if err != nil {
		return "", err
	}

	return r, nil
}
