package tabs

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
)

const (
	WebhooksTab int = iota
	DeliveriesTab
	numberOfTabs
)

var tabMap = make(map[string]int, numberOfTabs)

type Model struct {
	ctx       *tuicontext.Context
	activeTab int
}

type TabSwitchMsg int

// init is called automatically
func init() {
	tabMap = map[string]int{
		"Webhooks":   WebhooksTab,
		"Deliveries": DeliveriesTab,
	}
}

func New(ctx *tuicontext.Context) Model {
	return Model{
		ctx:       ctx,
		activeTab: WebhooksTab,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, m.tabSwitch(msg)
	}
	return m, nil
}

func (m Model) View() string {
	tabRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		renderTab("Webhooks", m.activeTab),
		renderTab("Deliveries", m.activeTab),
	)
	gap := tabGap.Render(strings.Repeat(" ", max(0, m.ctx.WindowWidth-lipgloss.Width(tabRow)-2)))
	tabRow = lipgloss.JoinHorizontal(lipgloss.Bottom, tabRow, gap)

	return tabRow
}

func (m Model) ActiveTab() int {
	return m.activeTab
}

func (m *Model) tabSwitch(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(msg, m.ctx.Keys.Left) {
		m.activeTab = getPreviousTab(numberOfTabs, m.activeTab)
	} else if key.Matches(msg, m.ctx.Keys.Right) {
		m.activeTab = getNextTab(numberOfTabs, m.activeTab)
	}
	return func() tea.Msg {
		return TabSwitchMsg(m.activeTab)
	}
}
