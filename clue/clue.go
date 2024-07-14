package clue

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	ROOM      int = 0
	WEAPON        = 1
	CHARACTER     = 2
)

var Characters []*Character
var Rooms []*Room
var Weapons []*Weapon
var Solve *Answer

var Cards []Card

type Card interface {
	PrintCard() string
	GetString() string
	GetType() int
	GetValue() string
	Mark(mark bool)
	RemoveNote()
	ClearNote()
	AddNote(note string)
}

type Answer struct {
	Room         *Room
	Murderer     *Character
	MurderWeapon *Weapon
}

func (c *Character) GetString() string {
	return fmt.Sprintf("%s%s%s", Colours[c.Colour], c.CharacterName, Colours["Reset"])
}

func (a *Answer) CheckAnswer(char *Character, weap *Weapon, room *Room) bool {
	return (strings.Compare(char.CharacterName, a.Murderer.CharacterName) == 0) && (strings.Compare(room.RoomName, a.Room.RoomName) == 0) && (strings.Compare(weap.Name, a.MurderWeapon.Name) == 0)
}

func FindRoom(room string) *Room {
	for _, i := range Rooms {
		if strings.Compare(room, i.RoomName) == 0 {
			return i
		}
	}

	return nil
}

func ConnectRooms() {
	room := Rooms[0]

	for i := 1; i < len(Rooms); i++ {
		room.AddNeighbor(Rooms[i], false)
		Rooms[i].AddNeighbor(room, false)

		room = Rooms[i]
	}

	Rooms[0].AddNeighbor(Rooms[len(Rooms)-1], false)
	Rooms[len(Rooms)-1].AddNeighbor(Rooms[0], false)

	Rooms[1].AddNeighbor(Rooms[5], true)
	Rooms[5].AddNeighbor(Rooms[1], true)

	Rooms[3].AddNeighbor(Rooms[8], true)
	Rooms[8].AddNeighbor(Rooms[3], true)
}

func (a *Answer) Print() {
	fmt.Printf("The Amswer =================\n")
	fmt.Printf("Murderer:\t%s\n", Solve.Murderer.CharacterName)
	fmt.Printf("Weapon:\t\t%s\n", Solve.MurderWeapon.Name)
	fmt.Printf("Room:\t\t%s\n", Solve.Room.RoomName)
	fmt.Printf("============================\n")
}

func DistCards(players []*Player) {

	pIndex := 0
	for _, card := range Cards {
		players[pIndex].Cards = append(players[pIndex].Cards, card)
		pIndex = (pIndex + 1) % len(players)
	}
}

func PrintWeapons() {
	for _, w := range Weapons {
		w.PrintCard()
	}
}

func PrintCards() {
	fmt.Printf("There are %d cards in the deck:\n", len(Cards))
	for _, card := range Cards {
		card.PrintCard()
	}
}

func DuplicateRooms() []*Room {
	DupRooms := make([]*Room, 0)
	for _, room := range Rooms {
		DupRooms = append(DupRooms, NewRoom(room.RoomName))
	}

	return DupRooms
}

func DuplicateChars() []*Character {
	DupChars := make([]*Character, 0)
	for _, char := range Characters {
		DupChars = append(DupChars, NewCharacter(char.CharacterName, char.Colour, char.ClosestRoom.RoomName))
	}

	return DupChars
}

func DuplicateWeap() []*Weapon {
	DupWeap := make([]*Weapon, 0)

	for _, weap := range Weapons {
		DupWeap = append(DupWeap, NewWeapon(weap.Name))
	}

	return DupWeap
}

func ShuffleCards() {
	for i := range Cards {
		j := rand.Intn(i + 1)
		Cards[i], Cards[j] = Cards[j], Cards[i]
	}
}

func FindClues() *Answer {
	i := rand.Intn(len(Characters))
	j := rand.Intn(len(Rooms))
	k := rand.Intn(len(Weapons))

	Solve = &Answer{
		Room:         Rooms[j],
		Murderer:     Characters[i],
		MurderWeapon: Weapons[k],
	}

	Rooms = append(Rooms[:j], Rooms[j+1:]...)
	Characters = append(Characters[:i], Characters[i+1:]...)
	Weapons = append(Weapons[:k], Weapons[k+1:]...)

	Cards = make([]Card, 0)
	for _, room := range Rooms {
		Cards = append(Cards, room)
	}
	for _, char := range Characters {
		Cards = append(Cards, char)
	}
	for _, weap := range Weapons {
		Cards = append(Cards, weap)
	}

	// ShuffleCards()
	return Solve
}
