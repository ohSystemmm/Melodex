package applicationController

import (
	"log"
	"os"
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

// FIX the clock so that it counts down

type Model struct {
	sharedState sharedState.SharedState

	width  int
	height int
	volume int

	pB      progress.Model
	timeSet textinput.Model

	sleep     bool
	timebegin time.Time
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

	if m.sleep && time.Since(m.timebegin) > parsedTime {
		// TODO Don't exit, rather pause
		os.Exit(0)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		if m.sharedState.Searching {
		} else if m.sharedState.SettingTime {
			switch msg.String() {
			case "esc":
				m.sharedState.SettingTime = false
				m.timeSet.Blur()
			case "enter":
				m.sharedState.SettingTime = false
				m.timeSet.Blur()
				m.sleep = !m.sleep
				if m.sleep {
					m.timebegin = time.Now()
				}
			}
		} else {
			switch msg.String() {
			case "j":
				m.volume = min(m.volume+1, 100)
			case "k":
				m.volume = max(m.volume-1, 0)
			case "t":
				m.sharedState.SettingTime = true
				return m, m.timeSet.Focus()
			}
		}
		music.SetVolume(m.volume)
	}

	m.timeSet, cmd = m.timeSet.Update(msg)

	return m, cmd
}
func (m Model) View() string {
	targetWidth := min(54, m.width-1)


	if targetWidth >= m.width+1 {
		targetWidth = m.width - 2
	}

	elements = append(elements, lipg.NewStyle().Bold(true).Render("Application Controller"))
	elements = append(elements, "")
	elements = append(elements, "Volume "+m.pB.ViewAs(float64(m.volume)/100)+strconv.Itoa(m.volume)+"%")
	sleepTimer := "Sleeptimer: "
	sleepTimer += m.timeSet.View()
	if m.sleep {
		sleepTimer += "󰱒 "
	} else {
		sleepTimer += "󰄱 "
	}
	elements = append(elements, sleepTimer)

	// var parsedTime time.Duration
	// if m.timeSet.Value() != "" {
	// 	parsedTime, _ = parseUserDuration(m.timeSet.Value())
	// }
	// elements = append(elements, parsedTime.String())

	elements := []string{
		lipg.NewStyle().Bold(true).Render("Application Controller"),
		"Volume",
		"Sleep Timer",
	}

	final := lipg.JoinVertical(lipg.Center, elements...)

	paddingAmount := max(0, (targetWidth-lipg.Width(final))/2)
	unevenPadding := (lipg.Width(final) % 2)

	if m.width >= minWidth && m.height >= minHeight {
		content := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).
			PaddingLeft(paddingAmount).
			PaddingRight(paddingAmount + unevenPadding).
			AlignVertical(lipg.Left).
			Render(final)
		return content
	} else {


func New(sharedState sharedState.SharedState) Model {
	pb := progress.New(
		progress.WithWidth(20),
		progress.WithGradient("#ff00cc", "#b800e6"),
		progress.WithoutPercentage(),
	)

	tS := textinput.New()
	tS.Width = 8
	tS.CharLimit = 8
	// tS.Placeholder = "00:30:00"
	tS.Placeholder = "20s"
	tS.Prompt = " "

	return Model{
		sharedState: sharedState,
		pB:          pb,
		sleep:       false,
		timeSet:     tS,
	}
}

// parseUserDuration parses "1h2m3s", "hh:mm:ss", "mm:ss", or "ss" formats.
func parseUserDuration(s string) (time.Duration, error) {
	// Try Go duration format first
	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}

	// Try colon format (hh:mm:ss, mm:ss, ss)
	parts := strings.Split(s, ":")
	var seconds int
	switch len(parts) {
	case 3:
		h, err1 := strconv.Atoi(parts[0])
		m, err2 := strconv.Atoi(parts[1])
		sec, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			log.Println("Invalid time format:", s)
			return 0, err1 // just return the first error found
		}
		seconds = h*3600 + m*60 + sec
	case 2:
		m, err1 := strconv.Atoi(parts[0])
		sec, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Println("Invalid time format:", s)
			return 0, err1
		}
		seconds = m*60 + sec
	case 1:
		sec, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Println("Invalid time format:", s)
			return 0, err
		}
		seconds = sec
	default:
		log.Println("Invalid time format:", s)
		return 0, nil
	}
	return time.Duration(seconds) * time.Second, nil

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
