package chat

// A simple program demonstrating the text area component from the Bubbles
// component library.

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

var (
	chat_style lipgloss.Style
)

type ChatWindow struct {
	viewport        viewport.Model
	messages        []string
	textarea        textarea.Model
	senderStyle     lipgloss.Style
	container_style *lipgloss.Style
	err             error
}

func NewChatWindow(cont_style *lipgloss.Style) ChatWindow {
	chat_style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#88d498"))
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 5)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return ChatWindow{
		textarea:        ta,
		messages:        []string{},
		viewport:        vp,
		senderStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		container_style: cont_style,
		err:             nil,
	}
}

func (c ChatWindow) Init() tea.Cmd {
	return textarea.Blink
}

func (c ChatWindow) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	c.textarea, tiCmd = c.textarea.Update(msg)
	c.viewport, vpCmd = c.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			c.messages = append(c.messages, c.senderStyle.Render("You: ")+c.textarea.Value())
			c.viewport.SetContent(strings.Join(c.messages, "\n"))
			c.textarea.Reset()
			c.viewport.GotoBottom()
		}

	// We handle errors just like any other message
	case errMsg:
		c.err = msg
		return c, nil
	}

	return c, tea.Batch(tiCmd, vpCmd)
}

func (c ChatWindow) View() string {
	chats := chat_style.Width(c.container_style.GetWidth()-8).Height((8*c.container_style.GetHeight())/10-4).Margin(0, 2).Padding(2)
	c.viewport.Height = chats.GetHeight()
	c.viewport.Width = chats.GetWidth()
	wind := chats.Render(c.viewport.View())

	tb := lipgloss.NewStyle().Width(c.container_style.GetWidth()-4).Height((2*c.container_style.GetHeight())/10-1).Margin(0, 2)
	c.textarea.SetWidth(lipgloss.Width(wind))
	c.textarea.SetHeight(tb.GetHeight())

	return lipgloss.JoinVertical(0, wind, tb.Render(c.textarea.View()))
}
