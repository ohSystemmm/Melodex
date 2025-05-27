package applicationController

import (
	"Melodex/src/backend/music"

	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	width  int
	height int
	volume int
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
			m.volume += 5
		case "k":
			m.volume -= 5
		}
		music.SetVolume(m.volume)
	}
	return m, nil
}
func (m Model) View() string {
	targetWidth := min(54, m.width-1)

	elements := []string{
		lipg.NewStyle().Bold(true).Render("Application Controller"),
		"Volume",
		"Sleep Timer",
	}

	final := lipg.JoinVertical(lipg.Center, elements...)

	paddingAmount := max(0, (targetWidth-lipg.Width(final))/2)
	unevenPadding := (lipg.Width(final) % 2)

	if m.width < lipg.Width(final) || m.height < lipg.Height(final) {
		return ""
	}

	return lipg.NewStyle().
		BorderStyle(lipg.ThickBorder()).
		PaddingLeft(paddingAmount).
		PaddingRight(paddingAmount + unevenPadding).
		Render(final)
}
func New() Model {
	return Model{}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
