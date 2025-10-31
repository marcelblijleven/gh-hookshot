package tuicontext

import (
	"fmt"

	"github.com/marcelblijleven/gh-hookshot/internal/tui/keys"
)

type Context struct {
	State        int
	Version      string
	WindowWidth  int
	WindowHeight int
	Keys         keys.KeyMapping

	HeaderHeight int
	FooterHeight int

	Owner string
	Repo  string
}

func (c Context) GetFullRepoName() string {
	return fmt.Sprintf("%s/%s", c.Owner, c.Repo)
}

func (c Context) CalculateContentHeight(h int) int {
	return h - c.HeaderHeight - c.FooterHeight
}
