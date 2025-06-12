package tui

import (
	"Melodex/src/log"
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"

	ApplicationController "Melodex/src/tui/applicationController"
	MusicList "Melodex/src/tui/musicList"
	MusicPlayer "Melodex/src/tui/musicPlayer"
	SharedState "Melodex/src/tui/sharedState"

	"os"
)

type MainModel struct {
	SharedState *SharedState.SharedState

	MList       MusicList.Model
	MPlayer     MusicPlayer.Model
	AController ApplicationController.Model

	width  int
	height int
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(
		m.MList.Init(),
		m.MPlayer.Init(),
		m.AController.Init(),
	)
}

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

func Application() {
	sharedState := SharedState.GetGlobalState()

	mainModel := MainModel{
		SharedState: sharedState,

		MList:       MusicList.New(sharedState),
		MPlayer:     MusicPlayer.New(sharedState),
		AController: ApplicationController.New(sharedState),
	}

	bzone.NewGlobal()
	p := tea.NewProgram(
		mainModel,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		log.Log.Errorf("Error running program: %v", err)
		os.Exit(1)
	}
}
