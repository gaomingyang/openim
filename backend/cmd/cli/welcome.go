package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var loginChoices = []string{"Log In", "Sign Up"}

type welcomeView struct {
	cursor int
	choice string
}

func (m welcomeView) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m welcomeView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = loginChoices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(loginChoices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(loginChoices) - 1
			}
		}
	}

	return m, nil
}

func (m welcomeView) View() string {
	s := strings.Builder{}
	s.WriteString("Login to the chat server or sign up?\n\n")

	for i := 0; i < len(loginChoices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(loginChoices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}
