package applicationController

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"Melodex/src/backend/music"
	"Melodex/src/tui/sharedState"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	sharedState *sharedState.SharedState

	width  int
	height int
	volume int

	pB      progress.Model
	timeSet textinput.Model

	sleep     bool
	timeBegin time.Time
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	var parsedTime time.Duration
	if m.timeSet.Value() != "" {
		parsedTime, _ = parseUserDuration(m.timeSet.Value())
	}

	if m.sleep && time.Since(m.timeBegin) > parsedTime {
		music.PauseSong()
	}

	switch action := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = action.Width, action.Height
	case tea.KeyMsg:
		if m.sharedState.Searching {
		} else if m.sharedState.SettingTime {
			switch action.String() {
			case "esc":
				m.sharedState.SettingTime = false
				m.timeSet.Blur()
			case "enter":
				m.sharedState.SettingTime = false
				m.timeSet.Blur()
				m.sleep = !m.sleep
				if m.sleep {
					m.timeBegin = time.Now()
				}
			}
		} else {
			switch action.String() {
			case "j":
				m.volume = min(m.volume+5, 100)
				music.DecreaseVolume(5)
			case "k":
				m.volume = max(m.volume-5, 0)
				music.IncreaseVolume(5)
			case "t":
				m.sharedState.SettingTime = true
				return m, m.timeSet.Focus()
			case "m":
				m.volume = 0 // TODO
			}
		}
	}

	m.timeSet, cmd = m.timeSet.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	var targetWidth = 54
	var elements = []string{}

	if targetWidth >= m.width+1 {
		targetWidth = m.width - 2
	}

	elements = append(elements, lipg.NewStyle().Bold(true).Render("Application Controller"))
	elements = append(elements, "")
	elements = append(elements, "Volume "+m.pB.ViewAs(float64(m.volume)/100)+strconv.Itoa(m.volume)+"%")

	var parsedTime time.Duration
	if m.timeSet.Value() != "" {
		parsedTime, _ = parseUserDuration(m.timeSet.Value())
	}

	if m.sleep && !m.sharedState.SettingTime {
		elapsed := time.Since(m.timeBegin)
		remaining := parsedTime - elapsed
		m.timeSet.SetValue(formatDuration(remaining))
	} else if !m.sharedState.SettingTime {
		m.timeSet.SetValue(formatDuration(parsedTime))
	}

	sleepTimer := "Sleep: "
	sleepTimer += m.timeSet.View()
	if m.sleep {
		sleepTimer += "󰱒 "
	} else {
		sleepTimer += "󰄱 "
	}
	elements = append(elements, sleepTimer)

	final := lipg.JoinVertical(lipg.Center, elements...)

	minWidth := lipg.Width(final)
	minHeight := lipg.Height(final)

	paddingAmount := (targetWidth - minWidth) / 2
	unevenPadding := 0

	if minWidth%2 == 1 {
		unevenPadding += 1
	}

	if m.height <= 30 {
		return ""
	} else if m.width >= minWidth && m.height >= minHeight {
		content := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).
			PaddingLeft(paddingAmount).
			PaddingRight(paddingAmount + unevenPadding).
			PaddingTop(1).
			PaddingBottom(1).
			AlignVertical(lipg.Left).
			Render(final)
		return content
	} else {
		return ""
	}
}

func New(newSharedState *sharedState.SharedState) Model {
	pb := progress.New(
		progress.WithWidth(20),
		progress.WithGradient("#ff00cc", "#b800e6"),
		progress.WithoutPercentage(),
	)

	tS := textinput.New()
	tS.Width = 8
	tS.Placeholder = "30s"
	tS.Prompt = " "

	return Model{
		sharedState: newSharedState,
		pB:          pb,
		sleep:       false,
		timeSet:     tS,
	}
}

func parseUserDuration(s string) (time.Duration, error) {
	s = strings.ReplaceAll(s, " ", "")

	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}

	// Accept units that start with h, m, s (case-insensitive)
	re := regexp.MustCompile(`(?i)(\d+)([hms])`)
	matches := re.FindAllStringSubmatch(s, -1)
	if len(matches) > 0 {
		var total time.Duration
		for _, m := range matches {
			val, _ := strconv.Atoi(m[1])
			switch strings.ToLower(m[2]) {
			case "h":
				total += time.Duration(val) * time.Hour
			case "m":
				total += time.Duration(val) * time.Minute
			case "s":
				total += time.Duration(val) * time.Second
			}
		}
		return total, nil
	}

	// colon-separated (hh:mm:ss, mm:ss, ss)
	parts := strings.Split(s, ":")
	var seconds int
	switch len(parts) {
	case 3:
		h, err1 := strconv.Atoi(parts[0])
		m, err2 := strconv.Atoi(parts[1])
		sec, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			return 0, errors.New("invalid time format")
		}
		seconds = h*3600 + m*60 + sec
	case 2:
		m, err1 := strconv.Atoi(parts[0])
		sec, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return 0, errors.New("invalid time format")
		}
		seconds = m*60 + sec
	case 1:
		sec, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, errors.New("invalid time format")
		}
		seconds = sec
	default:
		return 0, errors.New("invalid time format")
	}
	return time.Duration(seconds) * time.Second, nil
}

func formatDuration(d time.Duration) string {
	totalSeconds := max(int(d.Seconds()), 0)
	h := totalSeconds / 3600
	m := (totalSeconds % 3600) / 60
	s := totalSeconds % 60
	if h > 0 {
		return pad2(h) + ":" + pad2(m) + ":" + pad2(s)
	}
	return pad2(m) + ":" + pad2(s)
}

func pad2(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}
