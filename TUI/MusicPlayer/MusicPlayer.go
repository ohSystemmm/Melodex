package MusicPlayer

import (
	"Melodex/TUI/SharedState"
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	SharedState *SharedState.SharedState
	ProgressBar progress.Model

	Title           string
	Artist          string
	Album           string
	Length          int
	Progress        int
	Paused          bool
	Shuffling       bool
	Looping         bool
	SelectedPreview bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// m.progressBar.Width = msg.Width - 10 // Adjust width based on window size
	case tea.KeyMsg:
		if m.SharedState.Searching {
		} else {
			switch msg.String() {
			case "right":
				m.Progress = min(m.Progress+1, m.Length)
			case "left":
				m.Progress = max(m.Progress-1, 0)
			case "shift+right":
				// TODO Next Song
				m.Progress = min(m.Progress+1, m.Length)
			case "shift+left":
				// TODO Previous Song
				m.Progress = max(m.Progress-1, 0)
			case " ":
				m.Paused = !m.Paused
			case ",":
				m.Shuffling = !m.Shuffling
			case ".":
				m.Looping = !m.Looping
			case "tab":
				m.SelectedPreview = !m.SelectedPreview
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	percent := float64(m.Progress) / float64(m.Length)
	currentTime := formatTime(m.Progress)
	totalTime := formatTime(m.Length)

	selectedPreview := ""
	if m.SelectedPreview {
		selectedPreview += "Currently Playing"
	} else {
		selectedPreview += "Selected Preview"
	}

	// TODO make the spaces dynamic
	control := currentTime + "                 "
	if m.Shuffling {
		control += " "
	} else {
		control += " "
	}

	if m.Paused {
		control += "  "
	} else {
		control += "  "
	}

	if m.Looping {
		control += " "
	} else {
		control += "󰑗 "
	}

	control += "                " + totalTime

	content := lipg.JoinVertical(
		lipg.Center,
		selectedPreview,

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
		" ",

		m.Title,
		" ",
		m.Artist+" - "+m.Album,
		" ",
		" ",
		m.ProgressBar.ViewAs(percent),
		control,
	)

	final := lipg.NewStyle().Padding(2).Align(lipg.Center).BorderStyle(lipg.ThickBorder()).Render(content)

	return final
}

// NOTE When the time exceeds an hour it will just ho on, this might require fixing
func formatTime(seconds int) string {
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
}

func New(sharedState *SharedState.SharedState) Model {
	pb := progress.New(
		progress.WithGradient("#00ffcc", "#00b8e6"),
		// progress.WithDefaultGradient(),
		progress.WithoutPercentage(),
	)
	// NOTE Sets the progressbar width
	// TODO should be dynamic in the future
	pb.Width = 50

	return Model{
		SharedState: sharedState,
		Title:       "Example Title",
		Artist:      "Example Artist",
		Album:       "Example Album",
		Length:      181,
		Progress:    53,
		ProgressBar: pb,
		Paused:      false,
		Shuffling:   false,
		Looping:     false,
	}
}
