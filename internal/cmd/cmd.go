package cmd

import (
	"github.com/msisdev/dotato/internal/arg"
	"github.com/msisdev/dotato/pkg/config"
)

func Run() {
	args, err := arg.Parse()
	if err != nil {
		panic(err)
	}

	if args.Version {
		println(config.GetDotatoVersion())
		return
	}
}
