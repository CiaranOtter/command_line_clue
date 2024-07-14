package game_server

import (
	"command_line_clue/clue"
	"command_line_clue/clue/comm"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type Client struct {
	Stream   *comm.ClueService_GameStreamServer
	Name     string
	Player   *clue.Player
	Ready    bool
	ShowChan chan clue.Card
}

func NewClient(name string, stream *comm.ClueService_GameStreamServer) *Client {
	return &Client{
		Name:     name,
		Stream:   stream,
		Ready:    false,
		ShowChan: make(chan clue.Card),
	}
}

func (c *Client) SendCards() {
	s := make([]string, 0)

	for _, c := range c.Player.Cards {
		s = append(s, c.GetValue())
	}

	(*c.Stream).Send(&comm.Message{
		Data: &comm.Message_Cards{
			Cards: &comm.Cards{
				Index: s,
			},
		},
	})

}

type GameService struct {
	*comm.UnimplementedClueServiceServer
	Clients        map[string]*Client
	Characters     map[string]*clue.Character
	Weapons        map[string]*clue.Weapon
	Rooms          map[string]*clue.Room
	Number_players int
	CurrentPlayer  int
	PlayOrderIndex []string
	Solution       *clue.Answer
	Cards          map[string]clue.Card
	Running        bool
}

// func (g *GameService) SetChar()

func (g *GameService) AddPlayer(name string, stream *comm.ClueService_GameStreamServer) bool {
	_, e := g.Clients[name]

	// if the username alredy exists
	if e || strings.Compare(name, "") == 0 {
		g.Warn(fmt.Sprintf("The username %s has already been taken", name))

		(*stream).Send(&comm.Message{
			Data: &comm.Message_Con{
				Con: &comm.Connect{
					PlayerName: name,
					Success:    false,
				},
			},
		})
		return false
	}

	// create a new client
	g.Clients[name] = NewClient(name, stream)
	// log the join
	g.Info(fmt.Sprintf("%s has joined the server", name))

	// send the success message to the client
	(*stream).Send(&comm.Message{
		Data: &comm.Message_Con{
			Con: &comm.Connect{
				PlayerName: name,
				Success:    true,
			},
		},
	})

	// Inform the other players of the new connection
	g.SendToOpp(&comm.Message{
		Data: &comm.Message_Opp{
			Opp: &comm.Opponent{
				Data: &comm.Opponent_Con{
					Con: &comm.Connect{
						PlayerName: name,
					},
				},
			},
		},
	}, g.Clients[name])

	//
	for _, cl := range g.Clients {
		if cl == g.Clients[name] {
			continue
		}

		// Send this client the already existing players in the game
		g.Info(fmt.Sprintf("Updating %s of the the already connected player %s", name, cl.Name))
		(*g.Clients[name].Stream).Send(&comm.Message{
			Data: &comm.Message_Opp{
				Opp: &comm.Opponent{
					PlayerName: cl.Name,
					Data: &comm.Opponent_Con{
						Con: &comm.Connect{
							PlayerName: cl.Name,
						},
					},
				},
			},
		})

		// if the other player has a character selected send the choice to this client
		if cl.Player.Char != nil {
			g.Info(fmt.Sprintf("Updating %s of %s's character choice %s", name, cl.Name, cl.Player.Char.GetString()))
			(*g.Clients[name].Stream).Send(&comm.Message{
				Data: &comm.Message_Opp{
					Opp: &comm.Opponent{
						PlayerName: cl.Name,
						Data: &comm.Opponent_SetChar{
							SetChar: &comm.SetPlayer{
								CharacterName: cl.Player.Char.CharacterName,
							},
						},
					},
				},
			})
		}
	}
	return true
}

func (g *GameService) SendToOpp(mes *comm.Message, sender *Client) {
	for _, cl := range g.Clients {
		if sender == cl {
			continue
		}

		(*cl.Stream).Send(mes)
	}
}

func (g *GameService) SetChar(name string, charName string) {
	_, e := g.Clients[name]

	if !e {
		g.Error(fmt.Sprintf("An unamed player attempted to pick a character"))
		return
	}

	char, e := g.Characters[charName]

	if !e {
		g.Error(fmt.Sprintf("%s attempting to pick an invalid character", name))
		return
	}

	g.Clients[name].Player = clue.NewPlayer(name, char, clue.NewSheet(clue.DuplicateWeap(), clue.DuplicateRooms(), clue.DuplicateChars()))
	g.Clients[name].Ready = true
	g.Info(fmt.Sprintf("%s has chosen th chracter %s", g.Clients[name].Player.GetString(), g.Clients[name].Player.Char.GetString()))

	g.SendToOpp(&comm.Message{
		Data: &comm.Message_Opp{
			Opp: &comm.Opponent{
				PlayerName: name,
				Data: &comm.Opponent_SetChar{
					SetChar: &comm.SetPlayer{
						CharacterName: charName,
					},
				},
			},
		},
	}, g.Clients[name])

	// for _, cl := range g.Clients {
	// 	if cl == g.Clients[name] {
	// 		continue
	// 	}

	// 	if g.Clients[name].Player.Char != nil {
	// 		(*g.Clients[name].Stream).Send(&comm.Message{
	// 			Data: &comm.Message_Opp{
	// 				Opp: &comm.Opponent{
	// 					PlayerName: cl.Name,
	// 					Data: &comm.Opponent_SetChar{
	// 						SetChar: &comm.SetPlayer{
	// 							CharacterName: g.Clients[name].Player.Char.GetValue(),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		})
	// 	}
	// }

}

func NewService(uri string) *GameService {
	lis, err := net.Listen("tcp", uri)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()
	Service := &GameService{
		Clients:        make(map[string]*Client),
		Characters:     make(map[string]*clue.Character),
		Weapons:        make(map[string]*clue.Weapon),
		Rooms:          make(map[string]*clue.Room),
		Number_players: 3,
		Running:        false,
	}

	char, room, weap := clue.ReadFile("/Users/ciaranotter/Documents/personal/command_line_clue/data/characters.csv", "/Users/ciaranotter/Documents/personal/command_line_clue/data/rooms.csv", "/Users/ciaranotter/Documents/personal/command_line_clue/data/weapons.csv")

	for _, c := range char {
		Service.Characters[c.CharacterName] = c
	}

	for _, w := range weap {
		Service.Weapons[w.Name] = w
	}

	for _, r := range room {
		Service.Rooms[r.RoomName] = r
	}

	comm.RegisterClueServiceServer(grpcServer, Service)

	Service.Info(fmt.Sprintf("Starting server at %s", uri))

	grpcServer.Serve(lis)

	return Service
}

func (s *GameService) Info(mes string) {
	fmt.Printf("[\033[1mINFO\033[0m] %s\n", mes)
}

func (s *GameService) Warn(mes string) {
	fmt.Printf("[\033[33m\033[1mWarn\033[0m] %s\n", mes)
}

func (s *GameService) Error(mes string) {
	fmt.Printf("[\033[31m\033[1mWarn\033[0m] %s\n", mes)
}

func (s *GameService) FindSolution() {
	s.Solution = clue.FindClues()
	s.Info("Solution has been set.")
}

func (s *GameService) ShuffleCards() {
	s.Cards = make(map[string]clue.Card)
	for _, c := range s.Characters {
		if s.Solution.Murderer.IsCharacter(c.CharacterName) {
			continue
		}

		s.Cards[c.GetValue()] = c
	}

	for _, w := range s.Weapons {
		if s.Solution.MurderWeapon.IsWeapon(w.Name) {
			continue
		}

		s.Cards[w.GetValue()] = w
	}

	for _, r := range s.Rooms {
		if s.Solution.Room.IsRoom(r.RoomName) {
			continue
		}

		s.Cards[r.GetValue()] = r
	}

	s.Info("Cards have been shuffled.")
}

func (s *GameService) ShufflePlayOrder() {
	s.PlayOrderIndex = make([]string, len(s.Clients))

	iter := 0
	for key, _ := range s.Clients {
		s.PlayOrderIndex[iter] = key
		iter++
	}

	for i := range s.PlayOrderIndex {
		j := rand.Intn(i + 1)
		s.PlayOrderIndex[i], s.PlayOrderIndex[j] = s.PlayOrderIndex[j], s.PlayOrderIndex[i]
	}

	s.CurrentPlayer = 0

	s.Info("Player order has been decided.")
}

func (s *GameService) DealCards() {
	pl_index := 0

	for _, c := range s.Cards {
		s.Clients[s.PlayOrderIndex[pl_index]].Player.Cards = append(s.Clients[s.PlayOrderIndex[pl_index]].Player.Cards, c)
		pl_index = (pl_index + 1) % len(s.Clients)
	}

	for _, cl := range s.Clients {
		cl.SendCards()
	}

	s.Info("Cards have been dealt.")
}

func (s *GameService) StartGame() {
	for _, cl := range s.Clients {
		if !cl.Ready {
			return
		}
	}

	s.FindSolution()
	s.ShuffleCards()
	s.ShufflePlayOrder()
	s.DealCards()

	for _, cl := range s.Clients {
		(*cl.Stream).Send(&comm.Message{
			Data: &comm.Message_Start{
				Start: &comm.StartGame{},
			},
		})
	}

	s.Info("Starting game.")

	s.Running = true
	(*s.Clients[s.PlayOrderIndex[s.CurrentPlayer]].Stream).Send(&comm.Message{
		Data: &comm.Message_Turn{
			Turn: &comm.StartTurn{},
		},
	})

	s.Info(fmt.Sprintf("Starting %s's turn", s.Clients[s.PlayOrderIndex[s.CurrentPlayer]].Player.GetString()))

	s.SendToOpp(&comm.Message{
		Data: &comm.Message_Opp{
			Opp: &comm.Opponent{
				PlayerName: s.Clients[s.PlayOrderIndex[s.CurrentPlayer]].Name,
				Data:       &comm.Opponent_Startturn{},
			},
		},
	}, s.Clients[s.PlayOrderIndex[s.CurrentPlayer]])

}

func (s *GameService) MovePlayer(name string, roomName string) {
	s.Clients[name].Player.CurrentRoom = s.Rooms[roomName]

	s.SendToOpp(&comm.Message{
		Data: &comm.Message_Opp{
			Opp: &comm.Opponent{
				PlayerName: name,
				Data: &comm.Opponent_Move{
					Move: &comm.Move{
						RoomName: roomName,
					},
				},
			},
		},
	}, s.Clients[name])

	s.Info(fmt.Sprintf("%s has moved to the %s", s.Clients[name].Player.GetString(), s.Clients[name].Player.CurrentRoom.RoomName))
}

func (s *GameService) CheckAnswer(name string, accuse *comm.Accuse) bool {
	if strings.Compare(s.Solution.MurderWeapon.GetValue(), accuse.GetWeap()) != 0 {
		return false
	}

	if strings.Compare(s.Solution.Murderer.GetValue(), accuse.GetChar()) != 0 {
		return false
	}

	if strings.Compare(s.Solution.Room.GetValue(), accuse.GetRoom()) != 0 {
		return false
	}

	return true
}

func (s *GameService) HandleAccuse(name string, accuse *comm.Accuse) {
	temp_iter := (s.CurrentPlayer + 1) % len(s.Clients)

	c := make([]string, 0)
	c = append(c, accuse.GetChar())
	c = append(c, accuse.GetRoom())
	c = append(c, accuse.GetWeap())

	s.Info(fmt.Sprintf("%s has accused %s in the %s with the %s", s.Clients[name].Player.GetString(), s.Characters[accuse.GetChar()].GetString(), s.Rooms[accuse.GetRoom()].GetString(), s.Weapons[accuse.GetWeap()].GetString()))

	for temp_iter != s.CurrentPlayer {
		// ask the player to show a card

		s.Info(fmt.Sprintf("Asking %s if they have a card to show.", s.Clients[s.PlayOrderIndex[temp_iter]].Player.GetString()))

		(*s.Clients[s.PlayOrderIndex[temp_iter]].Stream).Send(&comm.Message{
			Data: &comm.Message_Opp{
				Opp: &comm.Opponent{
					PlayerName: name,
					Data: &comm.Opponent_Ask{
						Ask: &comm.AskShow{
							Cards: c,
						},
					},
				},
			},
		})

		// rep, err := (*s.Clients[s.PlayOrderIndex[temp_iter]].Stream).Recv()

		rep := <-s.Clients[s.PlayOrderIndex[temp_iter]].ShowChan

		if rep != nil {
			s.Info(fmt.Sprintf("%s has a card to show", s.Clients[s.PlayOrderIndex[temp_iter]].Player.GetString()))

			(*s.Clients[name].Stream).Send(&comm.Message{
				Data: &comm.Message_Opp{
					Opp: &comm.Opponent{
						PlayerName: s.Clients[s.PlayOrderIndex[temp_iter]].Name,
						Data: &comm.Opponent_Show{
							Show: &comm.Show{
								HasCard: true,
								Card:    rep.GetValue(),
							},
						},
					},
				},
			})
			return
		}
		temp_iter = (temp_iter + 1) % len(s.Clients)
	}

	// if no one has send and empty message
	(*s.Clients[name].Stream).Send(&comm.Message{
		Data: &comm.Message_Show{
			Show: &comm.Show{
				Card:       "",
				PlayerName: "",
				HasCard:    false,
			},
		},
	})

	s.Info("No player has any cards to show")
}

func (s *GameService) GameStream(stream comm.ClueService_GameStreamServer) error {
	s.Info("New connection started\n")
	var Name string
	for {

		if !s.Running && len(s.Clients) == s.Number_players {
			s.StartGame()
		}
		mes, err := stream.Recv()

		if err == io.EOF {
			s.Info("Player has left the server")
			break
		}

		if err != nil {
			s.Error("Error in the server")
			return err
		}

		s.Info("Message received")

		switch mes.GetData().(type) {
		// if a connect request comes through
		case *comm.Message_Con:
			Name = mes.GetCon().GetPlayerName()
			s.AddPlayer(mes.GetCon().GetPlayerName(), &stream)
		case *comm.Message_SetChar:
			s.SetChar(Name, mes.GetSetChar().GetCharacterName())
		case *comm.Message_Move:
			s.MovePlayer(Name, mes.GetMove().GetRoomName())
		case *comm.Message_Accuse:

			if mes.GetAccuse().GetFinal() {
				if s.CheckAnswer(Name, mes.GetAccuse()) {
					s.Info(fmt.Sprintf("%s player go answer right", s.Clients[Name].Player.GetString()))
				} else {
					s.Info(fmt.Sprintf("%s player go answer Wrong", s.Clients[Name].Player.GetString()))
				}
			} else {
				s.HandleAccuse(Name, mes.GetAccuse())

			}
		case *comm.Message_Show:
			if mes.GetShow().GetHasCard() {
				s.Clients[Name].ShowChan <- s.Cards[mes.GetShow().GetCard()]
			} else {
				s.Clients[Name].ShowChan <- nil
			}
		case *comm.Message_Endturn:
			s.Info(fmt.Sprintf("%s has ended their turn", s.Clients[Name].Player.GetString()))
			s.CurrentPlayer = (s.CurrentPlayer + 1) % len(s.Clients)

			(*s.Clients[s.PlayOrderIndex[s.CurrentPlayer]].Stream).Send(&comm.Message{
				Data: &comm.Message_Turn{
					Turn: &comm.StartTurn{},
				},
			})

		}
	}

	return nil
}
