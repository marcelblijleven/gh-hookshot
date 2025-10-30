package repository

import (
	"errors"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/marcelblijleven/gh-hookshot/internal/data"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/styles"
	"github.com/marcelblijleven/gh-hookshot/internal/tui/tuicontext"
	"github.com/marcelblijleven/gh-hookshot/internal/util"
)

type Model struct {
	ctx  *tuicontext.Context
	repo data.Repository
	err  error
}

func New(ctx *tuicontext.Context) Model {
	return Model{
		ctx: ctx,
	}
}

func (m Model) Init() tea.Cmd {
	fetchRepo := func() tea.Msg {
		repo, err := data.GetRepo(m.ctx.Owner, m.ctx.Repo)
		if err != nil {
			return dataFetchMsg{
				Err: err,
			}
		}

		return dataFetchMsg{
			Err:  err,
			Repo: repo,
		}
	}

	return fetchRepo
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case dataFetchMsg:
		var ret RepositoryDataMsg
		repoName := m.ctx.GetFullRepoName()

		if msg.Err != nil {
			var httpErr *api.HTTPError

			m.err = msg.Err
			if errors.As(msg.Err, &httpErr) {
				ret = RepositoryDataMsg{
					Valid: false,
					Err:   fmt.Errorf("could not retrieve repository %s (%d)", repoName, httpErr.StatusCode),
				}
			} else {
				ret = RepositoryDataMsg{
					Valid: false,
					Err:   fmt.Errorf("error occurred while retrieving repositoy %s: %s", repoName, msg.Err.Error()),
				}
			}

			return m, func() tea.Msg { return ret }
		}

		if !msg.Repo.IsAdmin() {
			ret = RepositoryDataMsg{
				Valid: false,
				Err:   fmt.Errorf("missing admin permissions for repository %s", repoName),
			}
			return m, func() tea.Msg { return ret }
		}

		m.repo = msg.Repo
		ret = RepositoryDataMsg{Valid: true, Err: nil}

		return m, func() tea.Msg { return ret }
	}

	return m, nil
}

func (m Model) View() string {
	repoStyle := lipgloss.NewStyle().Foreground(styles.ColorGray).Render(m.ctx.GetFullRepoName())
	icon := getRepoIcon(m)
	repoName := lipgloss.JoinHorizontal(lipgloss.Center, repoStyle, " ", icon)
	spacing := strings.Repeat(" ", util.Max(0, m.ctx.WindowWidth-lipgloss.Width(repoName)))

	return lipgloss.NewStyle().Width(m.ctx.WindowWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Center, repoName, spacing))
}
