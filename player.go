package main

import (
	"command_line_clue/clue"
	"command_line_clue/clue/comm"
	"command_line_clue/screen"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type Game struct {
	Stream         *comm.ClueService_GameStreamClient
	Player         *clue.Player
	Players        map[string]*clue.Player
	Screen         *screen.Screen
	GameHeading    *screen.Block
	GameBody       *screen.Block
	Waiting_Screen *screen.Block
	Running        bool
	Characters     map[string]*clue.Character
	Weapons        map[string]*clue.Weapon
	Rooms          map[string]*clue.Room
	Cards          map[string]clue.Card
}

var G *Game

func (g *Game) AddOpponent(opp *comm.Connect) {
	g.Players[opp.GetPlayerName()] = &clue.Player{
		PlayerName: opp.GetPlayerName(),
		Waiting_block: screen.NewBlock(5, 5, []screen.PrintInterface{
			screen.NewHeading(fmt.Sprintf("%s has joined the game", opp.GetPlayerName()), " "),
		}),
	}

	g.Waiting_Screen.AddBlock(g.Players[opp.GetPlayerName()].Waiting_block)
}

func (g *Game) SetChar(name string, opp *comm.SetPlayer) {
	g.Players[name] = clue.NewPlayer(name, g.Characters[opp.GetCharacterName()], clue.NewSheet(clue.DuplicateWeap(), clue.DuplicateRooms(), clue.DuplicateChars()))
	g.Waiting_Screen.RemoveBlock(g.Players[name].Waiting_block)
	g.Players[name].Waiting_block = screen.NewBlock(5, 5, []screen.PrintInterface{
		screen.NewHeading(fmt.Sprintf("%s has joined as %s", name, g.Players[name].Char.GetString()), " "),
	})
	g.Waiting_Screen.AddBlock(g.Players[name].Waiting_block)

}

func (g *Game) AskShow(name string, ask *comm.AskShow) {
	c := make(map[int]string)
	real := make(map[int]string)

	iter := 0
	for _, card := range g.Player.Cards {

		for _, a := range ask.GetCards() {
			//
			if strings.Compare(card.GetValue(), a) == 0 {
				iter++
				c[iter] = card.GetString()
				real[iter] = a
				break
			}
		}
	}

	if len(c) == 0 {
		(*g.Stream).Send(&comm.Message{
			Data: &comm.Message_Show{
				Show: &comm.Show{
					Card:    "",
					HasCard: false,
				},
			},
		})
		return
	}

	g.GameBody.Clear()
	cl, in := screen.NewChoiceList(fmt.Sprintf("Which of the following cards would you like to show %s", g.Players[name].GetString()), c)
	g.GameBody.AddBlock(cl)

	choice := <-in

	g.GameBody.Clear()

	(*g.Stream).Send(&comm.Message{
		Data: &comm.Message_Show{
			Show: &comm.Show{
				Card:       real[choice],
				PlayerName: g.Player.PlayerName,
				HasCard:    true,
			},
		},
	})
}

func (g *Game) OppTurn(name string) {
	g.GameBody.Clear()

	g.GameBody.AddBlock(screen.NewHeading(fmt.Sprintf("%s's turn", g.Players[name].GetString()), "+"))
}

func (g *Game) MovePlayer(name string, roomName string) {
	g.GameBody.AddBlock(screen.NewHeading(fmt.Sprintf("%s has moved to the %s", g.Players[name].GetString(), g.Rooms[roomName].GetString()), " "))
	g.Players[name].CurrentRoom = g.Rooms[roomName]
}

func (g *Game) HandleOpponent(opp *comm.Opponent) {
	switch opp.GetData().(type) {
	case *comm.Opponent_Con:
		g.AddOpponent(opp.GetCon())
	case *comm.Opponent_SetChar:
		g.SetChar(opp.GetPlayerName(), opp.GetSetChar())
	case *comm.Opponent_Ask:
		g.AskShow(opp.GetPlayerName(), opp.GetAsk())
	case *comm.Opponent_Move:
		g.MovePlayer(opp.GetPlayerName(), opp.GetMove().GetRoomName())
	case *comm.Opponent_Startturn:
		g.OppTurn(opp.GetPlayerName())
	}
}

func (s *Game) GetCards(cards []string) {
	s.Cards = make(map[string]clue.Card)
	for _, c := range s.Characters {
		s.Cards[c.GetValue()] = c
	}

	for _, w := range s.Weapons {
		s.Cards[w.GetValue()] = w
	}

	for _, r := range s.Rooms {
		s.Cards[r.GetValue()] = r
	}

	for _, card := range cards {
		s.Player.Cards = append(s.Player.Cards, s.Cards[card])
	}

	s.Player.AutoMark()
}

func (g *Game) StartGame() {
	g.GameBody.Clear()
	g.GameBody.AddBlock(screen.NewHeading("Starting game", "="))
}

func (g *Game) ShowClueSheet(cl *screen.Block) {

	cl.Clear()
	char := make([]string, 0)
	for _, c := range g.Player.Sheet.Chars {
		s := c.GetString()

		if c.Marked {
			s = fmt.Sprintf("\tX\t%s", s)
		} else {
			s = fmt.Sprintf("\t \t%s", s)
		}

		char = append(char, s)
	}

	weap := make([]string, 0)

	for _, w := range g.Player.Sheet.Weapons {
		s := w.GetString()

		if w.Marked {
			s = fmt.Sprintf("\tX\t%s", s)
		} else {
			s = fmt.Sprintf("\t \t%s", s)
		}

		weap = append(weap, s)
	}

	room := make([]string, 0)

	for _, r := range g.Player.Sheet.Rooms {
		s := r.GetString()

		if r.Marked {
			s = fmt.Sprintf("\tX\t%s", s)
		} else {
			s = fmt.Sprintf("\t \t%s", s)
		}

		room = append(room, s)
	}

	cl.AddBlock(screen.NewHeading("Characters:", ""))
	cl.AddBlock(screen.NewList(char))
	cl.AddBlock(screen.NewHeading("Weapons:", ""))
	cl.AddBlock(screen.NewList(weap))
	cl.AddBlock(screen.NewHeading("Rooms:", ""))
	cl.AddBlock(screen.NewList(room))
}

func (g *Game) ShowCards(cl *screen.Block) {
	cl.Clear()

	c := make([]string, 0)

	for _, card := range g.Player.Cards {
		c = append(c, card.GetString())
	}

	cl.AddBlock(screen.NewList(c))
}

func (g *Game) MakeMove(s *screen.Block) {

	c := make(map[int]string)
	realRoom := make(map[int]string)

	if g.Player.CurrentRoom == nil {
		c[1] = fmt.Sprintf("Move to the %s", g.Player.Char.ClosestRoom.RoomName)
		realRoom[1] = g.Player.Char.ClosestRoom.RoomName
	} else {
		c[1] = fmt.Sprintf("Stay in the %s", g.Player.CurrentRoom.RoomName)
		realRoom[1] = g.Player.CurrentRoom.RoomName
		iter := 1

		for _, r := range g.Player.CurrentRoom.Neighbors {
			iter++
			c[iter] = fmt.Sprintf("Move to the %s", r.Neighbor.RoomName)
			realRoom[iter] = r.Neighbor.RoomName

			if r.Passage {
				c[iter] = fmt.Sprintf("%s (secret passage)", c[iter])
			}
		}

	}

	cl, in := screen.NewChoiceList("", c)
	s.AddBlock(cl)

	choice := <-in

	g.Player.CurrentRoom = g.Rooms[realRoom[choice]]

	s.RemoveBlock(cl)
	s.AddBlock(screen.NewHeading(fmt.Sprintf("You have moved to the %s", g.Player.CurrentRoom.RoomName), " "))

	(*g.Stream).Send(&comm.Message{
		Data: &comm.Message_Move{
			Move: &comm.Move{
				RoomName: g.Player.CurrentRoom.RoomName,
			},
		},
	})
}

func (g *Game) PickWeap(s *screen.Block) *clue.Weapon {
	c := make(map[int]string)
	realChar := make(map[int]string)

	iter := 0
	for _, char := range g.Weapons {
		iter++
		c[iter] = char.GetString()
		realChar[iter] = char.GetValue()
	}

	cl, in := screen.NewChoiceList("Who do you think is the murderer?", c)

	s.AddBlock(cl)

	choice := <-in

	s.RemoveBlock(cl)
	return g.Weapons[realChar[choice]]
}

func (g *Game) AccChar(s *screen.Block) *clue.Character {
	c := make(map[int]string)
	realChar := make(map[int]string)

	iter := 0
	for _, char := range g.Characters {
		iter++
		c[iter] = char.GetString()
		realChar[iter] = char.GetValue()
	}

	cl, in := screen.NewChoiceList("Who do you think is the murderer?", c)

	s.AddBlock(cl)

	choice := <-in

	s.RemoveBlock(cl)
	return g.Characters[realChar[choice]]
}

func (g *Game) AccRoom(s *screen.Block) *clue.Room {
	c := make(map[int]string)
	realChar := make(map[int]string)

	iter := 0
	for _, char := range g.Rooms {
		iter++
		c[iter] = char.GetString()
		realChar[iter] = char.GetValue()
	}

	cl, in := screen.NewChoiceList("Who do you think is the murderer?", c)

	s.AddBlock(cl)

	choice := <-in

	s.RemoveBlock(cl)
	return g.Rooms[realChar[choice]]
}

func (g *Game) MakeAccusation(s *screen.Block, final bool) {
	char := g.AccChar(s)
	weap := g.PickWeap(s)

	var room *clue.Room
	if final {
		room = g.AccRoom(s)
	} else {
		room = g.Player.CurrentRoom
	}

	s.AddBlock(screen.NewHeading(fmt.Sprintf("You have accused %s in the %s with the %s", char.GetString(), weap.GetString(), g.Player.CurrentRoom.RoomName), " "))

	(*g.Stream).Send(&comm.Message{
		Data: &comm.Message_Accuse{
			Accuse: &comm.Accuse{
				Char:  char.GetValue(),
				Weap:  weap.GetValue(),
				Room:  room.GetValue(),
				Final: final,
			},
		},
	})

	mes, err := (*g.Stream).Recv()

	if err != nil {
		log.Fatal()
	}

	if !mes.GetOpp().GetShow().GetHasCard() {
		s.AddBlock(screen.NewHeading("No other player has these cards", " "))
		return
	}

	player := mes.GetOpp().GetPlayerName()
	card := mes.GetOpp().GetShow().GetCard()

	s.AddBlock(screen.NewHeading(fmt.Sprintf("%s has show you the card %s", g.Players[player].GetString(), g.Cards[card].GetString()), " "))
	g.Player.MarkCard(g.Cards[card])
}

func (g *Game) MyTurn() {
	g.GameBody.Clear()
	g.GameBody.AddBlock(screen.NewHeading("Your turn", "+"))

	if g.Player.CurrentRoom == nil {
		g.GameBody.AddBlock(screen.NewHeading(fmt.Sprintf("The %s is the closest room to you", g.Player.Char.ClosestRoom.RoomName), " "))
	} else {
		g.GameBody.AddBlock(screen.NewHeading(fmt.Sprintf("You are currently in the %s", g.Player.CurrentRoom.RoomName), " "))
	}

	c := make(map[int]string)

	c[1] = "Show clue sheet"
	c[2] = "Show Cards"
	c[3] = "Roll Dice"

	ch, choice := screen.NewChoiceList("What would you like to do?", c)

	gameBlock := screen.NewBlock(5, 5, []screen.PrintInterface{})
	g.GameBody.AddBlock(gameBlock)
	g.GameBody.AddBlock(ch)

	for {
		i := <-choice

		if i == 1 {
			g.ShowClueSheet(gameBlock)
		}

		if i == 2 {
			g.ShowCards(gameBlock)
		}

		if i == 3 {
			g.GameBody.RemoveBlock(ch)
			g.GameBody.RemoveBlock(gameBlock)
			break
		}
	}

	gameBlock.Clear()
	g.GameBody.AddBlock(gameBlock)
	g.MakeMove(gameBlock)

	g.MakeAccusation(gameBlock, false)

	c[3] = "Make Final accusation"
	c[4] = "End Turn"

	ch, choice = screen.NewChoiceList("What would you like to do?", c)
	g.GameBody.AddBlock(ch)

	for {
		i := <-choice

		if i == 1 {
			g.ShowClueSheet(gameBlock)
		}

		if i == 2 {
			g.ShowCards(gameBlock)
		}

		if i == 3 {
			g.MakeAccusation(gameBlock, true)
		}

		if i == 4 {
			(*g.Stream).Send(&comm.Message{
				Data: &comm.Message_Endturn{
					Endturn: &comm.EndTurn{},
				},
			})
			break
		}
	}

	gameBlock.Clear()
	g.GameBody.Clear()

}
func (g *Game) HandleMessage() {
	g.Running = true

	waiting := screen.NewBlock(5, 5, []screen.PrintInterface{
		screen.NewHeading("Waiting for game to start", "-"),
		g.Waiting_Screen,
	})

	g.GameBody.Clear()
	g.GameBody.AddBlock(waiting)

	for g.Running {
		mes, err := (*g.Stream).Recv()

		if err != nil {
			(*g.Stream).CloseSend()
			g.Running = false
		}

		switch mes.GetData().(type) {
		case *comm.Message_Opp:
			g.HandleOpponent(mes.GetOpp())
		case *comm.Message_Cards:
			g.GetCards(mes.GetCards().GetIndex())
		case *comm.Message_Start:
			g.StartGame()
		case *comm.Message_Turn:
			g.MyTurn()
		}
	}
}

func (g *Game) Login() {
	input, v := screen.NewTextInput("What's your name: ")

	g.GameBody.AddBlock(input)
	g.Screen.Input = true
	name := <-v
	g.Screen.Input = false
	g.GameBody.Clear()

	(*g.Stream).Send(&comm.Message{
		Data: &comm.Message_Con{
			Con: &comm.Connect{
				PlayerName: name,
			},
		},
	})

	resp, err := (*g.Stream).Recv()

	if err != nil {
		log.Fatal(err)
	}

	if !resp.GetCon().GetSuccess() {
		defer g.Login()
		return
	}

	g.Players[resp.GetCon().GetPlayerName()] = &clue.Player{
		PlayerName: resp.GetCon().GetPlayerName(),
	}

	g.Player = g.Players[resp.GetCon().GetPlayerName()]

	g.GameHeading.PutFirst(screen.NewHeading(g.Player.PlayerName, ""))

	g.PickChar()
}

func (g *Game) PickChar() {

	g.GameBody.Clear()

	c := make(map[int]string)
	real_name := make(map[int]string)

	iter := 0
	for _, char := range g.Characters {
		iter++
		c[iter] = char.GetString()

		for _, p := range g.Players {
			if p.Char != nil && p.Char.IsCharacter(char.CharacterName) {
				c[iter] = fmt.Sprintf("%s <- %s", c[iter], p.PlayerName)
			}
		}
		real_name[iter] = char.CharacterName
	}

	list, resp := screen.NewChoiceList("Pick your character", c)
	g.GameBody.AddBlock(list)

	ch := <-resp

	for _, p := range g.Players {
		if p.Char != nil && p.Char.IsCharacter(real_name[ch]) {
			defer g.PickChar()
			return
		}
	}

	g.GameBody.Clear()

	g.Player = clue.NewPlayer(g.Player.PlayerName, g.Characters[real_name[ch]], clue.NewSheet(clue.DuplicateWeap(), clue.DuplicateRooms(), clue.DuplicateChars()))

	(*g.Stream).Send(&comm.Message{
		Data: &comm.Message_SetChar{
			SetChar: &comm.SetPlayer{
				CharacterName: real_name[ch],
			},
		},
	})

	g.GameHeading.PutFirst(screen.NewHeading(g.Player.Char.GetString(), ""))
	g.GameBody.Clear()
	g.GameBody.AddBlock(g.Waiting_Screen)
}

func NewService(uri string) {

}

func refresh() {
	for G.Screen.Running {
		time.Sleep(time.Millisecond * 500)
		G.Screen.Refresh()
	}
}

func main() {
	// connect to the referee
	conn, err := grpc.NewClient("localhost:5000", grpc.WithInsecure())

	// if there is an errr break
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create a New client
	client := comm.NewClueServiceClient(conn)

	// GEt the stream from the connection
	stream, err := client.GameStream(context.Background())

	// break if there is an error
	if err != nil {
		log.Fatal(err)
	}

	G = &Game{
		Stream: &stream,
		Screen: screen.NewScreen(),
		GameHeading: screen.NewBlock(5, 5, []screen.PrintInterface{
			screen.NewHeading("CLUE", "#"),
		}),
		GameBody:       screen.NewBlock(5, 5, []screen.PrintInterface{}),
		Waiting_Screen: screen.NewBlock(5, 5, []screen.PrintInterface{}),
		Characters:     make(map[string]*clue.Character),
		Weapons:        make(map[string]*clue.Weapon),
		Rooms:          make(map[string]*clue.Room),
		Players:        make(map[string]*clue.Player),
	}

	char, room, weap := clue.ReadFile("/Users/ciaranotter/Documents/personal/command_line_clue/data/characters.csv", "/Users/ciaranotter/Documents/personal/command_line_clue/data/rooms.csv", "/Users/ciaranotter/Documents/personal/command_line_clue/data/weapons.csv")

	for _, c := range char {
		G.Characters[c.CharacterName] = c
	}

	for _, w := range weap {
		G.Weapons[w.Name] = w
	}

	for _, r := range room {
		G.Rooms[r.RoomName] = r
	}

	G.Screen.AddBlock(G.GameHeading)
	G.Screen.AddBlock(G.GameBody)

	go G.Screen.Keys()
	go refresh()
	G.Login()

	G.HandleMessage()
}
