package musicPlayer

import (
	"Melodex/src/backend/music"
	"Melodex/src/connection"
	"Melodex/src/tui/sharedState"

	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	sharedState *sharedState.SharedState
	progressBar progress.Model

	width  int
	height int

	// songFile string
	title  string
	artist string
	album  string
	length int
	// In percent
	progress        float32
	selectedPreview bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		if m.width <= 55 {
			// m.title = strconv.Itoa(m.width - 6)
			m.progressBar.Width = (m.width - 6)
			// m.ProgressBar.Width = 21
		} else {
			m.progressBar.Width = 50
		}

	case tea.KeyMsg:
		if m.sharedState.Searching {
		} else {
			switch msg.String() {
			case "right":
				// m.progress = min(m.progress+1, m.length)
			case "left":
				// m.progress = max(m.progress-1, 0)
			case "b":

			case "n":

			case "enter":
				m.title = connection.GetCurrentSong()
				m.artist = ""

				var err error
				m.length = 100 // TODO
				// m.length, err = music.SongLength(m.songFile, m.songFile)
				// if err != nil {
				// 	m.sharedState.Logger.Println("Music Player: Failed to find the Song Length")
				// }
				m.progress, err = music.GetSongPosition()
				if err != nil {
					m.sharedState.Logger.Println("Music Player: Failed to find the Song Length")
				}
			case ",":
				m.sharedState.Shuffling = !m.sharedState.Shuffling
			case ".":
				if m.sharedState.LoopPL {
					m.sharedState.LoopPL = false
					m.sharedState.Looping = true
				} else if m.sharedState.Looping {
					m.sharedState.Looping = false
				} else {
					m.sharedState.LoopPL = true
				}
			case " ":
				if music.IsPlaying() && m.sharedState.Paused {
					m.sharedState.Paused = !m.sharedState.Paused

				} else {
					m.sharedState.Paused = !m.sharedState.Paused
				}
			case "tab":
				m.selectedPreview = !m.selectedPreview
			}

		}
	}
	return m, nil
}

func (m Model) View() string {
	percent := float64(m.progress)
	totalTime := formatTime(m.length)
	currentTime := formatTime(int(float32(m.progress) * float32(m.length)))

	var elements = []string{}

	if m.selectedPreview {
		elements = append(elements, "Currently Viewing")
	} else {
		elements = append(elements, "Playing")
	}

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

	if m.sharedState.Paused {
		controlCenter += "\U000F040A "
	} else {
		controlCenter += "\U000F03E4 "
	}

	controlCenter += "󰒭 "

	if m.width >= 28 {
		if m.sharedState.Looping {
			controlCenter += "\U000F0458"
		} else if m.sharedState.LoopPL {
			controlCenter += "\U000F0456"
		} else {
			controlCenter += "\U000F0457"
		}
	}

	// TODO make the spaces dynamic
	control := currentTime

	if m.width <= 55 {
		control += strings.Repeat(" ", max(1, m.width/2-12))
	} else {
		control += strings.Repeat(" ", 16)
	}

	control += controlCenter

	if m.width <= 55 {
		control += strings.Repeat(" ", max(1, m.width/2-13))
	} else {
		control += strings.Repeat(" ", 15)
		// control += "               "
	}

	control += totalTime

	// NOTE When enough space is available
	m.album = ""
	elements = append(elements, m.title)
	elements = append(elements, m.artist+" \n "+m.album)
	// NOTE else
	// elements = append(elements, m.title+" "+m.artist+" - "+m.album)

	elements = append(elements, m.progressBar.ViewAs(percent))
	elements = append(elements, control)

	content := lipg.JoinVertical(lipg.Center, elements...)

	final := lipg.NewStyle().Padding(2).Align(lipg.Center).BorderStyle(lipg.ThickBorder()).Render(content)

	minWidth := lipg.Width(final)
	minHeight := lipg.Height(final)

	if m.width >= minWidth && m.height >= minHeight {
		return final
	} else {
		redStyle := lipg.NewStyle().Foreground(lipg.Color("#DD0000"))
		greenStyle := lipg.NewStyle().Foreground(lipg.Color("#00DD00"))

		widthColor := redStyle
		if m.width >= minWidth {
			widthColor = greenStyle
		}

		heightColor := redStyle
		if m.height >= minHeight {
			heightColor = greenStyle
		}

		musicListWidth := 0

		if m.width > 102 {
			musicListWidth = max(0, m.width-minWidth)
		}

		// BUG Height and the associated current height do not get bold, this is very jaring to see, so it should be fixed soon
		errorContent := lipg.JoinVertical(
			lipg.Center,
			"Terminal size too small for Player:",
			lipg.JoinHorizontal(
				lipg.Left,
				"Width = ",
				widthColor.Render(strconv.Itoa(m.width)),
				" Height = ",
				heightColor.Render(strconv.Itoa(m.height)),
			),
			"Needed for Melodex:",
			"Width = "+strconv.Itoa(minWidth)+" Height = "+strconv.Itoa(minHeight),
		)

		leftPadding := (m.width - lipg.Width(errorContent) - musicListWidth) / 2
		topPadding := (m.height - lipg.Height(errorContent)) / 2

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

// NOTE When the time exceeds an hour it will just go on, this might require fixing
func formatTime(seconds int) string {
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
}

func New(sharedState *sharedState.SharedState) Model {
	pb := progress.New(
		progress.WithGradient("#00ffcc", "#00b8e6"),
		// progress.WithDefaultGradient(),
		progress.WithoutPercentage(),
	)
	sharedState.Paused = true
	sharedState.LoopPL = true

	return Model{
		sharedState: sharedState,
		// songFile:    ,
		title:       "Example Title",
		artist:      "Example Artist",
		album:       "Example Album",
		length:      181,
		progress:    53,
		progressBar: pb,
	}
}
