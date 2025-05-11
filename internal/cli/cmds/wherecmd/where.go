package wherecmd

import (
	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/factory"
)

func WhereState(logger *log.Logger, args *args.WhereStateArgs) {
	println(factory.DotatoFilePathState)
}
