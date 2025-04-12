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
	content := lipg.NewStyle().
		Bold(true).
		Render("Appication Controller")

	return content
}

func New() Model {
	return Model{}
}
