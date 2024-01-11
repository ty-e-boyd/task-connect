package main

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// APPLICATION STATE MODEL
type model struct {
	list list.Model
}

// MAIN ENTRY
func main() {
	log.Print("Starting Application..")

	items := []list.Item{
		item{title: "Install", desc: "Install Software"},
		item{title: "Update", desc: "Update Software"},
		item{title: "Remove", desc: "Remove Software"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Task Connect Installer"

	p := tea.NewProgram(m, tea.WithAltScreen())

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
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

// VIEW
func (m model) View() string {
	return docStyle.Render(m.list.View())
}
