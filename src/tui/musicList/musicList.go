package musicList

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"Melodex/src/connection"
	SharedState "Melodex/src/tui/sharedState"

	"Melodex/src/backend/music"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	sharedState *SharedState.SharedState

	list      table.Model
	searchBar textinput.Model

	originalRows []table.Row

	playlistName   string
	totalListWidth int

	width  int
	height int
}

// Init implements the tea.Model interface
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles input events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			m.list.MoveUp(1)
		case "down":
			m.list.MoveDown(1)
		case "pgup":
			m.list.MoveUp(10)
		case "pgdown":
			m.list.MoveDown(10)
		case "home":
			m.list.GotoTop()
		case "end":
			m.list.GotoBottom()
		}

		if m.sharedState.Searching {
			switch msg.String() {
			case "esc", "enter":
				m.sharedState.Searching = false
				m.searchBar.Blur()
				// return m, nil
			}
		} else {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "enter":
				m.sharedState.Paused = false
			// case " ":
			// 	music.SetVolume(50)
			// 	// music.PlaySong()
			// 	m.sharedState.Paused = false
			case " ":
				music.PauseSong()
				// m.sharedState.Paused = !music.IsPlaying()
			case "f":
				m.sharedState.Searching = true
				return m, m.searchBar.Focus()
			}
		}
	case tea.WindowSizeMsg:
		// This is to handle the window resize
		m.width, m.height = msg.Width, msg.Height
		m.list.SetHeight(m.height - 4)
		newColumns := m.list.Columns()

		elseWidth := 70
		titleWidth := max(m.width-elseWidth, 33)
		// lengthWidth := int(0.1 * float64(availableWidth))
		lengthWidth := 8

		newColumns[0].Width = titleWidth
		newColumns[1].Width = lengthWidth

		m.totalListWidth = titleWidth + lengthWidth
		m.list.SetColumns(newColumns)

	case tea.MouseMsg:
		if msg.Y >= 4 && msg.Y <= m.height-1 && msg.X >= 1 && msg.X <= m.totalListWidth+4 {
			switch tea.MouseEvent(msg).Button {
			case tea.MouseButtonWheelUp:
				m.list.MoveUp(1)
			case tea.MouseButtonWheelDown:
				m.list.MoveDown(1)
			}
			switch tea.MouseAction(msg.Action) {
			case tea.MouseAction(tea.MouseButtonLeft):
				rowIdx := msg.Y - 4
				m.list.SetCursor(rowIdx)
			}
		}
	}
	m.searchBar, cmd = m.searchBar.Update(msg)

	// NOTE case-unsensitive
	searchTerm := strings.ToLower(m.searchBar.Value())
	if searchTerm != "" {
		filteredRows := m.filterRows(searchTerm)
		m.list.SetRows(filteredRows)
	} else {
		m.list.SetRows(m.originalRows)
	}

	return m, cmd
}

func (m Model) filterRows(searchTerm string) []table.Row {
	var filteredRows []table.Row
	for _, row := range m.originalRows {
		for _, cell := range row {
			// NOTE THe search is accross all columns
			if strings.Contains(strings.ToLower(cell), searchTerm) {
				filteredRows = append(filteredRows, row)
				break
			}
		}
	}
	return filteredRows
}

// View renders the music list
func (m Model) View() string {
	searchBar := m.searchBar.View()

	// NOTE Possible chars for the diffrent filters
	//  󰉹, , , , , 󱕉, 󱕋, 󱕊, 󱕌

	// HACK This should be temporary and be in a seperate function (the rest of the function):
	padding := m.totalListWidth - lipg.Width(" "+m.playlistName) - lipg.Width(searchBar) + 4
	padding = max(padding, 0)

	headerBorder := lipg.Border{
		Top:         "━",
		Bottom:      "━",
		Left:        "┃",
		Right:       "┃",
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┣",
		BottomRight: "┫",
	}

	header := lipg.NewStyle().
		BorderStyle(headerBorder).
		Render(
			lipg.NewStyle().
				Bold(true).Render(" " +
				m.playlistName +
				strings.Repeat(" ", padding) +
				searchBar))

	musicList := lipg.NewStyle().
		BorderStyle(lipg.ThickBorder()).
		BorderTop(false).
		BorderBottom(true).
		BorderLeft(true).
		BorderRight(true).
		Render(m.list.View())

	// NOTE Idealy this number should be dynamic, but this is not necessary
	if m.width > 102 {
		return lipg.JoinVertical(lipg.Top, header, musicList)
	} else {
		return ""
	}
}

// New initializes the music list
func New(sharedState *SharedState.SharedState) Model {
	// rows := []table.Row{
	// 	{"Bohemian Rhapsody", "5:55"},
	// 	{"Imagine", "3:03"},
	// 	{"Hotel California", "6:30"},
	// 	{"Stairway to Heaven", "8:02"},
	// 	{"Smells Like Teen Spirit", "5:01"},
	// 	{"Sweet Child O' Mine", "5:56"},
	// 	{"Billie Jean", "4:54"},
	// 	{"Wonderwall", "4:18"},
	// 	{"Hey Jude", "7:11"},
	// 	{"Comfortably Numb", "6:22"},
	// }

	rows := connection.ConnectSongs()

	longestTitle, longestTime := 35, 6

	for _, row := range rows {
		longestTitle = max(longestTitle, len(row[0]))
		longestTime = max(longestTime, len(row[1]))
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

	sB := textinput.New()
	sB.Width = 20
	sB.CharLimit = 0
	sB.Placeholder = "Search"

	// FIX find a way to somehow seperate the filter and search boxes in their own borders
	sB.Prompt = " "

	trimmedPath := strings.TrimRight("", "/")
	playlistName := filepath.Base(trimmedPath)

	// Log files
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Creating the Log file failed")
	}

	sharedState.Logger = log.New(file, "List", log.LstdFlags)

	return Model{
		sharedState:    sharedState,
		list:           t,
		originalRows:   rows,
		totalListWidth: tLW,
		playlistName:   playlistName,
		searchBar:      sB,
	}
}

// defineTableStyles sets the table styles
func defineTableStyles() table.Styles {
	styles := table.DefaultStyles()
	styles.Selected = styles.Selected.
		Foreground(lipg.Color("230")).
		Background(lipg.Color("63")).
		// Foreground(lipg.Color("#111111")).
		// Background(lipg.Color("#dddddd")).
		Bold(true)
	styles.Header = styles.Header.Bold(true).Background(lipg.Color("60"))
	return styles
}
