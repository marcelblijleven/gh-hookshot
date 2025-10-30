package footer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
)

type Model struct {
	ctx  *tuicontext.Context
	help help.Model
}

type ResizedMsg struct{}

func New(ctx *tuicontext.Context) Model {
	return Model{
		ctx:  ctx,
		help: help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	m.ctx.FooterHeight = lipgloss.Height(m.View())
	return m, nil
}

func (m Model) View() string {
	helpView := m.help.View(m.ctx.Keys)
	return helpView
}
