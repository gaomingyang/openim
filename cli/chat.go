package main

import (
	"fmt"
	"strings"

	"openim/ws"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

type chatModel struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	otherStyle  lipgloss.Style
	err         error
	username    string
	sendCh      chan<- ws.Message
	recvCh      <-chan ws.Message
}

func initialChatModel(username string, sendCh chan<- ws.Message, recvCh <-chan ws.Message) chatModel {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 20)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return chatModel{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("12")),
		otherStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("236")),
		err:         nil,
		username:    username,
		sendCh:      sendCh,
		recvCh:      recvCh,
	}
}

func (m chatModel) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		tea.ClearScreen,
		recvMsgCmd(m.recvCh),
	)
}

func recvMsgCmd(recvCh <-chan ws.Message) tea.Cmd {
	return func() tea.Msg {
		msg := <-recvCh
		return msg
	}
}

func (m chatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)
	cmds := []tea.Cmd{tiCmd, vpCmd}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			msg := m.textarea.Value()
			// m.messages = append(m.messages, m.senderStyle.Render(m.username+": ")+msg)
			// m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			// m.viewport.GotoBottom()

			go func() {
				m.sendCh <- ws.Message{
					UserName: m.username,
					Content:  msg,
				}
			}()
		}
	case ws.Message:
		render := m.otherStyle
		if m.username == msg.UserName {
			render = m.senderStyle
		}

		m.messages = append(m.messages, render.Render(msg.UserName+": ")+msg.Content)
		cmds = append(cmds, recvMsgCmd(m.recvCh))
		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
		// default:
		// fmt.Printf("unhandled message: %T\n", msg)
	}

	return m, tea.Batch(cmds...)
}

func (m chatModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
