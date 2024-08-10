package login

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Login struct {
	Login textinput.Model
}

func NewLogin() Login {
	l1 := textinput.New()
	l1.Placeholder = "Username"
	l1.Focus()
	l1.CharLimit = 50
	l1.Width = 50

	return Login{
		Login: l1,
	}
}

func (l Login) Init() tea.Cmd {
	return textinput.Blink
}

func (l Login) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return l, tea.Quit
		case tea.KeyEnter:
			return l, tea.Tick(time.Second/60, func(time.Time) tea.Msg {
				return l
			})
		}
	}

	l.Login, cmd = l.Login.Update(msg)
	return l, cmd
}

func (l Login) View() string {
	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s",
		l.Login.View(),
		"(esc to quit)",
	) + "\n"
}
