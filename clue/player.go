package clue

import (
	"fmt"
	"strconv"
	"strings"
)

type Player struct {
	Char        *Character
	PlayerName  string
	Sheet       *CSheet
	Cards       []Card
	CurrentRoom *Room
}

type CSheet struct {
	Chars   []*Character
	Rooms   []*Room
	Weapons []*Weapon
}

func (p *Player) MarkCard(card Card) {
	switch card.GetType() {
	case CHARACTER:
		p.Sheet.MarkChar(card)
		break
	case ROOM:
		p.Sheet.MarkRoom(card)
		break
	case WEAPON:
		p.Sheet.MarkWeapon(card)
	}

}

func (p *Player) GetString() string {
	return fmt.Sprintf("%s%s%s", Colours[p.Char.Colour], p.PlayerName, Colours["Reset"])
}

func (p *Player) FinalAccusation() bool {
	return Solve.CheckAnswer(p.PickCharacter(true), p.PickWeapon(), p.PickRoom())

}

func (p *Player) CheckHas(char *Character, weap *Weapon, room *Room) (bool, Card) {
	showCards := make([]Card, 0)

	for _, card := range p.Cards {
		switch card.GetType() {
		case CHARACTER:
			if strings.Compare(char.CharacterName, card.GetValue()) == 0 {
				showCards = append(showCards, card)
			}
			break
		case ROOM:
			if strings.Compare(room.RoomName, card.GetValue()) == 0 {
				showCards = append(showCards, card)
			}
			break
		case WEAPON:
			if strings.Compare(weap.Name, card.GetValue()) == 0 {
				showCards = append(showCards, card)
			}
			break
		}
	}

	if len(showCards) == 0 {
		return false, nil
	}

	return true, p.AskShow(showCards)
}

func (p *Player) AskShow(cards []Card) Card {
	fmt.Printf("Which of these cards would you like you show?\n")

	for i, c := range cards {
		fmt.Printf("(%d) %s\n", i+1, c.GetString())
	}

	fmt.Printf("Your choice: ")
	var choice string

	fmt.Scan(&choice)

	i, err := strconv.Atoi(choice)

	if (err != nil) || (i < 1 || i > len(cards)) {
		fmt.Printf("%s is not a valid choice.\n")
		return p.AskShow(cards)
	}

	return cards[i-1]
}

func (p *Player) PickRoom() *Room {
	fmt.Printf("Pick a room:\n")
	for i, room := range p.Sheet.Rooms {
		fmt.Printf("(%d) %s\n", i+1, room.GetString())
	}

	var pChoice string
	fmt.Scan(&pChoice)

	i, err := strconv.Atoi(pChoice)

	if (err != nil) || (i < 1 || i > len(p.Sheet.Rooms)) {
		fmt.Printf("%s is not a valid choice.\n")
		return p.PickRoom()
	}

	return p.Sheet.Rooms[i-1]
}

func (p *Player) PickCharacter(makred bool) *Character {
	fmt.Printf("Pick a character:\n")
	for i, char := range p.Sheet.Chars {
		fmt.Printf("(%d) %s", i+1, char.GetString())
		if makred && char.Marked {
			fmt.Printf(": X")
		}
		fmt.Printf("\n")
	}

	var pchoice string
	fmt.Printf("Your choice: ")
	fmt.Scan(&pchoice)

	i, err := strconv.Atoi(pchoice)

	if (err != nil) || (i < 1 || i > len(p.Sheet.Chars)) {
		fmt.Printf("%s is not a valid choice.")
		return p.PickCharacter(makred)

	}

	return p.Sheet.Chars[i-1]
}

func (p *Player) PickWeapon() *Weapon {
	fmt.Printf("Pick a murder weapon:\n")

	for i, weap := range p.Sheet.Weapons {
		fmt.Printf("(%d) %s\n", i+1, weap.GetString())
	}

	var pChoice string
	fmt.Printf("Your choice: ")
	fmt.Scan(&pChoice)

	i, err := strconv.Atoi(pChoice)

	if (err != nil) || (i < 1 || i > len(p.Sheet.Weapons)) {
		fmt.Printf("%s is not a valid choice.\n", pChoice)
		return p.PickWeapon()
	}

	return p.Sheet.Weapons[i-1]
}

func (p *Player) MakeAccusation() (*Character, *Weapon, *Room) {
	char := p.PickCharacter(true)
	weap := p.PickWeapon()
	room := p.CurrentRoom

	return char, weap, room
}

func (p *Player) MakeMove() {
	fmt.Printf("You can move to:\n")
	total := 0
	if p.CurrentRoom == nil {
		total++
		fmt.Printf("(%d) %s\n", total, p.Char.ClosestRoom.RoomName)
	} else {
		for _, r := range p.CurrentRoom.Neighbors {
			total++
			fmt.Printf("(%d) %s", total, r.Neighbor.RoomName)

			if r.Passage {
				fmt.Printf(" (secret passage)")
			}

			fmt.Printf("\n")
		}

		total++
		fmt.Printf("Or (%d) stay where you are", total)
	}

	var choice string
	fmt.Scan(&choice)

	i, err := strconv.Atoi(choice)

	if (err != nil) || (i < 1 || i > total) {
		fmt.Printf("%s is not a valid choice.\n")
		defer p.MakeMove()
		return
	}

	if p.CurrentRoom == nil {
		p.CurrentRoom = p.Char.ClosestRoom
		fmt.Printf("Moving to %s.\n", p.CurrentRoom.RoomName)
		return
	} else {
		if i == total {
			fmt.Printf("Staying where you are.\n")
			return
		}

		p.CurrentRoom = p.CurrentRoom.Neighbors[i].Neighbor
		fmt.Printf("Moving to %s. \n", p.CurrentRoom.RoomName)
		return
	}

}

func (p *Player) AfterTurn() {
	fmt.Printf("Would you like to do anything before ending your turn?\n")

	fmt.Printf("(1) Show clue sheet\n")
	fmt.Printf("(2) Look at Cards\n")
	fmt.Printf("(3) Make Accusation\n")
	fmt.Printf("(4) End Turn\n")

	fmt.Printf("Your choice:\n")

	var choice string
	fmt.Scan(&choice)

	i, err := strconv.Atoi(choice)

	if (err != nil) || (i < 1 || i > 4) {
		fmt.Printf("%s is an invalid choice.\n")
		defer p.AfterTurn()
		return
	}

	switch i {
	case 1:
		p.ClueSheet()
		defer p.AfterTurn()
		break
	case 2:
		p.PrintCards()
		defer p.AfterTurn()
		break
	case 3:
		p.FinalAccusation()
		break
	case 4:
		break
	}

	return
}

func (p *Player) BeforeTurn() {
	fmt.Printf("%s's Turn:\n", p.GetString())
	fmt.Printf("\n")

	if p.CurrentRoom == nil {
		fmt.Printf("%s is the closests room.\n", p.Char.ClosestRoom.RoomName)
	} else {
		p.CurrentRoom.PrintNeighbors()
	}
	fmt.Printf("\n")
	fmt.Printf("What would you like to do?\n")

	fmt.Printf("(1) Show Clue sheet.\n")
	fmt.Printf("(2) Edit Clue sheet.\n")
	fmt.Printf("(3) Show Cards.\n")
	fmt.Printf("(4) Roll dice.\n")

	var c string
	fmt.Scan(&c)

	i, err := strconv.Atoi(c)

	if (err != nil) || (i < 1 || i > 3) {
		fmt.Printf("%s is not a valid choice.\n")
		defer p.BeforeTurn()
		return
	}

	switch i {
	case 1:
		p.ClueSheet()
		defer p.BeforeTurn()
		return
	case 2:
		p.EditClueSheet()
		defer p.BeforeTurn()
		return
	case 3:
		p.PrintCards()
		defer p.BeforeTurn()
		return
	case 4:
		p.Roll()
		return
	}
}

func (p *Player) Roll() {
	fmt.Printf("Player rolls the dice\n")
}

func (p *Player) EditCard(card Card) {
	fmt.Printf("(1) Add a note.\n")
	fmt.Printf("(2) Remove last note.\n")
	fmt.Printf("(3) Clear Notes. \n")
	fmt.Printf("(4) Mark X.\n")
	fmt.Printf("(5) Unmark X.\n")
	fmt.Printf("(6) Finish.\n")

	fmt.Printf("Your choice: ")
	var choice string

	fmt.Scan(&choice)

	i, err := strconv.Atoi(choice)

	if (err != nil) || (i < 1 || i > 6) {
		fmt.Printf("%s is an invalid choice.\n")
		defer p.EditCard(card)
		return
	}

	switch i {
	case 1:
		var Note string
		fmt.Printf("Note: ")
		fmt.Scan(&Note)

		card.AddNote(Note)
		defer p.EditCard(card)
		return
	case 2:
		card.RemoveNote()
		defer p.EditCard(card)
		return
	case 3:
		card.ClearNote()
		defer p.EditCard(card)
		return
	case 4:
		card.Mark(true)
		defer p.EditCard(card)
		return
	case 5:
		card.Mark(false)
		defer p.EditCard(card)
		return
	case 6:
		return
	}
}
func (p *Player) EditClueSheet() {
	fmt.Printf("Edit:\n")

	total := 0

	for _, char := range p.Sheet.Chars {
		total++
		fmt.Printf("(%d) %s\n", total, char.GetString())
	}

	for _, weap := range p.Sheet.Weapons {
		total++
		fmt.Printf("(%d) %s\n", total, weap.GetString())
	}

	for _, room := range p.Sheet.Rooms {
		total++
		fmt.Printf("(%d) %s\n", total, room.GetString())
	}

	fmt.Printf("(%d) Finish.\n", total+1)

	fmt.Printf("Your choice: ")
	var choice string

	fmt.Scan(&choice)

	i, err := strconv.Atoi(choice)

	if (err != nil) || (i < 1 || i > total+1) {
		fmt.Printf("%s is not a valid choice.\n", choice)
		defer p.EditClueSheet()
		return
	}

	if i <= len(p.Sheet.Chars) {
		p.EditCard(p.Sheet.Chars[i-1])
		defer p.EditClueSheet()
		return
	}

	i = i - len(p.Sheet.Chars)

	if i <= len(p.Sheet.Weapons) {
		p.EditCard(p.Sheet.Weapons[i-1])
		defer p.EditClueSheet()
		return
	}

	i = i - len(p.Sheet.Weapons)

	if i <= len(p.Sheet.Rooms) {
		p.EditCard(p.Sheet.Rooms[i-1])
		defer p.EditClueSheet()
		return
	}

	return
}

func (p *Player) ClueSheet() {
	fmt.Printf("Characters:\n")
	for _, char := range p.Sheet.Chars {
		fmt.Printf("\t%s: ", char.GetString())
		if char.Marked {
			fmt.Printf("X")
		} else {
			fmt.Printf(" ")
		}

		for _, note := range char.Note {
			fmt.Printf("%s,", note)
		}

		fmt.Printf("\n")
	}

	fmt.Printf("Weapons:\n")
	for _, weap := range p.Sheet.Weapons {
		fmt.Printf("\t%s: ", weap.GetString())

		if weap.Marked {
			fmt.Printf("X")
		} else {
			fmt.Printf(" ")
		}

		for _, note := range weap.Note {
			fmt.Printf("%s,", note)
		}

		fmt.Printf("\n")
	}

	fmt.Printf("Rooms:\n")
	for _, r := range p.Sheet.Rooms {
		fmt.Printf("\t%s: ", r.GetString())

		if r.Marked {
			fmt.Printf("X")
		} else {
			fmt.Printf(" ")
		}

		for _, note := range r.Note {
			fmt.Printf("%s,", note)
		}

		fmt.Printf("\n")
	}
}

func (p *Player) Print() {
	fmt.Print("Player ==========\n")
	fmt.Printf("Name:\t\t%s\n", p.GetString())
	fmt.Printf("Character:\t%s\n", p.Char.GetString())
	fmt.Printf("================\n")
}

func NewSheet(weap []*Weapon, room []*Room, char []*Character) *CSheet {
	return &CSheet{
		Chars:   char,
		Rooms:   room,
		Weapons: weap,
	}
}

func (s *CSheet) MarkRoom(w Card) {
	for _, room := range s.Rooms {
		if strings.Compare(w.GetValue(), room.GetValue()) == 0 {
			room.Marked = true
			return
		}
	}
}

func (s *CSheet) MarkWeapon(w Card) {
	for _, weap := range s.Weapons {
		if strings.Compare(weap.GetValue(), w.GetValue()) == 0 {
			weap.Marked = true
			return
		}
	}
}

func (s *CSheet) MarkChar(c Card) {
	for _, char := range s.Chars {
		if strings.Compare(c.GetValue(), char.GetValue()) == 0 {
			char.Marked = true
			return
		}
	}
}

func (p *Player) AutoMark() {
	for _, card := range p.Cards {
		switch card.GetType() {
		case ROOM:
			p.Sheet.MarkRoom(card)
			break
		case CHARACTER:
			p.Sheet.MarkChar(card)
			break
		case WEAPON:
			p.Sheet.MarkWeapon(card)
			break
		}
	}
}

func (p *Player) PrintCards() {

	fmt.Printf("%s's Cards\n", p.GetString())
	for _, card := range p.Cards {
		card.PrintCard()
	}
}

func NewPlayer(name string, char *Character, sheet *CSheet) *Player {

	return &Player{
		Char:        char,
		PlayerName:  name,
		Sheet:       sheet,
		Cards:       make([]Card, 0),
		CurrentRoom: nil,
	}
}
