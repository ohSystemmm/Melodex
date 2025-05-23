package applicationController

import (
	"os"
	"strconv"
	"time"

	// "Melodex/src/tui/sharedState"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
)

type Model struct {
	width  int
	height int
	volume int

	pB progress.Model

	Sleep     bool
	timebegin time.Time
	// ready     bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	const sleepTime = time.Minute * 30

	if m.Sleep && time.Since(m.timebegin) > sleepTime {
		os.Exit(0)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			// TODO: Volume Up
			m.volume = min(m.volume+1, 100)
		case "k":
			// TODO: Volume Down
			m.volume = max(m.volume-1, 0)
		case "t":
			m.Sleep = !m.Sleep
			if m.Sleep {
				m.timebegin = time.Now()
			}
		}
	}
	return m, nil
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
	if m.Sleep {
		elements = append(elements, "Sleeptimer "+"󰱒")
	} else {
		elements = append(elements, "Sleeptimer "+"󰄱")
	}

	// +strconv.FormatFloat(m.pB.Percent(), 'f', -1, 64)
	// elements = append(elements, "Sleep Timer")

	final := lipg.JoinVertical(lipg.Center, elements...)

	minWidth := lipg.Width(final)
	minHeight := lipg.Height(final)

	paddingAmount := (targetWidth - minWidth) / 2
	unevenPadding := 0

	if minWidth%2 == 1 {
		unevenPadding += 1
	}

	if m.width >= minWidth && m.height >= minHeight {
		content := lipg.NewStyle().BorderStyle(lipg.ThickBorder()).
			PaddingLeft(paddingAmount).
			PaddingRight(paddingAmount + unevenPadding).
			AlignVertical(lipg.Left).
			Render(final)
		return content
	} else {
		return ""
	}
}

func New() Model {
	pb := progress.New(
		progress.WithWidth(20),
		progress.WithGradient("#ff00cc", "#b800e6"),
		// progress.WithDefaultGradient(),
		progress.WithoutPercentage(),
	)

	return Model{
		pB:    pb,
		Sleep: false,
	}
}
