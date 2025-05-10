package mxspinner

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/msisdev/dotato/internal/lib/store"
)

const (
	spinnerDuration	= time.Millisecond * 100
	watchDuration 	= time.Millisecond * 25
)

// Go's type system requires this format
type endMsg struct {
	error
}

// Interactive work procedure.
type Task func(store *store.Store[string], quit <-chan bool) error

type taskMsg struct {
	text	string	// text to show
	ok 		bool
}

type Spinner struct {
	spinner		spinner.Model					// spinner spins
	frame 		string								// buffer for view
	text 			string								// buffer for view

	task 			Task									// task is interactive
	store			*store.Store[string]	// used by task
	quit 			chan bool							// quit signal from model to work
	quitting	bool									// to prevent closing channel twice
	
	Error 		error									// error returned from task
}

func New(init string, f Task) Spinner {
	sp := spinner.New()
	sp.Spinner.FPS = spinnerDuration

	return Spinner{
		spinner:	sp,
		frame: 		sp.Spinner.Frames[0],
		text: 		init,
		task: 		f,
		store: 		store.New(init, true),
		quit: 		make(chan bool),
		quitting: false,
		Error: 		nil,
	}
}

func (m Spinner) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.watch(),
		m.work,
	)
}

// Wait for any signal
func (m Spinner) watch() tea.Cmd {
	return tea.Tick(watchDuration, func(time.Time) tea.Msg {
		// Check quit signal (non-blocking)
		select {
		case <-m.quit:
			return nil
		default:
		}

		// Check store
		text, ok := m.store.Pop()
		return taskMsg{text, ok}
	})
}

// Run task and wrap returning error with explicit type
func (m Spinner) work() tea.Msg {
	err := m.task(m.store, m.quit)
	return endMsg{err}
}

func (m Spinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case taskMsg:
		if msg.ok {
			m.text = msg.text
		}
		return m, m.watch()	// Run another wait

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

		// Drain store
		if text, ok := m.store.Pop(); ok {
			m.text = text
		}
		
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
