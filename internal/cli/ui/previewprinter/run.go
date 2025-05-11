package previewprinter

import (
	"fmt"

	"github.com/msisdev/dotato/internal/cli/app"
)

func run(ps []app.Preview, arrow string) {
	fmt.Printf("\n🔎 Preview: total %d\n\n", len(ps))
	for _, p := range ps {
		println(render(p, arrow))
	}
	println()
}

func RunPreviewImportFile(ps []app.Preview) {
	run(ps, arrowImportFile)
}

func RunPreviewImportLink(ps []app.Preview) {
	run(ps, arrowImportLink)
}

func RunPreviewExportFile(ps []app.Preview) {
	run(ps, arrowExportFile)
}

func RunPreviewExportLink(ps []app.Preview) {
	run(ps, arrowExportLink)
}

func RunPreviewUnlink(ps []app.Preview) {
	run(ps, arrowUnlink)
}
