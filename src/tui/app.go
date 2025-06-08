package tui

import (
	"Melodex/src/logger"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"

	ApplicationController "Melodex/src/tui/applicationController"
	MusicList "Melodex/src/tui/musicList"
	MusicPlayer "Melodex/src/tui/musicPlayer"
	SharedState "Melodex/src/tui/sharedState"

	"os"
)

// MainModel represents the primary UI model that integrates various components.
type MainModel struct {
	SharedState *SharedState.SharedState // Stores shared application state.

	MList       MusicList.Model             // Handles the music list UI.
	MPlayer     MusicPlayer.Model           // Handles the music player UI.
	AController ApplicationController.Model // Handles the application controller UI.

	width  int // Stores the window width.
	height int // Stores the window height.
}

// Init initializes the MainModel and its subcomponents.
func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.MList.Init(), m.MPlayer.Init())
}

// Update processes incoming messages and updates UI components accordingly.
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var command tea.Cmd
	var commands []tea.Cmd

	switch action := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = action.Width, action.Height

		var mList tea.Model
		mList, command = m.MList.Update(msg)
		m.MList = mList.(MusicList.Model)
		commands = append(commands, command)

		var mPlayer tea.Model
		mPlayer, command = m.MPlayer.Update(msg)
		m.MPlayer = mPlayer.(MusicPlayer.Model)
		commands = append(commands, command)

		var aController tea.Model
		aController, command = m.AController.Update(msg)
		m.AController = aController.(ApplicationController.Model)
		commands = append(commands, command)
	default:
		var mList tea.Model
		mList, command = m.MList.Update(msg)
		m.MList = mList.(MusicList.Model)
		commands = append(commands, command)

		var mPlayer tea.Model
		mPlayer, command = m.MPlayer.Update(msg)
		m.MPlayer = mPlayer.(MusicPlayer.Model)
		commands = append(commands, command)

		var aController tea.Model
		aController, command = m.AController.Update(msg)
		m.AController = aController.(ApplicationController.Model)
		commands = append(commands, command)
	}

	return m, tea.Batch(commands...)
}

// View renders the UI components in a structured layout.
func (m MainModel) View() string {
	return lipg.JoinHorizontal(
		lipg.Top,
		m.MList.View(),
		lipg.JoinVertical(
			lipg.Left,
			m.MPlayer.View(),
			m.AController.View(),
		),
	)
}

// Application initializes and runs the terminal user interface.
func Application() {
	sharedState := SharedState.GetGlobalState()

	mainModel := MainModel{
		SharedState: sharedState,

		MList:       MusicList.New(sharedState),
		MPlayer:     MusicPlayer.New(sharedState),
		AController: ApplicationController.New(*sharedState),
	}

	bzone.NewGlobal()
	p := tea.NewProgram(
		mainModel,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		logger.Log.Errorf("Error running program: %v", err)
		os.Exit(1)
	}
}
