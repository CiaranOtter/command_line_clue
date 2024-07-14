package clue

import (
	"command_line_clue/screen"
	"fmt"
	"strconv"
	"strings"
)

type Player struct {
	Char          *Character
	PlayerName    string
	Sheet         *CSheet
	Cards         []Card
	CurrentRoom   *Room
	Waiting_block *screen.Block
}

type CSheet struct {
	Chars   []*Character
	Rooms   []*Room
	Weapons []*Weapon
}

var Colours = map[string]string{
	"Yellow": "\033[33m",
	"Red":    "\033[31m",
	"Purple": "\033[35m",
	"Green":  "\033[32m",
	"White":  "\033[37m",
	"Blue":   "\033[34m",
	"Reset":  "\033[0m",
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

func (p *Player) GetCharacter(name string) *Character {
	for _, char := range p.Sheet.Chars {
		if char.IsCharacter(name) {
			return char
		}
	}

	return nil
}

func (p *Player) GetString() string {
	return fmt.Sprintf("%s%s%s", Colours[p.Char.Colour], p.PlayerName, Colours["Reset"])
}

func (p *Player) FinalAccusation(s *screen.Block) bool {
	return Solve.CheckAnswer(p.PickCharacter(true, s), p.PickWeapon(s), p.PickRoom())

}

func (p *Player) CheckHas(char *Character, weap *Weapon, room *Room) (bool, []Card) {
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

	return true, showCards
}

func (p *Player) GetCards(cardName []string) []Card {
	cards := make([]Card, 0)
	for _, card := range p.Cards {
		for i, name := range cardName {
			if strings.Compare(name, card.GetValue()) == 0 {
				cards = append(cards, card)
				cardName = append(cardName[:i], cardName[i+1:]...)
				break
			}
		}
	}

	return cards
}

func (p *Player) AskShow(cards []Card, s *screen.Block) Card {

	l := make(map[int]string)

	for i, c := range cards {
		l[i+1] = c.GetString()
	}

	cl, resp := screen.NewChoiceList("Which of these cards would you like you show?", l)

	s.AddBlock(cl)

	i := <-resp

	s.Clear()

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

func (p *Player) PickCharacter(makred bool, s *screen.Block) *Character {

	c := make(map[int]string)

	for i, char := range p.Sheet.Chars {
		c[i+1] = char.GetString()
		if makred && char.Marked {
			c[i+1] = fmt.Sprintf("%s: X", c[i+1])
		}
	}

	list, resp := screen.NewChoiceList("Pick a character:", c)

	s.AddBlock(list)

	i := <-resp

	s.RemoveBlock(list)

	return p.Sheet.Chars[i-1]
}

func (p *Player) PickWeapon(s *screen.Block) *Weapon {
	c := make(map[int]string)

	for i, weap := range p.Sheet.Weapons {
		c[i+1] = weap.GetString()
	}

	list, resp := screen.NewChoiceList("Pick a murder weapon:", c)

	s.AddBlock(list)

	i := <-resp

	s.RemoveBlock(list)

	return p.Sheet.Weapons[i-1]
}

func (p *Player) MakeAccusation(s *screen.Block) (*Character, *Weapon, *Room) {
	char := p.PickCharacter(true, s)
	weap := p.PickWeapon(s)
	room := p.CurrentRoom

	return char, weap, room
}

func (p *Player) ApplyMove(RoomName string) bool {
	if p.CurrentRoom == nil {

		if p.Char.ClosestRoom.IsRoom(RoomName) {
			p.CurrentRoom = p.Char.ClosestRoom
			return true
		}

		return false

	} else {

		if p.CurrentRoom.IsRoom(RoomName) {
			return true
		}

		for _, r := range p.CurrentRoom.Neighbors {
			if r.Neighbor.IsRoom(RoomName) {
				p.CurrentRoom = r.Neighbor
				return true
			}
		}

		return false
	}
}

func (p *Player) MakeMove(s *screen.Block) {
	total := 0

	c := make(map[int]string)

	if p.CurrentRoom == nil {
		total++
		c[total] = p.Char.ClosestRoom.RoomName
	} else {
		for _, r := range p.CurrentRoom.Neighbors {
			total++
			c[total] = r.Neighbor.RoomName

			if r.Passage {
				c[total] = fmt.Sprintf("%s (secret passage)", c[total])
			}
		}

		total++
		c[total] = fmt.Sprintf("Stay in the %s", p.CurrentRoom.GetString())

	}

	l, v := screen.NewChoiceList("You can move to:", c)
	s.AddBlock(l)

	i := <-v
	s.RemoveBlock(l)

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

func (p *Player) AfterTurn(s *screen.Block) {
	c := make(map[int]string)
	c[1] = "Show clue sheet"
	c[2] = "Look at Cards"
	c[3] = "Make Accusation"
	c[4] = "End Turn"

	list, resp := screen.NewChoiceList("Would you like to do anything before ending your turn?", c)

	s.AddBlock(list)

	chosen := false

	for !chosen {
		i := <-resp

		switch i {
		case 1:
			s.RemoveBlock(list)
			p.ClueSheet(s)
			s.AddBlock(list)
		case 2:
			s.RemoveBlock(list)
			p.PrintCards()
			s.AddBlock(list)
		case 3:
			s.RemoveBlock(list)
			p.FinalAccusation(s)
			s.AddBlock(list)
		case 4:
			chosen = true
		}
	}

	s.RemoveBlock(list)

	return
}

func (p *Player) BeforeTurn(s *screen.Block) {

	heading := screen.NewHeading("Your Turn", "=")
	var text *screen.Heading
	if p.CurrentRoom == nil {
		text = screen.NewHeading(fmt.Sprintf("%s is the closests room.\n", p.Char.ClosestRoom.RoomName), " ")
	} else {
		text = screen.NewHeading(p.CurrentRoom.PrintNeighbors(), " ")
	}

	c := make(map[int]string)
	c[1] = "Show Clue sheet."
	c[2] = "Edit Clue sheet."
	c[3] = "Show Cards."
	c[4] = "Roll dice."
	list, resp := screen.NewChoiceList("What would you like to do?\n", c)

	tb := screen.NewBlock(5, 5, []screen.PrintInterface{heading, text, list})
	s.AddBlock(tb)

	chosen := false
	for !chosen {
		chocie := <-resp

		switch chocie {
		case 1:
			s.RemoveBlock(tb)
			p.ClueSheet(s)
			s.AddBlock(tb)
		case 2:
			s.RemoveBlock(tb)
			p.EditClueSheet()
			s.AddBlock(tb)
		case 3:
			p.PrintCards()
		case 4:
			p.Roll()
			chosen = true
		}
	}

	s.RemoveBlock(tb)

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

func (p *Player) ClueSheet(s *screen.Block) {

	heading := "CLUE SHEET"
	sHeading := "Characters:"

	chars := make([]string, 0)

	for _, char := range p.Sheet.Chars {
		s := fmt.Sprintf("\t")
		if char.Marked {
			s = fmt.Sprintf("%sX ", s)
		} else {
			s = fmt.Sprintf("%s  ", s)
		}

		s = fmt.Sprintf("%s %s", s, char.GetString())

		for _, note := range char.Note {
			s = fmt.Sprintf("%s%s,", s, note)
		}

		chars = append(chars, s)
	}

	cList := screen.NewBlock(5, 5, []screen.PrintInterface{
		screen.NewHeading(sHeading, "-"),
		screen.NewList(chars),
	})

	wHeading := "Weapons:"
	weaps := make([]string, 0)

	for _, weap := range p.Sheet.Weapons {

		s := fmt.Sprintf("\t")
		if weap.Marked {
			s = fmt.Sprintf("%sX ", s)
		} else {
			s = fmt.Sprintf("%s  ", s)
		}

		s = fmt.Sprintf("%s%s: ", s, weap.GetString())

		for _, note := range weap.Note {
			s = fmt.Sprintf("%s%s,", s, note)
		}

		weaps = append(weaps, s)
	}

	wList := screen.NewBlock(5, 5, []screen.PrintInterface{
		screen.NewHeading(wHeading, "-"),
		screen.NewList(weaps),
	})

	rHeading := "Rooms:"
	rooms := make([]string, 0)
	for _, r := range p.Sheet.Rooms {

		s := fmt.Sprintf("\t")

		if r.Marked {
			s = fmt.Sprintf("%sX ", s)
		} else {
			s = fmt.Sprintf("%s  ", s)
		}

		s = fmt.Sprintf("%s%s: ", s, r.GetString())

		for _, note := range r.Note {
			s = fmt.Sprintf("%s%s,", s, note)
		}

		rooms = append(rooms, s)
	}

	rList := screen.NewBlock(5, 5, []screen.PrintInterface{
		screen.NewHeading(rHeading, "-"),
		screen.NewList(rooms),
	})

	exit, e := screen.NewChoiceList("", map[int]string{
		1: "done",
	})
	block := screen.NewBlock(5, 5, []screen.PrintInterface{
		screen.NewHeading(heading, "="),
		cList,
		wList,
		rList,
		exit,
	})

	s.AddBlock(block)

	<-e

	s.RemoveBlock(block)
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
