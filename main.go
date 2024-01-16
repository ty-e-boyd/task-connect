package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

// APPLICATION STATE MODEL
type model struct {
	message string
}

// MAIN ENTRY
func main() {
	m := model{message: "this is a message"}

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Errorf("Unable to run application: %v", err)
		os.Exit(1)
	}
}

// INIT
func (m model) Init() tea.Cmd {
	return nil
}

// UPDATE
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}

	return m, nil
}

// VIEW
func (m model) View() string {
	s := fmt.Sprintln("hey")

	if len(m.message) > 0 {
		s += m.message
	}

	return "\n" + s + "\n\n"
}
