package main

import (
	"fmt"
	"os"

	"github.com/Erodotos/bubbletea-examples/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	p := tea.NewProgram(models.NewRootModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
