package shared

import (
	"fmt"

	"github.com/msisdev/dotato/internal/dotato"
)

const (
	symbolNone      = "✔"
	symbolSkip			= "✘"
	symbolCreate    = "+"
	symbolOverwrite = "!"
	symbolUnknown   = "?"
)

func getSymbol(op dotato.FileOp) string {
	switch op {
	case dotato.FileOpNone:
		return symbolNone
	case dotato.FileOpSkip:
		return symbolSkip
	case dotato.FileOpCreate:
		return symbolCreate
	case dotato.FileOpOverwrite:
		return symbolOverwrite
	default:
		return symbolUnknown
	}
}

func SprintPreviewImportFile(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s\n -> %s",
		getSymbol(p.DttOp),
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
}
func PrintPreviewImportFile(ps []dotato.Preview) {
	printPreview(ps, SprintPreviewImportFile)
}

func SprintPreviewImportLink(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s\n -> %s %s",
		getSymbol(p.DotOp),
		p.Dot.Path.Abs(),
		getSymbol(p.DttOp),
		p.Dtt.Path.Abs(),
	)
}
func PrintPreviewImportLink(ps []dotato.Preview) {
	printPreview(ps, SprintPreviewImportLink)
}

func SprintPreviewExportFile(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s\n <- %s",
		getSymbol(p.DotOp),
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
}
func PrintPreviewExportFile(ps []dotato.Preview) {
	printPreview(ps, SprintPreviewExportFile)
}

func SprintPreviewExportLink(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s\n <- %s",
		getSymbol(p.DotOp),
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
}
func PrintPreviewExportLink(ps []dotato.Preview) {
	printPreview(ps, SprintPreviewExportLink)
}

func SprintPreviewUnlink(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s\n <-> %s",
		getSymbol(p.DotOp),
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
}
func PrintPreviewUnlink(ps []dotato.Preview) {
	printPreview(ps, SprintPreviewUnlink)
}

func printPreview(ps []dotato.Preview, f func(dotato.Preview) string) {
	fmt.Print("\n🔎 Preview\n\n")
	for _, p := range ps {
		fmt.Println(f(p))
		fmt.Println()
	}
	fmt.Println()
}
