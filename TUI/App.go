package TUI

/*
import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"
)

// Defines variables that can be used via model, oftentimes
type model struct {
	list   table.Model
	width  int
	height int
}

// The List function that defines the properties of the table
func List() model {
	// the Columns + the tiles + the width, more should probably not be explained
	columns := []table.Column{
		{Title: "Title", Width: 40},
	}

	// these should be in the future filled automaticly by the backend, and updated when neccesary
	rows := []table.Row{
		{"ASD"},
		{"asdwqr"},
		{"rfew"},
		{"fgdgs"},
		{"qwdsa"},
		{"frsd"},
		{"h56a"},
		{"ASD"},
		{"asdwqr"},
		{"rfew"},
		{"fgdgs"},
		{"qwdsa"},
		{"frsd"},
		{"h56a"},
		{"ASD"},
		{"asdwqr"},
		{"rfew"},
		{"fgdgs"},
		{"qwdsa"},
		{"frsd"},
		{"h56a"},
		{"ASD"},
		{"asdwqr"},
		{"rfew"},
		{"fgdgs"},
		{"qwdsa"},
		{"frsd"},
		{"h56a"},
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

	return model{list: t}
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
			// FIXME MY EYES, this should be fixed to function regardless of length
			m.list.MoveUp(1000000000)
		case "end":
			// FIXME MY EYES, this should be fixed to function regardless of length
			m.list.MoveDown(1000000000)
		case "enter":
			fmt.Printf("Selected Entry")
			// m.list.
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

	// a long time ago I attempted to make scroll wheel events but tghey failed, catastrophicaly, if it can be implemented pelase do it

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
	return lipg.NewStyle().BorderStyle(lipg.RoundedBorder()).Render("PLAYLISTNAME                 󰉹 SEARCH  ") + "\n" +
		lipg.NewStyle().BorderStyle(lipg.RoundedBorder()).Render(m.list.View())
}
*/
