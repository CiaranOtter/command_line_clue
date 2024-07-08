package main

import (
	"command_line_clue/clue"
	"command_line_clue/clue/comm"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

var Players map[string]*clue.Player
var Player *clue.Player
var Name string
var Stream comm.ClueService_GameStreamClient

func HandleCommand(message *comm.Command) {
	switch message.GetType() {
	case comm.CommandType_PICK_CHAR:
		Player = PickChar(Name)

		Players[Name] = Player

		Stream.Send(&comm.Message{
			Data: &comm.Message_SetChar{
				SetChar: &comm.SetPlayer{
					CharacterName: Player.Char.CharacterName,
				},
			},
		})

		break
	}
}

func HandleOpponent(message *comm.Opponent) {

	t := message.GetData()

	Name := message.GetPlayerName()

	switch t.(type) {
	case *comm.Opponent_Con:
		fmt.Printf("Another player has connected.\n")
		Players[message.GetCon().GetPlayerName()] = &clue.Player{
			PlayerName: message.GetCon().GetPlayerName(),
		}
		break
	case *comm.Opponent_SetChar:
		fmt.Printf("Setting opponents character.\n")

		for _, char := range clue.Characters {
			if strings.Compare(char.CharacterName, message.GetSetChar().GetCharacterName()) == 0 {
				Players[Name] = clue.NewPlayer(Name, char, nil)
				return
			}
		}

	}

}

func GetCards(message *comm.Cards) {

	clue.Cards = make([]clue.Card, 0)

	for _, room := range clue.Rooms {
		clue.Cards = append(clue.Cards, room)
	}

	for _, char := range clue.Characters {
		clue.Cards = append(clue.Cards, char)
	}

	for _, weap := range clue.Weapons {
		clue.Cards = append(clue.Cards, weap)
	}

	MyCards := message.GetIndex()

	Player.Cards = make([]clue.Card, 0)

	for _, index := range MyCards {
		Player.Cards = append(Player.Cards, clue.Cards[index])
	}
}

func RemoveCards(message *comm.RemoveAnswer) {
	clue.Characters = append(clue.Characters[:message.GetChar()], clue.Characters[message.GetChar()+1:]...)
	clue.Rooms = append(clue.Rooms[:message.GetRoom()], clue.Rooms[message.GetRoom()+1:]...)
	clue.Weapons = append(clue.Weapons[:message.GetWeap()], clue.Weapons[message.GetWeap()+1:]...)
}

func main() {
	conn, err := grpc.NewClient("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := comm.NewClueServiceClient(conn)

	Stream, err = client.GameStream(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	clue.ReadFile("data/characters.csv", "data/rooms.csv", "data/weapons.csv")

	clue.ConnectRooms()
	Players = make(map[string]*clue.Player, 0)

	fmt.Printf("What's your name: ")
	fmt.Scan(&Name)

	Stream.Send(&comm.Message{
		Data: &comm.Message_Con{
			Con: &comm.Connect{
				PlayerName: Name,
			},
		},
	})

	running := true

	for running {
		command, err := Stream.Recv()

		if err != nil {
			log.Fatal(err)
			continue
		}

		data := command.GetData()

		switch data.(type) {
		case *comm.Message_End:
			running = false
			break
		case *comm.Message_Com:
			HandleCommand(command.GetCom())
			break
		case *comm.Message_Opp:
			HandleOpponent(command.GetOpp())
			break
		case *comm.Message_Start:
			fmt.Printf("Starting game\n")
			break
		case *comm.Message_Cards:
			GetCards(command.GetCards())
			break

		case *comm.Message_RAns:
			RemoveCards(command.GetRAns())
			break
		}
	}

	fmt.Printf("Game Over\n")

}

func PickChar(name string) *clue.Player {

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

func CheckAvailable(index int) (bool, *clue.Player) {
	for _, pl := range Players {
		if (pl.Char != nil) && (strings.Compare(pl.Char.CharacterName, clue.Characters[index].CharacterName) == 0) {
			return false, pl
		}
	}

	return true, nil
}
