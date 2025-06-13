package musicPlayer

import (
	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/log"
	"Melodex/src/settings"
	"Melodex/src/tui/sharedState"
	"fmt"
	"strconv"
	"strings"
	"time"

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
		if m.sharedState.SettingTime {
		} else if m.sharedState.Searching {
		} else {
			switch action.String() {
			case "right":
				m.progress = min(m.progress+5, m.length)
				music.SetMediaPosition(music.CurrentMediaPosition() + 0.05)
			case "left":
				m.progress = max(m.progress-5, 0)
				music.SetMediaPosition(music.CurrentMediaPosition() - 0.05)
			case "b":
				music.PlayPreviousSong()
				settings.DecreaseUpdateCurrentSong()
			case "n":
				music.PlayNextSong()
				settings.IncreaseUpdateCurrentSong()
				//if m.sharedState.Shuffling {
				//	m.sharedState.Shuffling = false
				//} else {
				//	m.sharedState.Shuffling = true
				//}
				//settings.SetShuffle(m.sharedState.Shuffling)
				//log.Log.Infof("shuffled %v", m.sharedState.Shuffling)
			case ".":
				switch m.sharedState.SongOption {
				case -1:
					m.sharedState.SongOption = 0
				case 0:
					m.sharedState.SongOption = 1
				case 1:
					m.sharedState.SongOption = -1
				}

				settings.SetMode(m.sharedState.SongOption)
				log.Log.Infof("Mode: %v", m.sharedState.SongOption)
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
	m.length = timeToMinutes(settings.AppConfig.CurrentSongLength)
	return m, nil
}

func getGreeting() string {
	hour := time.Now().Hour()

	switch {
	case hour >= 5 && hour < 12:
		return "Good morning"
	case hour >= 12 && hour < 18:
		return "Good afternoon"
	case hour >= 18 && hour < 22:
		return "Good evening"
	default:
		return "Good night"
	}
}

func (m Model) View() string {
	percent := float64(m.progress) / float64(m.length)
	totalTime := formatTime(m.length)
	currentTime := formatTime(m.progress)

	var elements = []string{}

	elements = append(elements, getGreeting()+", @"+config.GetUser()+"!")

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
	m.title = "Currently Playing:\n" + settings.AppConfig.CurrentSong
	if settings.AppConfig.CurrentSong == "" {
		m.title = "Viewing Song List"
	}

	elements = append(elements, m.title+"\n\n")

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
	secs := seconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, secs)
}

func New(sharedState *sharedState.SharedState) Model {
	pb := progress.New(
		progress.WithGradient("#00FFCC", "#00b8e6"),
		progress.WithoutPercentage(),
	)
	sharedState.Paused = true

	return Model{
		sharedState: sharedState,
		length:      timeToMinutes(settings.AppConfig.SongList[0][settings.AppConfig.CurrentSongIndex]),
		progress:    0,
		progressBar: pb,
	}
}

func timeToMinutes(timeStr string) int {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return 0
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}

	return (hours * 60) + minutes
}
