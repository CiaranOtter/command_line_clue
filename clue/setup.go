package clue

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadFile(charFile string, roomFile string, weaponRoom string) ([]*Character, []*Room, []*Weapon) {

	file, err := os.Open(roomFile)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	Rooms = make([]*Room, 0)

	for _, row := range records {
		Rooms = append(Rooms, NewRoom(row[0]))
	}

	file, err = os.Open(weaponRoom)

	if err != nil {
		log.Fatal(err)
	}

	reader = csv.NewReader(file)
	records, _ = reader.ReadAll()

	Weapons = make([]*Weapon, 0)

	for _, row := range records {
		Weapons = append(Weapons, NewWeapon(row[0]))
	}

	file, err = os.Open(charFile)

	if err != nil {
		log.Fatal(err)
	}

	reader = csv.NewReader(file)
	records, _ = reader.ReadAll()

	Characters = make([]*Character, 0)

	for _, row := range records {
		Characters = append(Characters, NewCharacter(row[0], row[1], row[2]))
	}

	return Characters, Rooms, Weapons
}
