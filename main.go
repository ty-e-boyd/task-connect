package main

import (
	"context"
	"fmt"
	"os"

	configs "github.com/ty-e-boyd/task-connect/configs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// APPLICATION STATE MODEL
type model struct {
	calService *calendar.Service
}

// MSG TYPES
type initialSetupMsg struct {
	result *calendar.Service
}

// CMDs
func doInitialSetup() tea.Msg {
	ctx := context.Background()
	b, err := os.ReadFile("creds.json")
	if err != nil {
		log.Fatal("unable to read client secret")
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatal("unable to create config from creds")
	}

	client := configs.GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatal("unable to create new calendar service")
	}

	result := initialSetupMsg{result: srv}

	return result
}

// MAIN ENTRY
func main() {
	m := model{}

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
		m.calService = msg.result
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

	s += "did not fail, which is sorta like a success."

	return "\n" + s + "\n\n"
}
