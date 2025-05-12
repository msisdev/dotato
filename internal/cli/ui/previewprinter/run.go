package previewprinter

import (
	"fmt"

	"github.com/msisdev/dotato/internal/cli/app"
)

func countUpdate(ps []app.Preview) int {
	count := 0
	for _, p := range ps {
		if (p.DotOp != app.FileOpNone && p.DotOp != app.FileOpSkip) ||
			(p.DttOp != app.FileOpNone && p.DttOp != app.FileOpSkip) {
			count++
		}
	}
	return count
}

func run(ps []app.Preview, arrow string) int {
	count := countUpdate(ps)
	fmt.Printf("\n🔎 Preview: update %d / total %d\n\n", count, len(ps))
	for _, p := range ps {
		println(renderItem(p, arrow))
	}
	println(footer)

	return count
}

func RunPreviewImportFile(ps []app.Preview) int {
	return run(ps, arrowImportFile)
}

func RunPreviewImportLink(ps []app.Preview) int {
	return run(ps, arrowImportLink)
}

func RunPreviewExportFile(ps []app.Preview) int {
	return run(ps, arrowExportFile)
}

func RunPreviewExportLink(ps []app.Preview) int {
	return run(ps, arrowExportLink)
}

func RunPreviewUnlink(ps []app.Preview) int {
	return run(ps, arrowUnlink)
}
