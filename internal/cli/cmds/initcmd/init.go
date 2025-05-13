package initcmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/msisdev/dotato/internal/cli/args"
	"github.com/msisdev/dotato/internal/factory"
	"github.com/msisdev/dotato/internal/lib/filesystem"
)

func Init(logger *log.Logger, args *args.InitArgs) {
	ok, err := factory.WriteExampleConfig(filesystem.NewOSFS(), 0666)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	if ok {
		fmt.Println("✔ Init done")
	} else {
		fmt.Println("✔ Already initialized")
	}
}
