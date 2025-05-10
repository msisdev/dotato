package wherecmd

import (
	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/cli/shared"
)

func Where(logger *log.Logger, args *args.WhereArgs) {
	_, err := shared.New(logger)
	if err != nil {
		logger.Fatal(err)
		return
	}

	panic("unimplemented")
}
