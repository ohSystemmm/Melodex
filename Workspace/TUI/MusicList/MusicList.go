package MusicList

import (
	"Melodex/Backend/Music"
	"Melodex/Services"
	"Melodex/TUI/SharedState"
	"github.com/hajimehoshi/oto"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	SharedState *SharedState.SharedState

	List      table.Model
	SearchBar textinput.Model

	OriginalRows []table.Row

	PlaylistName   string
	TotalListWidth int

	Width  int
	Height int

	context     *oto.Context
	musicPlayer *Music.Player
}

var tempPath = "./Backend/Music/mp3_songs/"

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
		}

		if m.SharedState.Searching {
			switch msg.String() {
			case "esc", "enter":
				m.SharedState.Searching = false
				m.SearchBar.Blur()
				// return m, nil
			}
		} else {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "enter":
				song := m.List.SelectedRow()
				Services.SetSelectedSong(song[0])
				Services.PlaySelectedSong(tempPath, song[0])
			case "f":
				m.SharedState.Searching = true
				return m, m.SearchBar.Focus()
			}
		}
	case tea.WindowSizeMsg:
		// This is to handle the window resize
		m.Width, m.Height = msg.Width, msg.Height
		m.List.SetHeight(m.Height - 4)
		newColumns := m.List.Columns()

		elseWidth := 68
		titleWidth := max(m.Width-elseWidth, 33)
		// lengthWidth := int(0.1 * float64(availableWidth))
		lengthWidth := 6

		newColumns[0].Width = titleWidth
		newColumns[1].Width = lengthWidth

		m.TotalListWidth = titleWidth + lengthWidth
		m.List.SetColumns(newColumns)

	}
	m.SearchBar, cmd = m.SearchBar.Update(msg)

	// NOTE case-unsensitive
	searchTerm := strings.ToLower(m.SearchBar.Value())
	if searchTerm != "" {
		filteredRows := m.filterRows(searchTerm)
		m.List.SetRows(filteredRows)
	} else {
		m.List.SetRows(m.OriginalRows)
	}

	return m, cmd
}

func (m Model) filterRows(searchTerm string) []table.Row {
	var filteredRows []table.Row
	for _, row := range m.OriginalRows {
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
	searchBar := m.SearchBar.View()

	// NOTE May not be needed
	// NOTE Possible chars for the diffrent filters
	// filtering := " 󰉹 "
	// filtering := " "
	// filtering := " "

	// filtering := " "
	// filtering := " "

	// filtering := "󱕉 "
	// filtering := "󱕋 "
	// filtering := "󱕊 "
	// filtering := "󱕌 "

	// HACK This should be temporary and be in a seperate function (the rest of the function):
	padding := m.TotalListWidth - lipg.Width(" "+m.PlaylistName) - lipg.Width(searchBar) + 4
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
				m.PlaylistName +
				strings.Repeat(" ", padding) +
				searchBar))

	musicList := lipg.NewStyle().
		BorderStyle(lipg.ThickBorder()).
		BorderTop(false).
		BorderBottom(true).
		BorderLeft(true).
		BorderRight(true).
		Render(m.List.View())

	// NOTE Idealy this number should be dynamic, but this is not necessary
	if m.Width > 102 {
		return lipg.JoinVertical(lipg.Top, header, musicList)
	} else {
		return ""
	}
}

// New initializes the music list
func New(sharedState *SharedState.SharedState) Model {
	//rows := []table.Row{
	//	{"Song A", "3:40"},
	//	{"Song B", "4:20"},
	//	{"Song C", "2:50"},
	//	{"Song A", "3:40"},
	//	{"Song B", "4:20"},
	//	{"Song C", "2:50"},
	//}

	rows := Services.GetSongList(tempPath)

	// longestTitle, longestTime := 98, 15
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

	trimmedPath := strings.TrimRight(tempPath, "/")
	playlistName := filepath.Base(trimmedPath)

	return Model{
		SharedState:    sharedState,
		List:           t,
		OriginalRows:   rows,
		TotalListWidth: tLW,
		PlaylistName:   playlistName,
		SearchBar:      sB,
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
