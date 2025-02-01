package TUI

/*
import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"
)

// Defines variables that can be used via model, oftentimes
type model struct {
	list           table.Model
	totalListWidth int
	width          int
	height         int
}

// The List function that defines the properties of the table
func List() model {
	// the Columns + the tiles + the width, more should probably not be explained
	// these should be in the future filled automaticly by the backend, and updated when neccesary
	rows := []table.Row{
		{"ASqwfwasdqwdqwdwqdq", "3:40"},
		{"asdwqr", "3:40"},
		{"rfewsd", "3:40"},
		{"fgdgsd", "3:40"},
		{"qwdsaq", "3:40"},
		{"frsdgd", "3:40"},
		{"h56aea", "3:40"},
		{"ASDtyu", "3:40"},
		{"asdwqr", "3:40"},
		{"ASDasd", "3:40"},
		{"asdwqr", "3:40"},
		{"rfewsd", "3:40"},
		{"fgdgsd", "3:40"},
		{"qwdsaq", "3:40"},
		{"frsdgd", "3:40"},
		{"h56aea", "3:40"},
		{"ASDtyu", "3:40"},
		{"asdwqr", "3:40"},
		{"ASDasd", "3:40"},
		{"asdwqr", "3:40"},
		{"rfewsd", "3:40"},
		{"fgdgsd", "3:40"},
		{"qwdsaq", "3:40"},
		{"frsdgd", "3:40"},
		{"h56aea", "3:40"},
		{"ASDtyu", "3:40"},
		{"asdwqr", "3:40"},
	}

	longestTitle := 5 + 13
	for _, row := range rows {
		if len(row[0]) > longestTitle {
			longestTitle = len(row[0])
		}
	}

	longestTime := 6
	for _, row := range rows {
		if len(row[1]) > longestTime {
			longestTime = len(row[1])
		}
	}

	columns := []table.Column{
		{Title: "Title", Width: longestTitle},
		{Title: "Length", Width: longestTime},
	}

	// here the configuration of the table gets applied
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(2),
	)

	// here the style gets applied
	t.SetStyles(defineTableStyles())

	tLW := 0
	for _, col := range columns {
		tLW += col.Width
	}

	return model{list: t, totalListWidth: tLW}
}

// the style of the table, done via lipgloss
func defineTableStyles() table.Styles {
	styles := table.DefaultStyles()
	styles.Selected = styles.Selected.
		Foreground(lipg.Color("230")).
		Background(lipg.Color("63")).
		Bold(true)
	styles.Header = styles.Header.Bold(true).Background(lipg.Color("60"))
	return styles
}

// the function run in the main.go to start the TUI
func Run() {
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

// Runs when the Bubbletea gets initialized
func (m model) Init() tea.Cmd {
	return tea.Batch(
		// Sets the TItle to MElodex, probably no changes needed in the future
		tea.SetWindowTitle("Melodex"),
	)
}

// the update function that reapeats its actions, duh
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// thist switch is to listen to msgs, they are like events that occure and there are many of them, they get defined in the docs
	switch msg := msg.(type) {
	// listens for keypresses
	case tea.KeyMsg:
		switch msg.String() {
		// exits Update
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.list.MoveUp(1)
		case "down":
			m.list.MoveDown(1)
		case "pgup":
			// THe number could use some tweaking
			m.list.MoveUp(10)
		case "pgdown":
			// THe number could use some tweaking
			m.list.MoveDown(10)
		case "home":
			m.list.GotoTop()
		case "end":
			m.list.GotoBottom()
		case "enter":
			fmt.Printf("Selected Entry")
		}

	// listens for the mouse
	case tea.MouseMsg:
		// Mouse Type is depracated but I don't know how to use the others
		switch msg.Type {
		case tea.MouseLeft:
			if msg.Y >= 5 && msg.Y < (5+m.list.Height()) {
				rowIdx := msg.Y - 5
				m.list.SetCursor(rowIdx)
			}
		case tea.MouseWheelUp:
			m.list.MoveUp(1)
		case tea.MouseWheelDown:
			m.list.MoveDown(1)
		}
	// Sets the Size of the window, via msgs
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		// the tableheight
		m.list.SetHeight((m.height / 1) - 5)
	}

	return m, cmd
}

// this is waht ultimatly decides gets to be applied to the output, via the view of the modules
func (m model) View() string {
	playlistName := "PLAYLISTNAME"
	staticText := " 󰉹 SEARCH  "

	// the len method does not count unicode characters correctly, the last number aims to offset this
	padding := m.totalListWidth - len(playlistName) - len(staticText) + 13

	if padding < 0 {
		padding = 0
	}

	rendered := lipg.NewStyle().BorderStyle(lipg.RoundedBorder()).Render(
		playlistName+strings.Repeat(" ", padding)+staticText) + "\n" +
		lipg.NewStyle().BorderStyle(lipg.RoundedBorder()).Render(m.list.View())

	return rendered
}
*/
