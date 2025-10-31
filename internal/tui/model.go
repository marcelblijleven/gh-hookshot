package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/content"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/footer"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/header"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/keys"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
)

type Model struct {
	ctx     *tuicontext.Context
	header  header.Model
	content content.Model
	footer  footer.Model
}

func New(ctx *tuicontext.Context) Model {
	m := Model{
		ctx:     ctx,
		header:  header.New(ctx),
		content: content.New(ctx),
		footer:  footer.New(ctx),
	}
	return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd        tea.Cmd
		repoCmd    tea.Cmd
		headerCmd  tea.Cmd
		contentCmd tea.Cmd
		footerCmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Key pressed
		if key.Matches(msg, keys.Keys.Quit) {
			cmd = tea.Quit
		}

	case tea.WindowSizeMsg:
		// Initial window size or window resized
		m.ctx.WindowHeight = msg.Height
		m.ctx.WindowWidth = msg.Width
	}

	m.header, headerCmd = m.header.Update(msg)
	m.footer, footerCmd = m.footer.Update(msg)
	// Content height is determined by repo, header and footer so keep
	// this as last
	m.content, contentCmd = m.content.Update(msg)

	cmds := []tea.Cmd{cmd}
	cmds = append(cmds,
		repoCmd,
		headerCmd,
		footerCmd,
		contentCmd,
	)

	return m, tea.Batch(cmds...)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.header.Init(),
		m.footer.Init(),
	)
}

func (m Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.header.View(), m.content.View(), m.footer.View())
}
