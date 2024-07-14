package main

import game_server "command_line_clue/game_server"

func main() {
	game_server.NewService("localhost:5000")
}
