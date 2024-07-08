package clue

import "fmt"

type Weapon struct {
	Name   string
	Marked bool
	Note   []string
}

func (c *Weapon) ClearNote() {
	c.Note = make([]string, 0)
}

func (c *Weapon) Mark(mark bool) {
	c.Marked = mark
}

func (c *Weapon) RemoveNote() {
	c.Note = append(c.Note[:len(c.Note)-2])
}

func (c *Weapon) AddNote(note string) {
	c.Note = append(c.Note, note)
}
func (w *Weapon) GetString() string {
	return w.Name
}

func (w *Weapon) GetValue() string {
	return w.Name
}

func (w *Weapon) PrintCard() {
	fmt.Printf("Weapon: %s\n", w.Name)
}

func (w *Weapon) GetType() int {
	return WEAPON
}

func NewWeapon(name string) *Weapon {
	return &Weapon{
		Name:   name,
		Marked: false,
	}
}
