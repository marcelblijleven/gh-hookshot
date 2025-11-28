package deliverydetail

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/common"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/styles"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
	"github.com/marcelblijleven/gh-hookshot/internal/util"
)

type Model struct {
	ctx            *tuicontext.Context
	details        viewport.Model
	detailsFetched bool
	rawDetails     *hookDeliveryDetailItem
	err            error
}

const noDeliverySelected = "No delivery selected"

var frame = lipgloss.NewStyle().Padding(1)

func New(ctx *tuicontext.Context) Model {
	vp := viewport.New(0, 0)
	vp.SetContent(noDeliverySelected)

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
	case common.NoDeliveriesMsg:
		m.rawDetails = nil
		m.details.SetContent(noDeliverySelected)
		return m, nil

	case common.FetchDeliveryDetailMsg:
		m.detailsFetched = true
		if msg.Err != nil {
			m.err = msg.Err
			m.details.SetContent(m.err.Error())
			return m, nil
		}

		details, err := util.SyntaxHighlightStruct(msg.DeliveryDetail, m.ctx.Theme, m.contentWidth())
		if err != nil {
			m.err = err
			return m, nil
		}

		m.rawDetails = &hookDeliveryDetailItem{HookDeliveryDetail: msg.DeliveryDetail}
		m.details.SetContent(details)
		m.details.SetYOffset(0)
	}
	m.details, viewportCmd = m.details.Update(msg)
	return m, viewportCmd
}

func (m Model) View() string {
	return frame.Render(m.details.View())
}

func (m *Model) SetSize(width int, height int) {
	m.details.Width = width
	m.details.Height = height

	if m.rawDetails == nil {
		m.details.SetContent(noDeliverySelected)
		return
	}

	details, err := util.SyntaxHighlightStruct(m.rawDetails, m.ctx.Theme, m.contentWidth())
	if err != nil {
		m.err = err
	}

	m.details.SetContent(details)
}

func (m Model) contentWidth() int {
	frameWidth, _ := frame.GetFrameSize()
	return m.details.Width - frameWidth - styles.DeliveriesWidth
}
