package main

import (
	"clue_client/chat"
	"clue_client/login"
	pickchar "clue_client/pickchar"
	"command_line_clue/characters"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var P *tea.Program

var (
	header_style     lipgloss.Style
	body_style       lipgloss.Style
	side_panle_style *lipgloss.Style
	windowHeight     int
	windowWidth      int
)

type MainScreen struct {
	header     string
	height     int
	width      int
	container  string
	chatWindow tea.Model
}

func NewMainScreen() MainScreen {
	return MainScreen{
		chatWindow: chat.NewChatWindow(side_panle_style),
		container:  "test window",
	}
}

func (m MainScreen) Init() tea.Cmd {
	return nil
}

func (m MainScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd = []tea.Cmd{}

	m.chatWindow, cmd = m.chatWindow.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainScreen) View() string {
	header := lipgloss.NewStyle().Width((2 * windowWidth) / 3).Height(windowHeight / 10).Inherit(header_style).Render(m.header)
	main := lipgloss.NewStyle().Width(2 * windowWidth / 3).Height(windowHeight - lipgloss.Height(header)).Inherit(body_style).Render(m.container)
	*side_panle_style = lipgloss.NewStyle().Width(windowWidth / 3).Height(windowHeight)
	chat := side_panle_style.Render(m.chatWindow.View())
	return lipgloss.JoinHorizontal(0, lipgloss.JoinVertical(0, header, main), chat)
}

type Game struct {
	active_screen tea.Model
	login         login.Login
	charchioce    pickchar.PickChar
	mainscreen    MainScreen
	User_colour   string
}

func (g Game) Init() tea.Cmd {
	return g.active_screen.Init()
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = []tea.Cmd{}
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return g, tea.Quit
		}
	case login.Login:
		g.mainscreen.header = msg.Login.Value()
		g.active_screen = g.charchioce
	case pickchar.PickChar:
		g.mainscreen.header = fmt.Sprintf("%s: %s", g.mainscreen.header, msg.GetCharString())
		g.active_screen = g.mainscreen
		g.User_colour = msg.GetColour()
		(g.mainscreen.chatWindow).SetColour(g.User_colour)
	case tea.WindowSizeMsg:

		windowHeight = msg.Height
		windowWidth = msg.Width

		g.active_screen, cmd = g.active_screen.Update(msg)
		cmds = append(cmds, cmd)
		return g, tea.Batch(cmds...)
	}

	g.active_screen, cmd = g.active_screen.Update(msg)
	cmds = append(cmds, cmd)

	return g, tea.Batch(cmds...)
}

func (g Game) View() string {
	return g.active_screen.View()
}

func initialModel() tea.Model {

	char_file := os.Getenv("CHAR_FILE")
	login := login.NewLogin()
	header_style = lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), false, false, true)
	body_style = lipgloss.NewStyle()
	t := lipgloss.NewStyle()
	side_panle_style = &t

	game := Game{
		active_screen: login,
		login:         login,
		mainscreen:    NewMainScreen(),
		charchioce:    pickchar.NewChoice(characters.LoadCharacters(char_file)),
	}

	return game
}

func main() {

	P = tea.NewProgram(initialModel())
	if _, err := P.Run(); err != nil {
		log.Fatal(err)
	}
}
