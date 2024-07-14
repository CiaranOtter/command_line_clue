package screen

import "fmt"

type PrintInterface interface {
	Print(b PrintInterface) string
	Reload(b PrintInterface)
	GetPadding() map[string]int
	GetDim() (int, int)
	GetData() []PrintInterface
	PrintPadding(s string) string
	UpChoice()
	DownChoice()
	MakeChoice()
	AddChar(letter int32)
	RemoveBlock(b PrintInterface)
}

type UnimplementedPrintInterface struct {
}

func (u *UnimplementedPrintInterface) RemoveBlock(b PrintInterface) {

}

func (u *UnimplementedPrintInterface) AddChar(letter int32) {
	// fmt.Printf("Adding a Character")
	for _, i := range u.GetData() {
		fmt.Printf("Adding a Character")

		i.AddChar(letter)
	}
}

func (u *UnimplementedPrintInterface) MakeChoice() {
	fmt.Printf("Unimplemented choice\n")
	for _, c := range u.GetData() {
		c.MakeChoice()
	}
}
func (u *UnimplementedPrintInterface) UpChoice() {
	for _, c := range u.GetData() {
		c.UpChoice()
	}
}

func (u *UnimplementedPrintInterface) DownChoice() {
	for _, c := range u.GetData() {
		c.DownChoice()
	}
}

func (u *UnimplementedPrintInterface) PrintPadding(s string) string {
	return ""
}

func (u *UnimplementedPrintInterface) Print(p PrintInterface) string {
	for _, i := range u.GetData() {
		i.Print(u)
	}
	return ""
}

func (u *UnimplementedPrintInterface) GetDim() (int, int) {
	return 0, 0
}

func (u *UnimplementedPrintInterface) GetData() []PrintInterface {
	return nil
}

func (u *UnimplementedPrintInterface) GetPadding() map[string]int {
	return map[string]int{
		"Left":   0,
		"Right":  0,
		"Top":    0,
		"Bottom": 0,
	}
}

func (u *UnimplementedPrintInterface) Reload(p PrintInterface) {
	for _, d := range u.GetData() {
		d.Reload(u)
	}
}
