package importgroup

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	libstate "github.com/msisdev/dotato/internal/lib/state"
	"github.com/msisdev/dotato/pkg/dotato"
)

type PreviewPage struct {
	dtt				*dotato.Dotato
	group 		string
	resolver	string

	// goroutine var
	progvar libstate.State[int]

	// view buffer
	progbuf string
}

func newPreviewPage(dtt *dotato.Dotato, group string, resolver string) PreviewPage {
	p := PreviewPage{
		dtt:    dotato.New(),
		group:  "",
		resolver: "",
	}

	return p
}

func (p PreviewPage) Init() tea.Cmd {
	return tea.Tick(time.Duration(1), func(t time.Time) tea.Msg {
		
	})
}

func (p PreviewPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}
