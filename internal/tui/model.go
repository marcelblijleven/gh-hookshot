package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	spinner        spinner.Model

	contentHeight int
	repoValid     bool
	err           error
}

func New(ctx *tuicontext.Context) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("62"))

	m := Model{
		ctx:            ctx,
		header:         header.New(ctx),
		webhooks:       webhooks.New(ctx),
		deliveries:     deliveries.New(ctx),
		deliveryDetail: deliverydetail.New(ctx),
		footer:         footer.New(ctx),
		spinner:        s,
	}
	return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd               tea.Cmd
		headerCmd         tea.Cmd
		webhooksCmd       tea.Cmd
		deliveriesCmd     tea.Cmd
		deliveryDetailCmd tea.Cmd
		footerCmd         tea.Cmd
		spinnerCmd        tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Key pressed
		if key.Matches(msg, keys.Keys.Quit) {
			cmd = tea.Quit
		}

		if key.Matches(msg, m.ctx.Keys.Left) {
			m.ctx.PreviousColumn()
		}

		if key.Matches(msg, m.ctx.Keys.Right) {
			m.ctx.NextColumn()
		}

		if m.ctx.IsDeliveriesView() {
			m.deliveries, deliveriesCmd = m.deliveries.Update(msg)
		} else if m.ctx.IsDetailsView() {
			m.deliveryDetail, deliveryDetailCmd = m.deliveryDetail.Update(msg)
		}

		m.header, headerCmd = m.header.Update(msg)
		m.footer, footerCmd = m.footer.Update(msg)
		m.webhooks, webhooksCmd = m.webhooks.Update(msg)

		if msg.String() == "?" {
			// Recalculate sizes now that show help was toggled
			m.setSizes()
		}

		// Exit here to prevent 'leaking' keystrokes to inactive columns
		return m, tea.Batch(
			cmd,
			headerCmd,
			footerCmd,
			webhooksCmd,
			deliveriesCmd,
			deliveryDetailCmd,
			spinnerCmd,
		)

	case tea.WindowSizeMsg:
		// Initial window size or window resized
		m.ctx.WindowHeight = msg.Height
		m.ctx.WindowWidth = msg.Width
		m.header, headerCmd = m.header.Update(msg)
		m.webhooks, webhooksCmd = m.webhooks.Update(msg)
		m.footer, footerCmd = m.footer.Update(msg)

		m.setSizes()

		return m, tea.Batch(
			cmd,
			headerCmd,
			footerCmd,
			webhooksCmd,
			deliveriesCmd,
			deliveryDetailCmd,
			spinnerCmd,
		)
	case repository.RepositoryDataMsg:
		// Message received after repository.Model.Init
		if !msg.Valid {
			m.err = msg.Err
			return m, nil
		}

		m.err = nil
		m.repoValid = true
		return m, webhooks.FetchWebhooksCmd(m.ctx.Owner, m.ctx.Repo)

	default:
		// Update each nested model
		m.header, headerCmd = m.header.Update(msg)
		m.footer, footerCmd = m.footer.Update(msg)
		m.webhooks, webhooksCmd = m.webhooks.Update(msg)
		m.deliveries, deliveriesCmd = m.deliveries.Update(msg)
		m.deliveryDetail, deliveryDetailCmd = m.deliveryDetail.Update(msg)
		m.spinner, spinnerCmd = m.spinner.Update(msg)

		return m, tea.Batch(
			cmd,
			headerCmd,
			footerCmd,
			webhooksCmd,
			deliveriesCmd,
			deliveryDetailCmd,
			spinnerCmd,
		)
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.header.Init(),
		m.footer.Init(),
		m.spinner.Tick,
	)
}

func (m Model) View() string {
	headerView := m.header.View()
	footerView := m.footer.View()

	if !m.repoValid || m.err != nil {
		containerWidth, containerHeight := styles.Container.GetFrameSize()
		height := m.ctx.WindowHeight - containerHeight - lipgloss.Height(headerView) - lipgloss.Height(footerView)

		var text string
		if m.err == nil {
			text = fmt.Sprintf("%s Loading data...", m.spinner.View())
		} else {
			text = m.err.Error()
		}

		content := lipgloss.Place(
			m.ctx.WindowWidth-containerWidth,
			height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				text,
			),
		)
		return styles.Container.Render(lipgloss.JoinVertical(
			lipgloss.Left,
			headerView,
			content,
			footerView,
		),
		)
	}

	webhooksView := m.webhooks.View()
	deliveriesView := m.deliveries.View()
	detailsView := m.deliveryDetail.View()

	deliveriesColumnStyle := lipgloss.NewStyle().
		Height(m.contentHeight).
		Width(styles.DeliveriesWidth).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color("62")).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				styles.GetTitleStyle(m.ctx.IsDeliveriesView()).Render("Deliveries"),
				deliveriesView,
			),
		)
	detail := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.GetTitleStyle(m.ctx.IsDetailsView()).Render("Detail"),
		detailsView,
	)

	delivery := lipgloss.JoinHorizontal(
		lipgloss.Top,
		deliveriesColumnStyle,
		detail,
	)

	return styles.Container.Render(lipgloss.JoinVertical(
		lipgloss.Left,
		headerView,
		webhooksView,
		delivery,
		footerView,
	),
	)
}

func (m *Model) setSizes() {
	containerWidth, containerHeight := styles.Container.GetFrameSize()

	headerView := m.header.View()
	webhooksView := m.webhooks.View()
	footerView := m.footer.View()

	// Set size on deliveries and view port here
	availHeight := m.ctx.WindowHeight - containerHeight
	availHeight -= lipgloss.Height(headerView)

	availHeight -= lipgloss.Height(webhooksView)
	availHeight -= lipgloss.Height(footerView)

	m.contentHeight = availHeight - 2
	m.deliveries.SetSize(m.ctx.WindowWidth-containerWidth, m.contentHeight)
	m.deliveryDetail.SetSize(m.ctx.WindowWidth-containerWidth, m.contentHeight)
}
