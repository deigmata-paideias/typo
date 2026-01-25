package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/deigmata-paideias/typo/internal/types"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item struct {
	types.MatchResult
}

func (i item) Title() string { return i.Command }

func (i item) Description() string {
	return i.Desc
}

func (i item) FilterValue() string { return i.Command }

// model bubbletea
type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func NewModel(originalCmd string, matches []types.MatchResult) model {

	items := make([]list.Item, len(matches))
	for i, match := range matches {
		items[i] = item{
			MatchResult: types.MatchResult{
				Command: match.Command,
				Score:   match.Score,
				Desc:    match.Desc,
			},
		}
	}

	// Don't add original command option

	// Add cancel option
	items = append(items, item{
		MatchResult: types.MatchResult{
			Command: "Cancel",
			Score:   1000.0,
		},
	})

	l := list.New(items, itemDelegate{}, 10, 12)
	l.Title = fmt.Sprintf("Detected possible spelling error: %s", originalCmd)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return model{list: l}
}

// Init implements tea.Model interface
func (m model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model interface
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.Command
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View implements tea.Model interface
func (m model) View() string {
	if m.choice != "" {
		if m.choice == "Cancel" {
			return quitTextStyle.Render("Operation cancelled")
		}
		return quitTextStyle.Render(fmt.Sprintf("Selected: %s", m.choice))
	}

	if m.quitting {
		return quitTextStyle.Render("Operation cancelled")
	}

	return m.list.View()
}

// GetChoice returns the user's choice
func (m model) GetChoice() string {

	return m.choice
}

// itemDelegate is the list item renderer
type itemDelegate struct{}

func (d itemDelegate) Height() int { return 1 }

func (d itemDelegate) Spacing() int { return 0 }

func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	var text string
	if i.Command == "Cancel" {
		text = "Cancel"
	} else {
		text = fmt.Sprintf("%-12s %.1f%% %s", i.Command, i.Score*100, i.Desc)
		if i.Desc != "" {
			text = fmt.Sprintf("%-12s %.1f%% - %s", i.Command, i.Score*100, i.Desc)
		} else {
			text = fmt.Sprintf("%-12s %.1f%%", i.Command, i.Score*100)
		}
	}

	if index == m.Index() {
		w.Write([]byte(selectedItemStyle.Render("▸ " + text)))
		return
	}
	w.Write([]byte(itemStyle.Render("  " + text)))
}

// RunSelector runs the selector
func RunSelector(originalCmd string, matches []types.MatchResult) (string, error) {

	if len(matches) == 0 {
		return "", nil
	}

	m := NewModel(originalCmd, matches)
	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}

	if finalModel, ok := finalModel.(model); ok {

		choice := finalModel.GetChoice()
		if choice == "Cancel" || choice == "" {
			return "", nil
		}
		return choice, nil
	}

	return "", nil
}
