package content

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/content/tabs"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/repository"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
)

type Model struct {
	ctx    *tuicontext.Context
	height int

	tabs       tabs.Model
	webhooks   list.Model
	deliveries list.Model
	repoValid  bool
	err        error
}

func New(ctx *tuicontext.Context) Model {
	d := list.NewDefaultDelegate()
	return Model{
		ctx:        ctx,
		tabs:       tabs.New(ctx),
		webhooks:   list.New([]list.Item{}, d, 0, 0),
		deliveries: list.New([]list.Item{}, d, 0, 0),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		webhookCmd    tea.Cmd
		deliveriesCmd tea.Cmd
		tabCmd        tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = m.ctx.CalculateContentHeight(msg.Height) - lipgloss.Height(m.tabs.View())
	case repository.RepositoryDataMsg:
		if !msg.Valid {
			m.err = msg.Err
			return m, nil
		}

		m.repoValid = true
		m.height = m.ctx.CalculateContentHeight(m.ctx.WindowHeight) - lipgloss.Height(m.tabs.View())
		return m, fetchWebhooksCmd(m.ctx.Owner, m.ctx.Repo)
	case webhooksFetchMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		items := make([]list.Item, len(msg.Webhooks))
		for idx, item := range msg.Webhooks {
			items[idx] = item
		}

		log.Println(len(items), "foooooo")
		m.webhooks.SetItems(items)
	}

	m.tabs, tabCmd = m.tabs.Update(msg)
	m.webhooks, webhookCmd = m.webhooks.Update(msg)
	m.deliveries, deliveriesCmd = m.deliveries.Update(msg)

	return m, tea.Batch(tabCmd, webhookCmd, deliveriesCmd)
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Height(m.height).
		Width(m.ctx.WindowWidth).
		Align(lipgloss.Center, lipgloss.Center)

	if m.err != nil {
		return style.Render(m.err.Error())
	}

	if !m.repoValid {
		return lipgloss.JoinVertical(lipgloss.Center, m.tabs.View(), style.Render("Awaiting initial data"))
	}
	if m.tabs.ActiveTab() == tabs.WebhooksTab {
		return m.renderWebhooksView(style)
	}

	if m.tabs.ActiveTab() == tabs.DeliveriesTab {
		return m.renderDeliveriesView(style)
	}

	return lipgloss.JoinVertical(lipgloss.Center, m.tabs.View(), style.Render("Something went wrong..."))
}

func (m Model) renderWebhooksView(style lipgloss.Style) string {
	if len(m.webhooks.Items()) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, m.tabs.View(), style.Render("No webhooks found"))
	}

	return lipgloss.JoinVertical(lipgloss.Center, m.tabs.View(), m.webhooks.View())
}

func (m Model) renderDeliveriesView(style lipgloss.Style) string {
	if len(m.webhooks.Items()) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, m.tabs.View(), style.Render("No webhooks found"))
	}

	if len(m.deliveries.Items()) == 0 {
		return lipgloss.JoinVertical(lipgloss.Center, m.tabs.View(), style.Render("No webhooks found"))
	}
	return lipgloss.JoinVertical(lipgloss.Center, m.tabs.View(), m.deliveries.View())
}
