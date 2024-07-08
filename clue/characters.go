package clue

import (
	"fmt"
)

type Character struct {
	CharacterName string
	Colour        string
	Marked        bool
	Note          []string
	ClosestRoom   *Room
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

func (c *Character) PrintCard() {
	fmt.Printf("Character:\n")
	fmt.Printf("\tName: %s\n\tColour: %s\n", c.GetString(), c.Colour)
	fmt.Printf("\n")
}

func NewCharacter(name string, colour string, room string) *Character {
	return &Character{
		CharacterName: name,
		Colour:        colour,
		Marked:        false,
		ClosestRoom:   FindRoom(room),
	}
}
