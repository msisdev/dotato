package field

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

func makeInputModel(placeholder string) textinput.Model {
	t := textinput.New()
	t.Cursor.Style = cursorStyle
	t.Placeholder = placeholder
	t.Cursor.SetMode(cursor.CursorStatic)

	return t
}

type FieldModel struct {
	Focus 	int
	Inputs	[]textinput.Model
}

func New(placeholders []string) FieldModel {
	m := FieldModel{
		Focus: 0,
		Inputs: make([]textinput.Model, len(placeholders)),
	}

	for i, placeholder := range placeholders {
		m.Inputs[i] = makeInputModel(placeholder)
	}

	m.Inputs[0].Focus()
	m.Inputs[0].PromptStyle = focusedStyle
	m.Inputs[0].TextStyle = focusedStyle

	return m
}

func (m FieldModel) Init() tea.Cmd {
	return nil
}

func (m FieldModel) Update(msg tea.Msg) (FieldModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			//Did the user press enter while the submit button is focused?
			if s == "enter" && m.Focus == len(m.Inputs) {
				// Submit the form
				return m, tea.Quit
			}

			// Cyle indexes
			if s == "up" || s == "shift+tab" {
				m.Focus--
			} else {
				m.Focus++
			}

			if m.Focus > len(m.Inputs) {	// there is a submit button
				m.Focus = 0
			} else if m.Focus < 0 {
				m.Focus = len(m.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.Inputs))
			for i := 0; i <= len(m.Inputs)-1; i++ {
				if i == m.Focus {
					// Set focus decorations
					cmds[i] = m.Inputs[i].Focus()
					m.Inputs[i].PromptStyle = focusedStyle
					m.Inputs[i].TextStyle = focusedStyle
					continue
				}

				// Remove focus decorations
				m.Inputs[i].Blur()
				m.Inputs[i].PromptStyle = noStyle
				m.Inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m *FieldModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m FieldModel) View() string {
	var b strings.Builder

	for i := range m.Inputs {
		b.WriteString(m.Inputs[i].View())
		if i < len(m.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.Focus == len(m.Inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func (m FieldModel) GetValues() []string {
	values := make([]string, len(m.Inputs))
	for i, input := range m.Inputs {
		values[i] = strings.TrimSpace(input.Value())
	}
	return values
}
