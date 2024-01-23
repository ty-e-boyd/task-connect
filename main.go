package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

// APPLICATION STATE MODEL
type model struct {
	message string
}

// MSG TYPES
type initialSetupMsg struct {
	result string
}

// CMDs
func doInitialSetup() tea.Msg {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("cannot obtain user home directory")
	}

	app := homedir + "/bin/bash"
	arg1 := "-c"
	arg2 := "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

	cmd := exec.Command(app, arg1, arg2)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("unable to run brew install curl -- %v", err)
	}

	result := initialSetupMsg{result: "command ran, no fatal. exiting now.."}

	return result
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
	return doInitialSetup
}

// UPDATE
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// check for custom msgs
	case initialSetupMsg:
		m.message = msg.result
		return m, tea.Quit

	// check for key press
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
