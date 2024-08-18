package pickchar

import (
	"command_line_clue/characters"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

var Colours = map[string]string{
	"Yellow": "#FFC43D",
	"Red":    "#EF476F",
	"Purple": "#C98BF9",
	"Green":  "#54AF97",
	"White":  "#F8FFE5",
	"Blue":   "#1B65AB",
}

type PickChar struct {
	cursor  int
	choices CharChoices
	choice  CharChoice
}

func (p PickChar) GetCharString() string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(Colours[p.choices.choices[p.cursor].Char.Colour])).Render(p.choices.choices[p.cursor].Char.Name)
}

func (p PickChar) GetColour() string {
	return Colours[p.choice.Char.Colour]
}

func (p PickChar) Init() tea.Cmd {
	return nil
}

func (p PickChar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Send the choice on the channel and exit.
			p.choice = p.choices.choices[p.cursor]
			p.choices.choices[p.cursor].taken = true
			// return p, tea.Quit
			return p, tea.Tick(time.Second/60, func(time.Time) tea.Msg {
				return p
			})

		case "down", "j":
			p.cursor++
			if p.cursor >= len(p.choices.choices) {
				p.cursor = 0
			}
			for p.choices.choices[p.cursor].taken {
				if p.cursor >= len(p.choices.choices) {
					p.cursor = 0
				}
				p.cursor++
			}

		case "up", "k":
			p.cursor--
			if p.cursor < 0 {
				p.cursor = len(p.choices.choices) - 1
			}
			for p.choices.choices[p.cursor].taken {
				if p.cursor < 0 {
					p.cursor = len(p.choices.choices) - 1
				}
				p.cursor--
			}
		}
	}

	return p, nil
}

func (p PickChar) View() string {
	s := strings.Builder{}
	s.WriteString("What kind of Bubble Tea would you like to order?\n\n")

	for i := 0; i < len(p.choices.choices); i++ {
		if p.choices.choices[i].taken {
			s.WriteString("(X) ")
			style := lipgloss.NewStyle().Strikethrough(true).Foreground(lipgloss.Color(Colours[p.choices.choices[i].Char.Colour])).Render(p.choices.choices[i].Char.Name)
			s.WriteString(style)
			s.WriteString("\n")
			continue
		} else if p.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(Colours[p.choices.choices[i].Char.Colour])).Render(p.choices.choices[i].Char.Name)
		s.WriteString(style)
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func NewChoice(c characters.CharacterList) PickChar {
	return PickChar{
		cursor:  0,
		choices: LoadChoices(c),
	}
}
