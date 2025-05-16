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
		switch msg.String() {
		case "j":
			// TODO: Volume Up
		case "k":
			// TODO: Volume Down
		}
	}
	return m, nil
}

func (m Model) View() string {
	var targetWidth = 54
	var elements = []string{}

	if targetWidth >= m.width-1 {
		targetWidth = m.width - 1
	}

	elements = append(elements, lipg.NewStyle().Bold(true).Render("Application Controller"))
	elements = append(elements, "Volume")
	elements = append(elements, "Sleep Timer")

	final := lipg.JoinVertical(lipg.Center, elements...)

	minWidth := lipg.Width(final)
	minHeight := lipg.Height(final)

	paddingAmount := (targetWidth - minWidth) / 2
	unevenPadding := 0

	if minWidth%2 == 1 {
		unevenPadding += 1
	}

	if m.width >= minWidth && m.height >= minHeight {
		content := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).
			PaddingLeft(paddingAmount).
			PaddingRight(paddingAmount + unevenPadding).
			Render(final)
		return content
	} else {
		return ""
	}
}

func New() Model {
	return Model{}
}
