package MusicPlayer

import (
	"Melodex/Services"
	"Melodex/TUI/SharedState"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	SharedState *SharedState.SharedState
	ProgressBar progress.Model

	Width  int
	Height int

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
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
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
		selectedPreview += "Selection Preview"
	}

	// TODO make the spaces dynamic
	control := currentTime + "                "
	if m.Shuffling {
		control += " "
	} else {
		control += " "
	}

	control += "󰒫"

	if m.Paused {

		control += "  "
	} else {
		control += "  "
	}

	control += "󰒬 "

	if m.Looping {
		control += " "
	} else {
		control += "󰑗 "
	}

	control += "              " + totalTime

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
		"\n\n\n\n\n\n",
		m.ProgressBar.ViewAs(percent),
		" ",
		control,
	)

	final := lipg.NewStyle().Padding(2).Align(lipg.Center).PaddingBottom(10).PaddingTop(5).BorderStyle(lipg.ThickBorder()).Render(content)

	if m.Width >= 56 && m.Height >= 50 {
		return final
	} else {

		minWidth := lipg.Width(final)
		minHeight := lipg.Height(final)

		if m.Width >= minWidth && m.Height >= minHeight {
			return final
		} else {
			redStyle := lipg.NewStyle().Foreground(lipg.Color("#DD0000"))
			greenStyle := lipg.NewStyle().Foreground(lipg.Color("#00DD00"))

			widthColor := redStyle
			if m.Width >= minWidth {
				widthColor = greenStyle
			}

			heightColor := redStyle
			if m.Height >= minHeight {
				heightColor = greenStyle
			}

			musicListWidth := 0

			if m.Width > 102 {
				musicListWidth = max(0, m.Width-minWidth)
			}

			// BUG Height and the associated current height do not get bold, this is very jaring to see, so it should be fixed soon
			errorContent := lipg.JoinVertical(
				lipg.Center,
				"Terminal size too small for Player:",
				lipg.JoinHorizontal(
					lipg.Left,
					"Width = ",
					widthColor.Render(strconv.Itoa(m.Width)),
					" Height = ",
					heightColor.Render(strconv.Itoa(m.Height)),
				),
				"Needed for Melodex:",
				"Width = 56 Height = 50",
			)

			leftPadding := (m.Width - lipg.Width(errorContent) - musicListWidth) / 2
			topPadding := (m.Height - lipg.Height(errorContent)) / 2

			leftPadding = max(leftPadding, 0)
			topPadding = max(topPadding, 0)

			error := lipg.NewStyle().
				PaddingTop(topPadding).
				PaddingLeft(leftPadding).
				Bold(true).
				Render(errorContent)
			return error
		}
	}
}

// NOTE When the time exceeds an hour it will just go on, this might require fixing
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
		Title:       Services.GetSelectedSong(),
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
