package login

import (
	"clc_services/profile"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"google.golang.org/grpc"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type ExitMsg struct{}
type Logout struct {
	Message *profile.Reply
}

type AccountInterface interface {
	GetQuestion() string
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	GetMessage() *profile.Account
}

type Register struct {
	AccountInterface
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func (r Register) GetQuestion() string {
	return "Register new username"
}

func (r Register) GetMessage() *profile.Account {
	return &profile.Account{
		Name:     r.inputs[0].Value(),
		Surname:  r.inputs[1].Value(),
		Username: r.inputs[2].Value(),
		Register: true,
	}
}

func NewRegister() Register {
	m := Register{
		inputs: make([]textinput.Model, 3),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Nickname"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Email"
			t.CharLimit = 64
		case 2:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.inputs[i] = t
	}

	return m
}

func (m Register) Init() tea.Cmd {
	return textinput.Blink
}

func (m Register) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Tick(time.Second/60, func(time.Time) tea.Msg {
				return ExitMsg{}
			})

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Tick(time.Second/60, func(time.Time) tea.Msg {
					return m
				})
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *Register) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Register) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

type Login struct {
	AccountInterface
	Login textinput.Model
}

func (l Login) GetQuestion() string {
	return "Login"
}

func NewLogin() AccountInterface {
	l1 := textinput.New()
	l1.Placeholder = "Username"
	l1.Focus()
	l1.CharLimit = 50
	l1.Width = 50

	return Login{
		Login: l1,
	}
}

func (l Login) GetMessage() *profile.Account {
	return &profile.Account{
		Username: l.Login.Value(),
		Register: false,
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
		case tea.KeyCtrlC:
			return l, tea.Tick(time.Second/60, func(time.Time) tea.Msg {
				return ExitMsg{}
			})
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
	) + "\n"
}

type LoginOrRegsiter struct {
	cursor      int
	choices     []AccountInterface
	chosen      bool
	accountServ profile.AccountServiceClient
	spinner     spinner.Model
	Username    string
}

func (l LoginOrRegsiter) Logout() (tea.Model, tea.Cmd) {
	resp, err := l.accountServ.Logout(context.Background(), &profile.Account{
		Username: l.Username,
	})

	if err != nil {
		panic(err)
	}

	return l, tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return Logout{
			Message: resp,
		}
	})

}

func NewLoginChoice(conn *grpc.ClientConn) tea.Model {
	account_conn := profile.NewAccountServiceClient(conn)

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return LoginOrRegsiter{
		cursor: 0,
		choices: []AccountInterface{
			NewLogin(),
			NewRegister(),
		},
		accountServ: account_conn,
		spinner:     s,
	}
}

func (c LoginOrRegsiter) Init() tea.Cmd {
	var cmds []tea.Cmd = []tea.Cmd{}
	for _, m := range c.choices {
		cmds = append(cmds, m.Init())
	}

	cmds = append(cmds, c.spinner.Tick)
	return tea.Batch(cmds...)
}

func (c LoginOrRegsiter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if c.chosen {

		switch msg := msg.(type) {
		case tea.KeyMsg:
			// if the user presses control q then go back to the choice screen
			switch msg.Type {
			case tea.KeyEsc:
				c.chosen = false
				return c, nil
			}

		case Login, Register:
			// The user has submitted a login
			req := c.choices[c.cursor].GetMessage()
			response, err := c.accountServ.Login(context.Background(), req)

			if err != nil {
				panic(err)
			}

			if response.GetSuccess() {
				c.Username = req.GetUsername()
				// var cmd tea.Cmd
				// var inter tea.Model
				// inter, cmd = c.Done()
				// c = inter.(LoginOrRegsiter)
				return c, tea.Tick(time.Second/60, func(time.Time) tea.Msg {
					return c
				})
			}
		}

		// else update the Appropriate screen
		var command tea.Cmd
		var inter tea.Model
		inter, command = c.choices[c.cursor].Update(msg)
		c.choices[c.cursor] = inter.(AccountInterface)
		return c, command
	}

	// if the choice screen is being displayed
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// choice was made
			c.chosen = true
			m, cmd := c.choices[c.cursor].Update(nil)
			c.choices[c.cursor] = m.(AccountInterface)
			return c, cmd

		case "down", "j":

			if len(c.choices) > c.cursor {
				c.cursor = (c.cursor + 1) % len(c.choices)
			}

		case "up", "k":
			if c.cursor > 0 {
				c.cursor = (c.cursor - 1) % len(c.choices)
			}
		}
	}

	return c, nil
}

func (c LoginOrRegsiter) View() string {
	s := strings.Builder{}

	if !c.chosen {
		for i, m := range c.choices {
			if c.cursor == i {
				s.WriteString("(•) ")
			} else {
				s.WriteString("( ) ")
			}

			s.WriteString(m.GetQuestion())
			s.WriteString("\n")
		}
	} else {
		s.WriteString(c.choices[c.cursor].View())
	}
	s.WriteString("(esc to quit)")
	return s.String()
}
