package spinner

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type endMsg struct {
	error
}

// Interactive work procedure. Must close up channel when done.
type Task func(up chan<- State, quit <-chan bool) error

type State struct {
	Frame		string	// use as spinner frame when task is done
	Text 		string	// text to show
	End			bool		// true if task is done
}

type Spinner struct {
	spinner	spinner.Model			// spinner spins

	task 		Task							// task is interactive
	up 			chan State				// used by work
	quit 		chan bool					// quit signal from model to work

	Error 	error							// error from work
	State		State
}

func New(init State, f Task) Spinner {
	s := Spinner{
		spinner:	spinner.New(),
		task: 		f,
		up: 			make(chan State),
		quit: 		make(chan bool),
		Error: 		nil,
		State: 		init,
	}

	s.spinner.Spinner.FPS = time.Millisecond * 100

	return s
}

func (m Spinner) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.wait,
		m.work,
	)
}

// Wait for any signal
func (m Spinner) wait() tea.Msg {
	select {
	case state := <-m.up:
		return state
	case <-m.quit:
		return nil
	}
}

// Run task and wrap returning error with explicit type
func (m Spinner) work() tea.Msg {
	err := m.task(m.up, m.quit)
	return endMsg{err}
}

func (m Spinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case State:
		m.State = msg			// Update state
		return m, m.wait	// Run another wait

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd			// Restart tick

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			close(m.quit)	// Send quit signal
			m.spinner.Spinner = spinner.Pulse
			return m, nil		// Quit is not happening yet
		}

	case endMsg:
		m.Error = msg.error
		return m, tea.Quit
	}

	return m, nil
}

func (m Spinner) View() string {
	if m.Error != nil {
		return fmt.Sprintf("%s %s\n", m.State.Frame, m.State.Text)
	}

	return fmt.Sprintf("%s %s\n", m.spinner.View(), m.State.Text)
}

func Run(init State, f Task) error {
	m := New(init, f)
	p := tea.NewProgram(m)

	if model, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	} else {
		m = model.(Spinner)
	}

	return m.Error
}
