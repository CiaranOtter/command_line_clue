package screen

import (
	"fmt"
	"sort"
)

type ChoiceList struct {
	UnimplementedPrintInterface
	Choices       map[int]string
	Message       string
	CurrentChoice int
	Output        chan int
	Keys          []int
}

func (cl *ChoiceList) AddChocie(i int, s string) {
	cl.Choices[i] = s
}

func (cl *ChoiceList) MakeChoice() {
	fmt.Printf("You have made the choice.")
	cl.Output <- cl.CurrentChoice
}

func (cl *ChoiceList) DownChoice() {
	if cl.CurrentChoice > cl.Keys[0] {
		cl.CurrentChoice--
	}
}

func (cl *ChoiceList) UpChoice() {
	if cl.CurrentChoice < cl.Keys[len(cl.Keys)-1] {
		cl.CurrentChoice += 1
	}
}

func (cl *ChoiceList) Print(b PrintInterface) string {
	s := fmt.Sprintf("%s%s\n", b.PrintPadding("Left"), cl.Message)

	for _, key := range cl.Keys {
		// add the item to the list
		s = fmt.Sprintf("%s%s", s, b.PrintPadding("Left"))

		// This is the print of the item in the list
		s = fmt.Sprintf("%s%s", s, cl.Choices[key])

		// if the index is currently on this item
		if key == cl.CurrentChoice {

			// if the is the index at the current choice
			s = fmt.Sprintf("%s%s", s, " <-")
		}

		s = fmt.Sprintf("%s\n", s)
	}

	return s
}

func NewChoiceList(mes string, choices map[int]string) (*ChoiceList, chan int) {

	var key int

	for key, _ = range choices {
		break
	}

	KeysList := make([]int, 0)
	for key, _ := range choices {
		KeysList = append(KeysList, key)
	}

	sort.Ints(KeysList)

	out := make(chan int)
	return &ChoiceList{
		Choices:       choices,
		Message:       mes,
		CurrentChoice: key,
		Output:        out,
		Keys:          KeysList,
	}, out
}
