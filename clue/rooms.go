package clue

import "fmt"

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

func (r *Room) PrintNeighbors() {
	fmt.Printf("Current Room: %s\n", r.RoomName)
	for _, n := range r.Neighbors {
		fmt.Printf("Neighbor: %s\n", n.Neighbor.RoomName)
	}
	fmt.Printf("\n")

}

func (r *Room) GetValue() string {
	return r.RoomName
}

func (r *Room) PrintCard() {
	fmt.Printf("Room: %s\n", r.RoomName)
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
