package msg

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	Login int = 0
	Pick      = 1
)

type (
	FrameMsg struct {
		Frame int
	}
)

func Frame(i int) tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return FrameMsg{
			Frame: i,
		}
	})
}
