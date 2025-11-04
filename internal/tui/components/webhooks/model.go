package webhooks

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/bubbles-hlist/hlist"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/deliveries"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
)

type Model struct {
	ctx          *tuicontext.Context
	delegate     hlist.ItemDelegate
	hooks        hlist.Model
	hooksFetched bool
	err          error
}

func New(ctx *tuicontext.Context) Model {
	delegate := hlist.NewDefaultDelegate()
	delegate.SetHeight(3)
	delegate.SetWidth(40)

	hooks := hlist.New([]hlist.Item{}, delegate, 0, 0)
	hooks.Title = "Webhooks"
	hooks.KeyMap = KeyMap
	hooks.SetFilteringEnabled(false)
	hooks.SetShowHelp(false)
	hooks.SetShowStatusBar(false) // Number of items will be displayed in title

	return Model{ctx: ctx, delegate: delegate, hooks: hooks}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var hooksCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.hooks.SetSize(msg.Width, m.delegate.Height()+3)
		m.ctx.WebhooksHeight = lipgloss.Height(m.hooks.View())
	case tea.KeyMsg:
		// Only allow [ and ] for this model
		if m.hooksFetched && key.Matches(msg, KeyMap.CursorLeft, KeyMap.CursorRight) {
			m.hooks, hooksCmd = m.hooks.Update(msg)
			item := m.hooks.SelectedItem()
			if item != nil {
				hookID := item.(WebhookItem).ID
				if hookID != m.ctx.SelectedWebhookID {
					m.ctx.SelectedWebhookID = hookID
					return m, tea.Batch(
						hooksCmd,
						deliveries.FetchWebhookDeliveriesCmd(
							m.ctx.Owner,
							m.ctx.Repo,
							hookID,
						),
					)
				}
			}
		}
	case webhooksFetchMsg:
		m.hooksFetched = true
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		items := make([]hlist.Item, len(msg.Webhooks))
		for idx, item := range msg.Webhooks {
			items[idx] = item
		}

		m.hooks.SetItems(items)
		m.hooks.Title = fmt.Sprintf("Webhooks (%d)", len(items))

		if item := m.hooks.SelectedItem(); item != nil {
			m.ctx.SelectedWebhookID = item.(WebhookItem).ID
			return m, deliveries.FetchWebhookDeliveriesCmd(m.ctx.Owner, m.ctx.Repo, m.ctx.SelectedWebhookID)
		}

	}
	return m, hooksCmd
}

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(borderColor)
	return style.Render(m.hooks.View())
}
