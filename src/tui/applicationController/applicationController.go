package applicationController

import (
	// "Melodex/TUI/SharedState"

	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	width  int
	height int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
	}
	return m, nil
}

func (m Model) View() string {
	var elements = []string{}

	elements = append(elements, lipg.NewStyle().Bold(true).Render("Application Controller"))
	elements = append(elements, "Settings")
	elements = append(elements, "Select Playlist")
	elements = append(elements, "Edit Tags")

	final := lipg.JoinVertical(lipg.Center, elements...)

	minWidth := lipg.Width(final)
	minHeight := lipg.Height(final)

	if m.width >= minWidth && m.height >= minHeight {
		content := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).
			PaddingLeft(16).
			PaddingRight(16).
			Render(final)
		return content
	} else {
		return ""
	}
}

func New() Model {
	return Model{}
}
