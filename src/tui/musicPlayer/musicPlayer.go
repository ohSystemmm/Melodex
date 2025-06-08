package musicPlayer

import (
	"Melodex/src/logger"
	"Melodex/src/tui/sharedState"

	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

// Model represents the music player UI model, handling song playback and progress tracking.
type Model struct {
	sharedState *sharedState.SharedState // Stores shared application state.
	progressBar progress.Model           // Displays the song progress bar.

	width  int // Stores the terminal width.
	height int // Stores the terminal height.

	title    string // Stores the current song title.
	artist   string // Stores the current artist name.
	album    string // Stores the current album name.
	length   int    // Stores the total song length in seconds.
	progress int    // Tracks the current playback position in seconds.

	selectedPreview bool // Tracks whether a song preview is selected.
}

// Init initializes the music player model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles user input and updates the music player state.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch action := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = action.Width, action.Height
		if m.width <= 55 {
			m.progressBar.Width = m.width - 6
		} else {
			m.progressBar.Width = 50
		}

	case tea.KeyMsg:
		if m.sharedState.Searching || m.sharedState.SettingTime {
		} else {
			switch action.String() {
			case "right":
				m.progress = min(m.progress+5, m.length)
			case "left":
				m.progress = max(m.progress-5, 0)
			case "b":
				// Handle previous song logic
			case "n":
				// Handle next song logic
			case "enter":
				// Handle current song selection logic
			case ",":
				m.sharedState.Shuffling = !m.sharedState.Shuffling
				logger.Log.Infof("Shuffled %v", m.sharedState.Shuffling)
			case ".":
				switch m.sharedState.SongOption {
				case -1:
					m.sharedState.SongOption = 0
				case 0:
					m.sharedState.SongOption = 1
				case 1:
					m.sharedState.SongOption = -1
				}
				logger.Log.Infof("Mode: %v", m.sharedState.SongOption)
			case " ":
				m.sharedState.Paused = !m.sharedState.Paused
			}
		}
	}
	return m, nil
}

// View renders the music player UI, including playback controls and progress bar.
func (m Model) View() string {
	percent := float64(m.progress) / float64(m.length)
	totalTime := formatTime(m.length)
	currentTime := formatTime(m.progress)

	var elements = []string{"Welcome to Melodex!"}

	// Adds visual elements if terminal height allows
	if m.height >= 23 {
		elements = append(elements,
			" ",
			"--------------------",
			"           ████     ",
			"-----------██████---",
			"           ██    ██ ",
			"-----------██-------",
			"           ██       ",
			"-----████████-------",
			"   ██████████       ",
			"---██████████-------",
			"     ██████         ",
			"--------------------",
			" ")
	}

	controlCenter := ""
	if m.width >= 27 {
		if m.sharedState.Shuffling {
			controlCenter += "\U000F049F "
		} else {
			controlCenter += "\U000F049E "
		}
	}
	controlCenter += "󰒮 "
	switch m.sharedState.Paused {
	case true:
		controlCenter += "\U000F040A "
	case false:
		controlCenter += "\U000F03E4 "
	}
	controlCenter += "󰒭 "

	if m.width >= 28 {
		switch m.sharedState.SongOption {
		case -1:
			controlCenter += "\U000F0457"
		case 0:
			controlCenter += "\U000F0456"
		case 1:
			controlCenter += "\U000F0458"
		}
	}

	control := fmt.Sprintf("%s %s %s", currentTime, controlCenter, totalTime)

	elements = append(elements, m.title, fmt.Sprintf("%s \n %s", m.artist, m.album))
	elements = append(elements, m.progressBar.ViewAs(percent), control)

	return lipg.JoinVertical(lipg.Center, elements...)
}

// formatTime formats a given duration (in seconds) into "mm:ss" format.
func formatTime(seconds int) string {
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
}

// New initializes a new music player model with default values.
func New(sharedState *sharedState.SharedState) Model {
	pb := progress.New(
		progress.WithGradient("#00FFCC", "#00b8e6"),
		progress.WithoutPercentage(),
	)
	sharedState.Paused = true

	return Model{
		sharedState: sharedState,
		length:      181,
		progress:    50,
		progressBar: pb,
	}
}
