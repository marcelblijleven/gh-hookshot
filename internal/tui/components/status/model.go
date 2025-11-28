package status

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
)

var statusStyle = lipgloss.NewStyle().
	Padding(1).
	Height(1).
	Background(lipgloss.Color("62")).
	Foreground(lipgloss.Color("230"))

type Model struct {
	ctx    *tuicontext.Context
	status string
}

func New(ctx *tuicontext.Context) Model {
	return Model{
		ctx: ctx,
	}
}

// Init satisfies tea.Model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update satisfies tea.Model
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.status = string(msg)
		return m, statusTick()
	case statusTickMsg:
		m.status = ""
	}
	return m, nil
}

// View satisfies tea.Model
func (m Model) View() string {
	if m.status == "" {
		return ""
	}

	return statusStyle.Render(m.status)
}
