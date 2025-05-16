package funcMenu

import (
	// "strconv"
	"Melodex/src/tui/sharedState"

	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

// Model struct
type Model struct {
	SharedState *sharedState.SharedState
	Button1     bool
	Button2     bool
	Button3     bool

	TotalListWidth int
	List           table.Model
	Width          int
	Height         int
	SearchBar      textinput.Model
	Sleep          bool
	time           time.Duration
	timebegin      time.Time
	ready          bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

/* func (m *Background) Init() tea.Cmd {
	return nil
} */

// Update handles input events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.time = 30 * time.Minute
	if time.Since(m.timebegin) > m.time {
		os.Exit(0)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "s":
			m.SharedState.Searching = true
		case "p": //button 2-

		case "t":
			m.Sleep = !m.Sleep
			if m.Sleep {
				m.timebegin = time.Now()
			}
		}
		if m.SharedState.Searching {
			switch msg.String() {
			case "esc", "enter":
				m.SharedState.Searching = false
				m.SearchBar.Blur()
			}
		} else {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "enter":
				// const content = "eeee"
				// NOTE Handle selection
			case "f":
				m.SharedState.Searching = true
				return m, m.SearchBar.Focus()
			}
		}

	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height

		availableWidth := m.Width - 10
		titleWidth := int(0.3 * float64(availableWidth))
		lengthWidth := int(0.1 * float64(titleWidth))
		m.TotalListWidth = titleWidth + lengthWidth
	}

	m.SearchBar, cmd = m.SearchBar.Update(msg)

	return m, cmd
}

// View renders the music list
func (m Model) View() string {
	searchBar := m.SearchBar.View()

	content := ""

	if m.Sleep {
		content = lipg.JoinVertical(
			lipg.Center,
			"󰄱",
		)
	} else {
		content = lipg.JoinVertical(
			lipg.Center,
			"󰱒",
		)
	}

	content2 := lipg.JoinVertical(
		lipg.Center,
	)
	final := lipg.NewStyle().Padding(2).Align(lipg.Center).BorderStyle(lipg.ThickBorder()).Render(content)
	final1 := lipg.JoinHorizontal(
		lipg.Center,
		final,
		lipg.NewStyle().Padding(2).Align(lipg.Center).BorderStyle(lipg.ThickBorder()).Render(searchBar),
	)
	final2 := lipg.JoinHorizontal(
		lipg.Center,
		final1,
		lipg.NewStyle().Padding(2).Align(lipg.Center).BorderStyle(lipg.ThickBorder()).Render(content2),
	)

	// header := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(
	// 	playlistName + strings.Repeat(" ", padding) + staticText)

	// funcmenul := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(m.List.View())
	// funcmenum := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(m.List.View())
	// funcmenur := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).Render(m.List.View())

	// return lipg.JoinHorizontal(lipg.Top, lipg.JoinVertical(lipg.Top, header, musicList), musicPlayer)
	return final2
}
func New(sharedState *sharedState.SharedState) Model {
	sB := textinput.New()
	sB.Width = 20
	sB.CharLimit = 0
	sB.Placeholder = "Directory"

	return Model{
		SharedState: sharedState,
		SearchBar:   sB,
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
