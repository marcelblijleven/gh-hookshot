package webhooks

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/marcelblijleven/bubbles-hlist/hlist"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/keys"
)

var KeyMap = hlist.KeyMap{
	// Browsing.
	CursorLeft:  keys.Keys.HookLeft,
	CursorRight: keys.Keys.HookRight,
	PrevPage: key.NewBinding(
		key.WithKeys("<"),
		key.WithHelp("<", "prev page"),
	),
	NextPage: key.NewBinding(
		key.WithKeys(">"),
		key.WithHelp(">", "next page"),
	),
	GoToStart: key.NewBinding(
		key.WithKeys("home", "g"),
		key.WithHelp("g/home", "go to start"),
	),
	GoToEnd: key.NewBinding(
		key.WithKeys("end", "G"),
		key.WithHelp("G/end", "go to end"),
	),
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	),
	ClearFilter: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "clear filter"),
	),

	// Filtering.
	CancelWhileFiltering: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel"),
	),
	AcceptWhileFiltering: key.NewBinding(
		key.WithKeys("enter", "tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
		key.WithHelp("enter", "apply filter"),
	),

	// Toggle help.
	ShowFullHelp: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "more"),
	),
	CloseFullHelp: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "close help"),
	),

	// Quitting.
	Quit: key.NewBinding(
		key.WithKeys("q", "esc"),
		key.WithHelp("q", "quit"),
	),
	ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),
}
