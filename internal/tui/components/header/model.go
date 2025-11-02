package header

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/repository"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/styles"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
	"github.com/marcelblijleven/gh-hookshot/internal/util"
)

type Model struct {
	ctx        *tuicontext.Context
	repository repository.Model
}

func New(ctx *tuicontext.Context) Model {
	return Model{
		ctx:        ctx,
		repository: repository.New(ctx),
	}
}

func (m Model) Init() tea.Cmd {
	repoCmd := m.repository.Init()
	return tea.Batch(repoCmd)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var repoCmd tea.Cmd
	m.ctx.HeaderHeight = lipgloss.Height(m.View())
	m.repository, repoCmd = m.repository.Update(msg)
	return m, repoCmd
}

func (m Model) View() string {
	frame := lipgloss.NewStyle().Width(m.ctx.WindowWidth).Padding(1, 3, 1, 3)
	frameWidth, _ := frame.GetFrameSize()
	logo := lipgloss.NewStyle().Bold(true).Render("Hookshot 🏹")

	version := lipgloss.NewStyle().Foreground(styles.ColorGray).Render(m.ctx.Version)
	spacing := strings.Repeat(
		" ",
		util.Max(0, m.ctx.WindowWidth-frameWidth-lipgloss.Width(logo)-lipgloss.Width(version)),
	)
	spacing = lipgloss.NewStyle().Foreground(styles.ColorGray).Render(spacing)

	return frame.Render(lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			logo,
			spacing,
			version),
		m.repository.View()))
}
