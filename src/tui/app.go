package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	lipg "github.com/charmbracelet/lipgloss"
	bzone "github.com/lrstanley/bubblezone"

<<<<<<<< HEAD:src/tui/app.go
	ApplicationController "Melodex/src/tui/applicationController"
	MusicList "Melodex/src/tui/musicList"
	MusicPlayer "Melodex/src/tui/musicPlayer"
	SharedState "Melodex/src/tui/sharedState"
========
	"Melodex/backend/music"
	"Melodex/tui/musicList"
	"Melodex/tui/musicPlayer"
	sharedState "Melodex/tui/sharedState"
>>>>>>>> main:Workspace/tui/app.go

	"fmt"
	"log"
	"os"
)

type MainModel struct {
	sharedState *sharedState.SharedState

<<<<<<<< HEAD:src/tui/app.go
	MList       MusicList.Model
	MPlayer     MusicPlayer.Model
	AController ApplicationController.Model
========
	MList   musicList.Model
	MPlayer musicPlayer.Model
>>>>>>>> main:Workspace/tui/app.go

	width  int
	height int
}

var (
	m MainModel
)

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.MList.Init(), m.MPlayer.Init())
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height

		var mList tea.Model
		mList, cmd = m.MList.Update(msg)
		m.MList = mList.(musicList.Model)
		cmds = append(cmds, cmd)

		var mPlayer tea.Model
		mPlayer, cmd = m.MPlayer.Update(msg)
		m.MPlayer = mPlayer.(musicPlayer.Model)
		cmds = append(cmds, cmd)

		var aController tea.Model
		aController, cmd = m.AController.Update(msg)
		m.AController = aController.(ApplicationController.Model)
		cmds = append(cmds, cmd)
	default:
		var mList tea.Model
		mList, cmd = m.MList.Update(msg)
		m.MList = mList.(musicList.Model)
		cmds = append(cmds, cmd)

		var mPlayer tea.Model
		mPlayer, cmd = m.MPlayer.Update(msg)
		m.MPlayer = mPlayer.(musicPlayer.Model)
		cmds = append(cmds, cmd)

		var aController tea.Model
		aController, cmd = m.AController.Update(msg)
		m.AController = aController.(ApplicationController.Model)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
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

func initLogger(shs *sharedState.SharedState) {
	// Log files
	file, err := os.Create("log.txt")
	if err != nil {
		log.Println("Creating the Log file failed")
	}
	shs.Logger = log.New(file, "List", log.LstdFlags)
}

func Application() {
	shs := sharedState.GetGlobalState()
	initLogger(shs)

	mainModel := MainModel{
		sharedState: shs,

<<<<<<<< HEAD:src/tui/app.go
		MList:       MusicList.New(sharedState),
		MPlayer:     MusicPlayer.New(sharedState),
		AController: ApplicationController.New(),
========
		MList:   musicList.New(shs),
		MPlayer: musicPlayer.New(shs),
>>>>>>>> main:Workspace/tui/app.go
	}

	music.SetLogger(shs.Logger)
	music.Init()

	bzone.NewGlobal()
	p := tea.NewProgram(
		mainModel,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
