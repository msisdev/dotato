package previewprinter

import (
	"github.com/msisdev/dotato/internal/cli/app"
)

func run(ps []app.Preview, viewAll bool, arrow string) int {
	// count
	updates := 0
	for _, p := range ps {
		if isUpdate(p) {
			updates++
		}
	}

	println(renderHeader(updates, len(ps)))
	println()
	printPreviews(ps, viewAll, arrow)
	println()
	println(renderBlock())

	return updates
}

func isUpdate(p app.Preview) bool {
	// If both dot and dtt are none, it is up to date
	if (p.DotOp == app.FileOpNone || p.DotOp == app.FileOpSkip) &&
		(p.DttOp == app.FileOpNone || p.DttOp == app.FileOpSkip) {
		return false
	}

	return true
}

func isPrintSkipped(p app.Preview) bool {
	if p.DotOp == app.FileOpNone && p.DttOp == app.FileOpNone {
		return true
	}

	return false
}

func printPreviews(ps []app.Preview, viewAll bool, arrow string) {
	idx := 0
	printed := false

	// Try to print at least one item
	for ; idx < len(ps); idx++ {
		if !viewAll && isPrintSkipped(ps[idx]) {
			continue
		}

		println(renderItem(ps[idx], arrow))
		idx++
		printed = true
		break
	}

	// Print the rest of the items
	for ; idx < len(ps); idx++ {
		if !viewAll && isPrintSkipped(ps[idx]) {
			continue
		}

		println()
		println(renderItem(ps[idx], arrow))
		printed = true
	}

	if !printed {
		println("☀️ All files are ok.")
	}
}

func RunPreviewImportFile(ps []app.Preview, viewAll bool) int {
	return run(ps, viewAll, arrowImportFile)
}

func RunPreviewImportLink(ps []app.Preview, viewAll bool) int {
	return run(ps, viewAll, arrowImportLink)
}

func RunPreviewExportFile(ps []app.Preview, viewAll bool) int {
	return run(ps, viewAll, arrowExportFile)
}

func RunPreviewExportLink(ps []app.Preview, viewAll bool) int {
	return run(ps, viewAll, arrowExportLink)
}

func RunPreviewUnlink(ps []app.Preview, viewAll bool) int {
	return run(ps, viewAll, arrowUnlink)
}
