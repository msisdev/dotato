package initcmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/dotato"
)

func Init(logger *log.Logger, args *args.InitArgs) {
	dotato := dotato.New()
	ok, err := dotato.Init()
	if err != nil {
		logger.Fatal(err)
		return
	}

	if ok {
		fmt.Println("✔ Init done")
	} else {
		fmt.Println("✔ Already initialized")
	}
}
