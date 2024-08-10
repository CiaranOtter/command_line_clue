package pickchar

import (
	"command_line_clue/characters"
)

type CharChoices struct {
	choices []CharChoice
}

type CharChoice struct {
	Char  characters.Character
	taken bool
}

func LoadChoices(chars characters.CharacterList) CharChoices {
	choices := CharChoices{
		choices: []CharChoice{},
	}
	for _, char := range chars.Characters {
		choices.choices = append(choices.choices, CharChoice{
			Char:  char,
			taken: false,
		})
	}

	return choices
}
