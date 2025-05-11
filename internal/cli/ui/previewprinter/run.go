package previewprinter

import "github.com/msisdev/dotato/internal/cli/app"

func run(ps []app.Preview, f func(app.Preview) string) {
	println("\n🔎 Preview\n\n")
	for _, p := range ps {
		println(f(p))
	}
	println()
}

func RunPreviewImportFile(ps []app.Preview) {
	run(ps, sprintPreviewImportFile)
}

func RunPreviewImportLink(ps []app.Preview) {
	run(ps, sprintPreviewImportLink)
}

func RunPreviewExportFile(ps []app.Preview) {
	run(ps, sprintPreviewExportFile)
}

func RunPreviewExportLink(ps []app.Preview) {
	run(ps, sprintPreviewExportLink)
}

func RunPreviewUnlink(ps []app.Preview) {
	run(ps, sprintPreviewUnlink)
}
