package deliveries

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/common"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
)

type Model struct {
	ctx               *tuicontext.Context
	deliveries        list.Model
	deliveriesFetched bool
	err               error
}

func New(ctx *tuicontext.Context) Model {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = true
	deliveries := list.New([]list.Item{}, delegate, 0, 0)
	deliveries.SetStatusBarItemName("delivery", "deliveries")
	deliveries.SetShowHelp(false)
	deliveries.SetShowTitle(false) // Handled by tui.Model

	return Model{
		ctx:        ctx,
		deliveries: deliveries,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var listCmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.ctx.Keys.Up, m.ctx.Keys.Down) {
			m.deliveries, listCmd = m.deliveries.Update(msg)

			if m.deliveries.SelectedItem() != nil {
				selectedId := m.deliveries.SelectedItem().(hookDeliveryItem).ID
				if selectedId != m.ctx.SelectedDeliveryID {
					m.ctx.SelectedDeliveryID = m.deliveries.SelectedItem().(hookDeliveryItem).ID
					return m, tea.Batch(listCmd, deliverySelectedCmd(m.ctx.SelectedWebhookID, m.ctx.SelectedDeliveryID))
				}
			}
		}
	case common.FetchDeliveriesMsg:
		m.deliveriesFetched = true
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		items := make([]list.Item, len(msg.Deliveries))
		for idx, item := range msg.Deliveries {
			items[idx] = hookDeliveryItem{HookDelivery: item}
		}

		m.deliveries.SetItems(items)

		if len(m.deliveries.Items()) == 0 {
			m.ctx.SelectedDeliveryID = 0
			return m, common.SetNoDeliveriesCmd(m.ctx.SelectedWebhookID)
		}

		if len(m.deliveries.Items()) > 0 && m.deliveries.SelectedItem() != nil {
			m.ctx.SelectedDeliveryID = m.deliveries.SelectedItem().(hookDeliveryItem).ID
			return m, deliverySelectedCmd(m.ctx.SelectedWebhookID, m.ctx.SelectedDeliveryID)
		}
	case tea.WindowSizeMsg:
		fmt.Println("foo", m.ctx.SelectedDeliveryID)
		return m, nil
	}

	m.deliveries, listCmd = m.deliveries.Update(msg)
	return m, listCmd
}

func (m Model) View() string {
	if len(m.deliveries.Items()) == 0 {
		return ""
	}
	return m.deliveries.View()
}

func (m *Model) SetSize(width, height int) {
	m.deliveries.SetSize(50, height)
}
