package inputconfirm

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	charLimit = 3
)

type model struct {
	textInput textinput.Model
	title			string
	err 			error
}

func initialModel(title string) model {
	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = "yes"
	ti.CharLimit = charLimit
	ti.Width = charLimit
	
	return model{
		textInput: ti,
		title: title,
		err: nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n",
		m.title,
		m.textInput.View(),
	)
}

func Run(title string) (bool, error) {
	m := initialModel(title)

	p := tea.NewProgram(m)

	if final, err := p.Run(); err != nil {
		return false, err
	} else {
		m = final.(model)
	}

	var ok bool
	switch m.textInput.Value() {
	case "yes":
		ok = true
	default:
		ok = false
	}

	return ok, m.err
}
