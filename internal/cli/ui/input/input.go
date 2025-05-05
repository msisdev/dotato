package input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/msisdev/dotato/internal/cli/ui"
)

const (
	charLimit = 156
	width = 20
)

type (
	errMsg error
)

type Input struct {
	textInput	textinput.Model
	err 			error
}

func New(placeholder string) Input {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.CharLimit = charLimit
	ti.Width = width

	return Input{
		textInput: ti,
	}
}

func (m Input) Init() tea.Cmd {
	return nil
}

func (m Input) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		
		}

	// handle error
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Input) View() string {
	return m.textInput.View()
}

func (m Input) Focus() {
	m.textInput.Focus()
	m.textInput.PromptStyle = ui.FocusedStyle
	m.textInput.TextStyle = ui.FocusedStyle
}

// De-focus
func (m Input) Blur() {
	m.textInput.Blur()
	m.textInput.PromptStyle = ui.BlurredStyle
	m.textInput.TextStyle = ui.BlurredStyle
}
