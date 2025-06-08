package musicList

import (
	"path/filepath"
	"strings"

	"Melodex/src/backend/music"
	"Melodex/src/connection"
	SharedState "Melodex/src/tui/sharedState"

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

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch action := msg.(type) {
	case tea.KeyMsg:
		switch action.String() {
		case "ctrl+c", "q":
			music.Cleanup()
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

		if m.sharedState.SettingTime {
		} else {
			if m.sharedState.Searching {
				switch action.String() {
				case "esc", "enter":
					m.sharedState.Searching = false
					m.searchBar.Blur()
				}
			} else {
				switch action.String() {
				case "q":
					return m, tea.Quit
				case "enter":
					m.sharedState.Paused = false
				case " ":
					music.PauseSong()
				case "f":
					m.sharedState.Searching = true
					return m, m.searchBar.Focus()
				}
				if m.sharedState.Searching {
					switch action.String() {
					case "esc", "enter":
						m.sharedState.Searching = false
						m.searchBar.Blur()
					}
				} else {
					switch action.String() {
					case "q":
						music.Cleanup()
						return m, tea.Quit
					case "enter":
						m.sharedState.Paused = false
						connection.SetCurrentSong(strings.TrimSuffix(m.list.SelectedRow()[0], filepath.Ext(m.list.SelectedRow()[0])))
						// TODO
					case " ":
						music.PauseSong()
						// TODO
					case "f":
						m.sharedState.Searching = true
						return m, m.searchBar.Focus()

					}
				}
			}
		}
	case tea.WindowSizeMsg:
		m.width, m.height = action.Width, action.Height
		m.list.SetHeight(m.height - 4)
		newColumns := m.list.Columns()

		elseWidth := 70
		titleWidth := max(m.width-elseWidth, 33)
		lengthWidth := 8

		newColumns[0].Width = titleWidth
		newColumns[1].Width = lengthWidth

		m.totalListWidth = titleWidth + lengthWidth
		m.list.SetColumns(newColumns)

	case tea.MouseMsg:
		if action.Y >= 4 && action.Y <= m.height-1 && action.X >= 1 && action.X <= m.totalListWidth+4 {
			switch tea.MouseEvent(action).Button {
			case tea.MouseButtonWheelUp:
				m.list.MoveUp(1)
			case tea.MouseButtonWheelDown:
				m.list.MoveDown(1)
			}
			switch tea.MouseAction(action.Action) {
			case tea.MouseAction(tea.MouseButtonLeft):
				rowIdx := action.Y - 4
				m.list.SetCursor(rowIdx)
			}
		}
	}
	m.searchBar, cmd = m.searchBar.Update(msg)

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
			if strings.Contains(strings.ToLower(cell), searchTerm) {
				filteredRows = append(filteredRows, row)
				break
			}
		}
	}
	return filteredRows
}

func (m Model) View() string {
	searchBar := m.searchBar.View()
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

	if m.width > 102 {
		return lipg.JoinVertical(lipg.Top, header, musicList)
	} else {
		return ""
	}
}

// New initializes the music list
func New(sharedState *SharedState.SharedState) Model {
	rows := connection.ConnectSongs()

	longestTitle, longestTime := 35, 6

	for _, row := range rows {
		longestTitle = max(longestTitle, len(row[0]))
		longestTime = max(longestTime, len(row[1]))
	}

	columns := []table.Column{
		{Title: "Title", Width: longestTitle},
		{Title: " Length", Width: longestTime},
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

	sB.Prompt = " "
	playlistName := "Playlist: " + connection.GetPlaylistName()

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
		Bold(true)
	styles.Header = styles.Header.Bold(true).Background(lipg.Color("60"))
	return styles
}
