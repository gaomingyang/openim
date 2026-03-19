package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	username := login()

	sendCh, recvCh := websocketConnection()

	p := tea.NewProgram(initialChatModel(username, sendCh, recvCh))
	_, err := p.Run()
	errorCheck(err)

}

func login() string {
	p := tea.NewProgram(welcomeView{})
	m, err := p.Run()
	errorCheck(err)

	var username, password string

	switch m.(welcomeView).choice {
	case "Log In":
		username, password = loginView()
	case "Sign Up":
		for {
			// tea.ClearScreen()

			m, err := tea.NewProgram((initialRegsiterModel())).Run()
			errorCheck(err)

			username = m.(registerModel).inputs[0].Value()
			password = m.(registerModel).inputs[1].Value()
			repeatPassword := m.(registerModel).inputs[2].Value()

			if password != repeatPassword {
				fmt.Println("Passwords do not match!")
				continue
			}
			regiserRequest(username, password)
		}
	default:
		fmt.Println("Goodbye!")
		os.Exit(0)
	}

	loginRequest(username, password)
	return username
}

func loginView() (string, string) {
	m, err := tea.NewProgram(initialLoginModel()).Run()
	errorCheck(err)
	return m.(registerModel).inputs[0].Value(), m.(registerModel).inputs[1].Value()
}
