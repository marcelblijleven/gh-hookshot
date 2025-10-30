package header

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/styles"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
	"github.com/marcelblijleven/gh-hookshot/internal/util"
)

type Model struct {
	ctx *tuicontext.Context
}

func New(ctx *tuicontext.Context) Model {
	return Model{
		ctx: ctx,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	m.ctx.HeaderHeight = lipgloss.Height(m.View())
	return m, nil
}

func (m Model) View() string {
	logo := lipgloss.NewStyle().Bold(true).Render("Hookshot 🏹")

	version := lipgloss.NewStyle().Foreground(styles.ColorGray).Render(m.ctx.Version)
	spacing := strings.Repeat(
		" ",
		util.Max(0, m.ctx.WindowWidth-2-lipgloss.Width(logo)-lipgloss.Width(version)),
	)

	return lipgloss.NewStyle().Width(m.ctx.WindowWidth).
		PaddingLeft(1).
		PaddingRight(1).
		Render(lipgloss.JoinHorizontal(lipgloss.Center, logo, spacing, version))
}
