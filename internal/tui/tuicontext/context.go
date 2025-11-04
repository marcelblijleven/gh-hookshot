package tuicontext

import (
	"fmt"

	"github.com/marcelblijleven/gh-hookshot/internal/tui/keys"
)

const (
	deliveriesView int = iota
	detailsView
	totalColumns
)

type Context struct {
	State        int
	Version      string
	WindowWidth  int
	WindowHeight int
	Keys         keys.KeyMapping
	Theme        string

	HeaderHeight   int
	WebhooksHeight int
	FooterHeight   int

	Owner              string
	Repo               string
	SelectedWebhookID  int
	SelectedDeliveryID int

	activeColumn int
}

func (c Context) GetFullRepoName() string {
	return fmt.Sprintf("%s/%s", c.Owner, c.Repo)
}

func (c *Context) NextColumn() {
	if c.activeColumn < totalColumns-1 {
		c.activeColumn++
	}
}

func (c *Context) PreviousColumn() {
	if c.activeColumn-1 >= 0 {
		c.activeColumn--
	}
}

func (c Context) ActiveColumn() int {
	return c.activeColumn
}

func (c Context) IsDeliveriesView() bool {
	return c.ActiveColumn() == deliveriesView
}

func (c Context) IsDetailsView() bool {
	return c.ActiveColumn() == detailsView
}
