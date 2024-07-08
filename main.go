package main

import (
	"command_line_clue/clue"
	"fmt"
	"strconv"
	"strings"
)

var Players []*clue.Player

func main() {
	clue.ReadFile("data/characters.csv", "data/rooms.csv", "data/weapons.csv")

	// clue.PrintWeapons()

	playerCount := 4

	Players = make([]*clue.Player, 0)
	for i := 0; i < playerCount; i++ {
		Players = append(Players, PickChar())
	}

	// clue.PrintWeapons()
	for _, p := range Players {
		p.Print()
	}

	clue.ConnectRooms()
	// clue.PrintWeapons()

	clue.FindClues() // decide on an answer

	fmt.Printf("Murder has happened\n")

	clue.DistCards(Players)

	for _, p := range Players {
		p.AutoMark()
	}

	for i, p := range Players {
		p.BeforeTurn()
		p.MakeMove()
		char, weap, room := p.MakeAccusation()

		fmt.Printf("%s is accusing %s with the %s in the %s\n", p.GetString(), char.GetString(), weap.Name, room.RoomName)

		temp_index := i + 1

		for j := 0; j < len(Players)-1; j++ {
			has, card := Players[temp_index].CheckHas(char, weap, room)

			if has {
				fmt.Printf("%s has shown the card: %s\n", Players[temp_index].GetString(), card.GetString())
				p.MarkCard(card)
				break
			}

			temp_index = (temp_index + 1) % playerCount
		}

		p.AfterTurn()
	}
	// clue.PrintWeapons()

	// clue.PrintCards()

	// clue.PrintWeapons()

	// player.BeforeTurn()
	// clue.PrintWeapons()

}

func CheckAvailable(index int) (bool, *clue.Player) {
	for _, pl := range Players {
		if (pl.Char != nil) && (strings.Compare(pl.Char.CharacterName, clue.Characters[index].CharacterName) == 0) {
			return false, pl
		}
	}

	return true, nil
}

func PickChar() *clue.Player {

	fmt.Printf("Whats your name: ")
	var name string

	fmt.Scan(&name)

	fmt.Printf("Pick a character:\n")
	total := 0
	for i, char := range clue.Characters {
		total++
		fmt.Printf("(%d) %s", total, char.GetString())

		if av, p := CheckAvailable(i); !av {
			fmt.Printf(" -> %s", p.GetString())
		}

		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	var choice string
	var charChoice *clue.Character

	for {
		fmt.Scan(&choice)

		i, err := strconv.Atoi(choice)

		if (err != nil) || (i <= 0 && i > total) {
			fmt.Printf("%s is not a valid choice\n")
			continue
		}

		av, p := CheckAvailable(i - 1)
		if !av {
			fmt.Printf("%s has already been taken by %s\n", clue.Characters[i-1].GetString(), p.GetString())
			continue
		}

		charChoice = clue.Characters[i-1]
		break
	}

	s := clue.NewSheet(clue.DuplicateWeap(), clue.DuplicateRooms(), clue.DuplicateChars())

	return clue.NewPlayer(name, charChoice, s)
}
