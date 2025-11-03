package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/deliveries"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/deliverydetail"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/footer"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/header"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/components/webhooks"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/keys"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/repository"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/styles"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
)

type Model struct {
	ctx            *tuicontext.Context
	header         header.Model
	webhooks       webhooks.Model
	deliveries     deliveries.Model
	deliveryDetail deliverydetail.Model
	footer         footer.Model
	repoValid      bool
	err            error
}

func New(ctx *tuicontext.Context) Model {
	m := Model{
		ctx:            ctx,
		header:         header.New(ctx),
		webhooks:       webhooks.New(ctx),
		deliveries:     deliveries.New(ctx),
		deliveryDetail: deliverydetail.New(ctx),
		footer:         footer.New(ctx),
	}
	return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd               tea.Cmd
		repoCmd           tea.Cmd
		headerCmd         tea.Cmd
		webhooksCmd       tea.Cmd
		deliveriesCmd     tea.Cmd
		deliveryDetailCmd tea.Cmd
		footerCmd         tea.Cmd
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
		m.header, headerCmd = m.header.Update(msg)
		m.webhooks, webhooksCmd = m.webhooks.Update(msg)
		m.footer, footerCmd = m.footer.Update(msg)

		containerWidth, containerHeight := styles.Container.GetFrameSize()

		headerView := m.header.View()
		webhooksView := m.webhooks.View()
		footerView := m.footer.View()

		// Set size on deliveries and view port here
		availHeight := m.ctx.WindowHeight - containerHeight
		availHeight -= lipgloss.Height(headerView)
		availHeight -= lipgloss.Height(webhooksView)
		availHeight -= lipgloss.Height(footerView)

		m.deliveries.SetSize(m.ctx.WindowWidth-containerWidth, availHeight)
		m.deliveryDetail.SetSize(m.ctx.WindowWidth-containerWidth, availHeight)

		return m, tea.Batch(cmd, headerCmd, webhooksCmd, deliveriesCmd, deliveryDetailCmd, footerCmd)
	case repository.RepositoryDataMsg:
		// Message received after repository.Model.Init
		if !msg.Valid {
			m.err = msg.Err
			return m, nil
		}

		m.err = nil
		m.repoValid = true
		return m, webhooks.FetchWebhooksCmd(m.ctx.Owner, m.ctx.Repo)

	}

	// Update each nested model
	m.header, headerCmd = m.header.Update(msg)
	m.footer, footerCmd = m.footer.Update(msg)
	m.webhooks, webhooksCmd = m.webhooks.Update(msg)
	m.deliveries, deliveriesCmd = m.deliveries.Update(msg)
	m.deliveryDetail, deliveryDetailCmd = m.deliveryDetail.Update(msg)

	return m, tea.Batch(cmd, repoCmd, headerCmd, footerCmd, webhooksCmd, deliveriesCmd, deliveryDetailCmd)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.header.Init(),
		m.footer.Init(),
	)
}

func (m Model) View() string {
	headerView := m.header.View()
	webhooksView := m.webhooks.View()
	footerView := m.footer.View()

	deliveries := lipgloss.JoinVertical(lipgloss.Left, styles.ActiveTitle.Render("Deliveries"), m.deliveries.View())
	deliveries = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, false, false).BorderForeground(lipgloss.Color("62")).Render(deliveries)
	detail := lipgloss.JoinVertical(lipgloss.Left, styles.ActiveTitle.Render("Detail"), m.deliveryDetail.View())

	delivery := lipgloss.JoinHorizontal(lipgloss.Bottom, deliveries, detail)

	return styles.Container.Render(lipgloss.JoinVertical(
		lipgloss.Left,
		headerView,
		webhooksView,
		delivery,
		footerView,
	),
	)
}
