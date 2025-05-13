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
	title := fmt.Sprintf("🥔 Preview │ update %d │ total %d", count, len(ps))
	println(renderBlock(normalBorder, title))

	if len(ps) > 0 {
		println(renderItem(ps[0], arrow))
	}
	if len(ps) > 1 {
		for _, p := range ps[1:] {
			println()
			println(renderItem(p, arrow))
		}
	}
	// for _, p := range ps {
	// 	print(renderItem(p, arrow))
	// }
	
	println(renderBlock(dotBorder, footer))

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
