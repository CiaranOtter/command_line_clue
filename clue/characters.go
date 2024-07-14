package clue

import (
	"fmt"
	"strings"
)

type Character struct {
	CharacterName string
	Colour        string
	Marked        bool
	Note          []string
	ClosestRoom   *Room
}

func (c *Character) IsCharacter(name string) bool {
	if strings.Compare(c.CharacterName, name) == 0 {
		return true
	}

	return false
}

func (c *Character) ClearNote() {
	c.Note = make([]string, 0)
}

func (c *Character) Mark(mark bool) {
	c.Marked = mark
}

func (c *Character) RemoveNote() {
	c.Note = append(c.Note[:len(c.Note)-2])
}

func (c *Character) AddNote(note string) {
	c.Note = append(c.Note, note)
}

func (c *Character) GetValue() string {
	return c.CharacterName
}

func (c *Character) GetType() int {
	return CHARACTER
}

func (c *Character) PrintCard() string {

	s := "+------------------+\n"
	s = fmt.Sprintf("%s| %s", s, c.GetString())

	for i := 2 + len(c.GetString()); i < 19; i++ {
		s = fmt.Sprintf("%s ", s)
	}

	s = fmt.Sprintf("%s|\n", s)

	for i := 0; i < 10; i++ {
		s = fmt.Sprintf("%s|                  |\n", s)
	}

	s = fmt.Sprintf("%s|", s)

	for i := len(c.GetString()) + 2; i < 19; i++ {
		s = fmt.Sprintf("%s ", s)
	}

	s = fmt.Sprintf("%s%s |\n", s, c.GetString())
	s = fmt.Sprintf("%s+------------------+\n", s)

	return s
}

func NewCharacter(name string, colour string, room string) *Character {
	return &Character{
		CharacterName: name,
		Colour:        colour,
		Marked:        false,
		ClosestRoom:   FindRoom(room),
	}
}
