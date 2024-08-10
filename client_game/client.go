package main

import (
	"clue_client/login"
	pickchar "clue_client/pickchar"
	"command_line_clue/characters"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

var P *tea.Program

type Game struct {
	header        string
	active_screen tea.Model
	login         login.Login
	charchioce    pickchar.PickChar
}

func (g Game) Init() tea.Cmd {
	return g.active_screen.Init()
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return g, tea.Quit
		}
	case login.Login:
		g.header = msg.Login.Value()
		g.active_screen = g.charchioce
	}

	g.active_screen, cmd = g.active_screen.Update(msg)
	return g, cmd
}

func (g Game) View() string {
	return fmt.Sprintf("%s\n%s", g.header, g.active_screen.View())
}

func initialModel() tea.Model {
	login := login.NewLogin()
	game := Game{
		active_screen: login,
		login:         login,
		charchioce:    pickchar.NewChoice(characters.LoadCharacters("/Users/ciaranotter/Documents/personal/command_line_clue/command_line_clue/data/characters.csv")),
	}

	return game
}

func main() {
	P = tea.NewProgram(initialModel())
	if _, err := P.Run(); err != nil {
		log.Fatal(err)
	}
}
