package deliverydetail

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
	"github.com/marcelblijleven/gh-hookshot/internal/util/markdown"
)

type Model struct {
	ctx            *tuicontext.Context
	details        viewport.Model
	detailsFetched bool
	err            error
}

func New(ctx *tuicontext.Context) Model {
	vp := viewport.New(0, 0)
	vp.SetContent("No delivery selected")
	return Model{
		ctx:     ctx,
		details: vp,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var viewportCmd tea.Cmd

	switch msg := msg.(type) {
	case deliveryDetailFetchMsg:
		m.detailsFetched = true
		if msg.Err != nil {
			m.err = msg.Err
			m.details.SetContent(m.err.Error())
			return m, nil
		}

		md, err := markdown.StructToMarkdown(msg.DeliveryDetail)
		if err != nil {
			m.err = err
			return m, nil
		}

		m.details.SetContent(md)
	}
	m.details, viewportCmd = m.details.Update(msg)
	return m, viewportCmd
}

func (m Model) View() string {
	return m.details.View()
}

func (m *Model) SetSize(width int, height int) {
	m.details.Width = width
	m.details.Height = height
}
