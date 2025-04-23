package arg

import (
	"github.com/alexflint/go-arg"
)

func Parse() (Args, error) {
	var args Args
	err := arg.Parse(&args)
	return args, err
}
