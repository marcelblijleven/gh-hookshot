package keys

import "github.com/charmbracelet/bubbles/key"

type KeyMapping struct {
	Up        key.Binding
	Down      key.Binding
	Left      key.Binding
	Right     key.Binding
	HookLeft  key.Binding
	HookRight key.Binding
	Redeliver key.Binding
	Help      key.Binding
	Quit      key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view
func (k KeyMapping) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the extended help view
func (k KeyMapping) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.HookLeft, k.HookRight, k.Redeliver},
		{k.Help, k.Quit},
	}
}

var Keys = &KeyMapping{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	HookLeft: key.NewBinding(
		key.WithKeys("["),
		key.WithHelp("[", "move hook left"),
	),
	HookRight: key.NewBinding(
		key.WithKeys("]"),
		key.WithHelp("]", "move hook right"),
	),
	Redeliver: key.NewBinding(
		key.WithKeys("R"),
		key.WithHelp("R", "redeliver selected delivery"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
