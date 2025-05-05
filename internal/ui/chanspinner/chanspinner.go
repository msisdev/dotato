package chanspinner

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Go's type system requires this format
type endMsg struct {
	error
}

// Interactive work procedure.
type Task func(up chan<- string, quit <-chan bool) error

type State struct {
	Text string	// text to show
}

type Spinner struct {
	spinner		spinner.Model			// spinner spins
	frame 		string						// buffer for view
	text 			string						// buffer for view

	task 			Task							// task is interactive
	up 				chan string				// used by task
	quit 			chan bool					// quit signal from model to work
	quitting	bool							// to prevent closing channel twice
	
	Error 		error							// error returned from task
}

func New(init string, f Task) Spinner {
	sp := spinner.New()
	sp.Spinner.FPS = time.Millisecond * 100

	return Spinner{
		spinner:	sp,
		frame: 		sp.Spinner.Frames[0],
		text: 		init,
		task: 		f,
		up: 			make(chan string, 1),	// this extra space is needed
		quit: 		make(chan bool),
		quitting: false,
		Error: 		nil,
	}
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
	case text, ok := <-m.up:
		if ok {
			return State{text}
		} else {
			return nil
		}
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
		m.text = msg.Text	// Update state
		return m, m.wait	// Run another wait

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		m.frame = m.spinner.View()
		return m, cmd			// Restart tick

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if !m.quitting {
				m.quitting = true
				close(m.quit)	// Send quit signal
				m.spinner.Spinner = spinner.Pulse	
			}
			
			return m, nil		// Quit is not happening yet
		}

	case endMsg:
		// Handle error and frame
		m.Error = msg.error
		if m.Error != nil {
			m.frame = "✖"
		} else {
			m.frame = "✔"
		}

		// Drain up channel
		empty := false
		for !empty {
			select {
			case text := <-m.up:
				m.text = text
			default:
				empty = true
			}
		}

		close(m.up)
		
		return m, tea.Quit	// Quit
	}

	return m, nil
}

func (m Spinner) View() string {
	return fmt.Sprintf("%s %s\n", m.frame, m.text)
}

func Run(init string, f Task) error {
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
