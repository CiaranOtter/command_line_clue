package main

import (
	"command_line_clue/clue"
	"fmt"
)

func main() {

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
