package TUI

import (
	// "strconv"
	// "strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

// Model struct
type Model struct {
	List           table.Model
	PlayListName   string
	TotalListWidth int
	width          int
	height         int
}

// Init implements the tea.Model interface
func (m Model) Init() tea.Cmd {
	return nil // No initial command needed
}

// Update handles input events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.List.MoveUp(1)
		case "down":
			m.List.MoveDown(1)
		case "pgup":
			m.List.MoveUp(10)
		case "pgdown":
			m.List.MoveDown(10)
		case "home":
			m.List.GotoTop()
		case "end":
			m.List.GotoBottom()
		case "enter":
			// Handle selection
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.List.SetHeight(m.height - 5)

		newColumns := m.List.Columns()
		availableWidth := m.width - 10
		titleWidth := int(0.3 * float64(availableWidth))
		lengthWidth := int(0.1 * float64(titleWidth))

		newColumns[0].Width = titleWidth
		newColumns[1].Width = lengthWidth

		m.TotalListWidth = titleWidth + lengthWidth
		m.List.SetColumns(newColumns)

	}

	return m, cmd
}

// View renders the music list
func (m Model) View() string {
	playlistName := "PLAYLISTNAME"
	staticText := " 󰉹 SEARCH  "

	padding := m.TotalListWidth - lipg.Width(playlistName) - lipg.Width(staticText) + 4
	if padding < 0 {
		padding = 0
	}

	// header := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(
	// 	playlistName + strings.Repeat(" ", padding) + staticText)

	// funcmenul := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(m.List.View())
	// funcmenum := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(m.List.View())
	// funcmenur := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(m.List.View())

	// return lipg.JoinHorizontal(lipg.Top, lipg.JoinVertical(lipg.Top, header, musicList), musicPlayer)
	return ""
}

// MList initializes the music list
func MList(m Model) Model {
	rows := []table.Row{
		{"Song A", "3:40"},
		{"Song B", "4:20"},
		{"Song C", "2:50"},
		{"Song A", "3:40"},
		{"Song B", "4:20"},
		{"Song C", "2:50"},
	}

	longestTitle, longestTime := 20, 6

	for _, row := range rows {
		if len(row[0]) > longestTitle {
			longestTitle = len(row[0])
		}
		if len(row[1]) > longestTime {
			longestTime = len(row[1])
		}
	}

	columns := []table.Column{
		{Title: "Title", Width: longestTitle},
		{Title: "Length", Width: longestTime},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	t.SetStyles(defineTableStyles())

	tLW := 0
	for _, col := range columns {
		tLW += col.Width
	}

	return Model{List: t, TotalListWidth: tLW}
}

// defineTableStyles sets the table styles
func defineTableStyles() table.Styles {
	styles := table.DefaultStyles()
	styles.Selected = styles.Selected.
		Foreground(lipg.Color("230")).
		Background(lipg.Color("63")).
		Bold(true)
	styles.Header = styles.Header.Bold(true).Background(lipg.Color("60"))
	return styles
}
