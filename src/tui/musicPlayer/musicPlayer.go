package musicPlayer

import (
	"Melodex/src/backend/music"
	"Melodex/src/logger"
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

	title    string
	artist   string
	album    string
	length   int
	progress int

	selectedPreview bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

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
				//music.SetMediaPosition(0) // TODO
			case "left":
				m.progress = max(m.progress-5, 0)
				//music.SetMediaPosition(0) // TODO
			case "b":
				//playlist.DecreaseIndex()
				//connection.Play(music.GetIndex())
				//playlist.GetPreviousSong() // TODO
			case "n":
				//playlist.IncreaseIndex()
				//connection.Play(music.GetIndex()) // TODO
			case "enter":
				//m.title = connection.GetCurrentSong() // TODO
			case ",":
				if m.sharedState.Shuffling {
					m.sharedState.Shuffling = false
				} else {
					m.sharedState.Shuffling = true
				}
				//connection.SetShuffle(m.sharedState.Shuffling) // TODO
				logger.Log.Infof("shuffled %v", m.sharedState.Shuffling)
			case ".":
				switch m.sharedState.SongOption {
				case -1:
					m.sharedState.SongOption = 0
				case 0:
					m.sharedState.SongOption = 1
				case 1:
					m.sharedState.SongOption = -1
				}

				//connection.SetMode(m.sharedState.SongOption) TODO
				logger.Log.Infof("Mode: %v", m.sharedState.SongOption)
				/* NOTE
				* -1 = No Repeat
				*  0 = Repeat Playlist
				*  1 = Repeat Song
				 */
			case " ":
				if music.IsPlaying() && m.sharedState.Paused {
					m.sharedState.Paused = !m.sharedState.Paused
				} else {
					m.sharedState.Paused = !m.sharedState.Paused
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	percent := float64(m.progress) / float64(m.length)
	totalTime := formatTime(m.length)
	currentTime := formatTime(m.progress)

	var elements = []string{}

	elements = append(elements, "Welcome to Melodex!")

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
	}

	control += totalTime

	elements = append(elements, m.title)
	elements = append(elements, m.artist+" \n "+m.album)

	elements = append(elements, m.progressBar.ViewAs(percent))
	elements = append(elements, control)

	content := lipg.JoinVertical(lipg.Center, elements...)
	contentHeight := lipg.Height(content)
	targetHeight := m.height - 14

	if targetHeight >= 0 && contentHeight < targetHeight {
		padTotal := targetHeight - contentHeight
		padTop := padTotal / 4
		padBottom := padTotal - padTop

		content = lipg.NewStyle().
			PaddingTop(padTop).
			PaddingBottom(padBottom).
			Render(content)
	}
	final := lipg.NewStyle().Padding(2).Align(lipg.Center).BorderStyle(lipg.ThickBorder()).Render(content)

	minWidth := lipg.Width(final)
	minHeight := lipg.Height(final) + 1

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

		style := lipg.NewStyle().
			PaddingTop(topPadding).
			PaddingLeft(leftPadding).
			Bold(true).
			Render(errorContent)
		return style
	}
}

func formatTime(seconds int) string {
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
}

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
