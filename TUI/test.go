package TUI

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"
)

var TableStyle = func(row, col int) lipgloss.Style {
	switch {
	case row == -1:
		return lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Center)
	case row%2 == 0:
		return lipgloss.NewStyle().Padding(0, 1)
	default:
		return lipgloss.NewStyle().Padding(0, 1)
	}
}

type model struct {
	list   table.Model
	width  int
	height int
	index  int8
}

func List() model {
	columns := []table.Column{
		{Title: "Title", Width: 40},
	}
	rows := []table.Row{
		{"---"},
		{"***"},
		{"BEEP"},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(2),
	)
	t.SetStyles(defineTableStyles())

	return model{list: t}
}
func defineTableStyles() table.Styles {
	styles := table.DefaultStyles()
	styles.Selected = styles.Selected.
		Foreground(lipgloss.Color("230")).
		Background(lipgloss.Color("63")).
		Bold(true)
	styles.Header = styles.Header.Bold(true).Background(lipgloss.Color("60"))
	return styles
}
func Runb() {
	bzone.NewGlobal()
	// the bubbletea "Heart", here the properties of the TUI get defined and the Modules get included
	p := tea.NewProgram(
		List(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)
	// WHen an Error occures, exits the program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
func (m model) Init() tea.Cmd {
	return tea.Batch(
		// Sets the TItle to MElodex, probably no changes needed in the future
		tea.SetWindowTitle("Melodex"),
	)
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// thist switch is to listen to msgs, they are like events that occure and there are many of them, they get defined in the docs
	switch msg := msg.(type) {
	// listens for keypresses
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.list.MoveUp(1)
		case "down":
			m.list.MoveDown(1)
		}

	// listens for the mouse
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			// thi9 is currently to set the row in the list via mouse when to lazy for keys
			if msg.Y >= 2 && msg.Y < (2+m.list.Height()) {
				rowIdx := msg.Y - 2
				m.list.SetCursor(rowIdx)
			}
		}

	// Sets the Size of the window, via msgs
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		// the tableheight
		m.list.SetHeight((m.height / 1) - 2)
	}

	return m, cmd
}
func (m model) View() string {
	//return lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Render(m.list.View())

	return lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Align(lipgloss.Right).Render(m.list.View()), lipgloss.JoinHorizontal(lipgloss.Bottom)
}
