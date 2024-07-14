package clue

import (
	"fmt"
	"strings"
)

type Neighbor struct {
	Neighbor *Room
	Passage  bool
}

type Room struct {
	RoomName  string
	Marked    bool
	Note      []string
	Neighbors []*Neighbor
}

func (r *Room) IsRoom(name string) bool {
	if strings.Compare(name, r.RoomName) == 0 {
		return true
	}

	return false
}

func (c *Room) ClearNote() {
	c.Note = make([]string, 0)
}

func (c *Room) Mark(mark bool) {
	c.Marked = mark
}

func (c *Room) RemoveNote() {
	c.Note = append(c.Note[:len(c.Note)-2])
}

func (c *Room) AddNote(note string) {
	c.Note = append(c.Note, note)
}

func (r *Room) GetString() string {
	return r.RoomName
}

func (r *Room) PrintNeighbors() string {
	s := fmt.Sprintf("Current Room: %s\n", r.RoomName)
	for _, n := range r.Neighbors {
		s = fmt.Sprintf("%sNeighbor: %s\n", s, n.Neighbor.RoomName)
	}
	s = fmt.Sprintf("%s\n", s)

	return s
}

func (r *Room) GetValue() string {
	return r.RoomName
}

func (r *Room) PrintCard() string {

	return r.RoomName
}

func (r *Room) GetType() int {
	return ROOM
}

func (r *Room) AddNeighbor(room *Room, pass bool) {

	r.Neighbors = append(r.Neighbors, NewNeighbor(room, pass))
}

func NewNeighbor(room *Room, pass bool) *Neighbor {
	return &Neighbor{
		Neighbor: room,
		Passage:  pass,
	}
}

func NewRoom(name string) *Room {
	return &Room{
		RoomName:  name,
		Marked:    false,
		Neighbors: make([]*Neighbor, 0),
	}
}
