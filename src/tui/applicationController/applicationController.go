package applicationController

import (
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

// Model represents the application controller, managing volume and sleep timer settings.
type Model struct {
	sharedState sharedState.SharedState // Stores shared application state.

	width  int // Stores the terminal width.
	height int // Stores the terminal height.
	volume int // Stores the volume level.

	pB      progress.Model  // Progress bar for volume control.
	timeSet textinput.Model // Input field for sleep timer.

	sleep     bool      // Indicates whether the sleep timer is active.
	timeBegin time.Time // Stores the time when sleep timer was activated.
}

// Init initializes the model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update processes user inputs and updates the application controller state.
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
			}
		}
	}

	m.timeSet, cmd = m.timeSet.Update(msg)
	return m, cmd
}

// View renders the application controller UI.
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

	sleepTimer := "Sleep: " + m.timeSet.View()
	if m.sleep {
		sleepTimer += "󰱒 "
	} else {
		sleepTimer += "󰄱 "
	}
	elements = append(elements, sleepTimer)

	final := lipg.JoinVertical(lipg.Center, elements...)
	return final
}

// New initializes a new application controller model.
func New(sharedState sharedState.SharedState) Model {
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
		sharedState: sharedState,
		pB:          pb,
		sleep:       false,
		timeSet:     tS,
	}
}

// parseUserDuration parses a time string into a time.Duration.
func parseUserDuration(s string) (time.Duration, error) {
	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}

	parts := strings.Split(s, ":")
	var seconds int

	switch len(parts) {
	case 3:
		h, _ := strconv.Atoi(parts[0])
		m, _ := strconv.Atoi(parts[1])
		sec, _ := strconv.Atoi(parts[2])
		seconds = h*3600 + m*60 + sec
	case 2:
		m, _ := strconv.Atoi(parts[0])
		sec, _ := strconv.Atoi(parts[1])
		seconds = m*60 + sec
	case 1:
		sec, _ := strconv.Atoi(parts[0])
		seconds = sec
	default:
		return 0, nil
	}

	return time.Duration(seconds) * time.Second, nil
}

// formatDuration converts a time.Duration into a formatted "HH:MM:SS" string.
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

// pad2 ensures a number is formatted with two digits.
func pad2(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}
