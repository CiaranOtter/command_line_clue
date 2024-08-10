package characters

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Character struct {
	Name   string
	Colour string
}

type CharacterList struct {
	Characters []Character
}

func LoadCharacters(filename string) CharacterList {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	charlist := CharacterList{
		Characters: []Character{},
	}

	for {
		player, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		charlist.Characters = append(charlist.Characters, Character{
			Name:   player[0],
			Colour: player[1],
		})
	}

	return charlist
}
