package confirm

import "github.com/msisdev/dotato/internal/component/textinput"

const (
	yes = "yes"
)

func Run(title string) (bool, error) {
	input, err := textinput.Run(title, yes)
	if err != nil {
		return false, err
	}

	return input == yes, nil
}
