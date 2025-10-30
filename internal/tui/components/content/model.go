package content

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/repository"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
	"github.com/marcelblijleven/gh-hookshot/internal/util/markdown"
)

const (
	webhooksView int = iota
	deliveriesView
	deliveryDetailView
	totalViews
)

type Model struct {
	ctx    *tuicontext.Context
	height int

	webhooks       list.Model
	deliveries     list.Model
	deliveryDetail viewport.Model

	// Sentinels
	repoValid          bool
	webhooksFetched    bool
	deliveriesFetched  bool
	selectedWebhookID  int
	selectedDeliveryID int
	err                error
}

func New(ctx *tuicontext.Context) Model {
	// Create lists
	d1 := list.NewDefaultDelegate()
	d1.SetHeight(3)
	d2 := list.NewDefaultDelegate()
	webhooksList := list.New([]list.Item{}, d1, 0, 0)
	deliveriesList := list.New([]list.Item{}, d2, 0, 0)
	detailViewport := viewport.New(0, 0)

	// Configure lists (todo: move to helper func)
	webhooksList.Title = "Webhooks"
	webhooksList.SetShowTitle(false)
	webhooksList.SetShowHelp(false)
	webhooksList.SetStatusBarItemName("webhook", "webhooks")
	deliveriesList.Title = "Deliveries"
	deliveriesList.SetShowHelp(false)
	deliveriesList.SetShowTitle(false)
	deliveriesList.SetStatusBarItemName("delivery", "deliveries")

	detailViewport.SetContent("No delivery selected")

	return Model{
		ctx:            ctx,
		webhooks:       webhooksList,
		deliveries:     deliveriesList,
		deliveryDetail: detailViewport,
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
	case tea.KeyMsg:
		if msg.String() == "enter" {
			// TODO handle per list
			m.selectedWebhookID = m.webhooks.Items()[m.webhooks.GlobalIndex()].(WebhookItem).ID
			return m, fetchWebhookDeliveriesCmd(m.ctx.Owner, m.ctx.Repo, m.selectedWebhookID)
		}
	case tea.WindowSizeMsg:
		m.height = m.ctx.CalculateContentHeight(msg.Height)
		m.webhooks.SetSize(msg.Width/3, m.height)
		m.deliveries.SetSize(msg.Width/3, m.height)
		m.deliveryDetail.Height = m.height
		m.deliveryDetail.Width = msg.Width / 3
	case repository.RepositoryDataMsg:
		if !msg.Valid {
			m.err = msg.Err
			return m, nil
		}

		m.err = nil
		m.repoValid = true
		m.height = m.ctx.CalculateContentHeight(m.ctx.WindowHeight)
		return m, fetchWebhooksCmd(m.ctx.Owner, m.ctx.Repo)
	case webhooksFetchMsg:
		m.webhooksFetched = true
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		items := make([]list.Item, len(msg.Webhooks))
		for idx, item := range msg.Webhooks {
			items[idx] = item
		}

		m.webhooks.SetItems(items)

		if len(m.webhooks.Items()) > 0 {
			m.selectedWebhookID = m.webhooks.SelectedItem().(WebhookItem).ID

			return m, fetchWebhookDeliveriesCmd(m.ctx.Owner, m.ctx.Repo, m.selectedWebhookID)

		}
	case deliveriesFetchMsg:
		m.deliveriesFetched = true
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		items := make([]list.Item, len(msg.Deliveries))
		for idx, item := range msg.Deliveries {
			items[idx] = item
		}

		m.deliveries.SetItems(items)

		if len(m.deliveries.Items()) > 0 && m.deliveries.SelectedItem() != nil {
			m.selectedDeliveryID = m.deliveries.SelectedItem().(HookDeliveryItem).ID
			return m, fetchWebhookDeliveryDetailCmd(m.ctx.Owner, m.ctx.Repo, m.selectedWebhookID, m.selectedDeliveryID)
		}
	case deliveryDetailFetchMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		md, err := markdown.StructToMarkdown(msg.DeliveryDetail)
		if err != nil {
			m.err = err
			return m, nil
		}
		m.deliveryDetail.SetContent(md)

	}

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
		return lipgloss.JoinVertical(lipgloss.Center, style.Render("Awaiting initial data"))
	}

	return containerStyle.Render(m.renderColumns())
}

func (m Model) renderColumns() string {
	return lipgloss.JoinHorizontal(lipgloss.Left, m.renderWebhookColumn(), m.renderDeliveriesColumn(), m.renderDetailColumn())
}

func (m Model) renderWebhookColumn() string {
	title := "Webhooks"
	width, _, _ := m.columnWidths()
	column := lipgloss.NewStyle().Width(width)
	return lipgloss.JoinVertical(lipgloss.Left, titleStyle.Render(title), column.Render(m.webhooks.View()))
}

func (m Model) renderDeliveriesColumn() string {
	title := "Deliveries"
	_, width, _ := m.columnWidths()
	column := lipgloss.NewStyle().Width(width)
	return lipgloss.JoinVertical(lipgloss.Left, titleStyle.Render(title), column.Render(m.deliveries.View()))
}

func (m Model) renderDetailColumn() string {
	title := "Delivery detail"
	_, _, width := m.columnWidths()
	column := lipgloss.NewStyle().Width(width)
	return lipgloss.JoinVertical(lipgloss.Left, titleStyle.Render(title), column.Render(m.deliveryDetail.View()))
}

func (m Model) columnWidths() (int, int, int) {
	containerWidth, _ := containerStyle.GetFrameSize()
	windowWidth := m.ctx.WindowWidth - containerWidth
	parts := windowWidth / 5

	return parts * 2, parts * 1, parts * 3
}
