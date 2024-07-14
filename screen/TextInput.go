package screen

import "fmt"

type TextInput struct {
	UnimplementedPrintInterface
	Message string
	Value   string
	Output  chan string
}

func NewTextInput(Message string) (*TextInput, chan string) {
	c := make(chan string)
	return &TextInput{
		Message: Message,
		Output:  c,
	}, c
}

func (t *TextInput) MakeChoice() {
	t.Output <- t.Value
}

func (t *TextInput) Print(b PrintInterface) string {
	s := fmt.Sprintf("%s%s", t.Message, t.Value)
	return s
}

func (t *TextInput) AddChar(letter int32) {
	fmt.Printf("Adding a character to the value")
	t.Value = fmt.Sprintf("%s%s", t.Value, string(letter))
}
