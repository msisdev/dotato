package shared

import (
	"fmt"

	"github.com/msisdev/dotato/pkg/dotato"
)

const (
	symbolNone = "✔"
	symbolCreate = "+"
	symbolOverwrite = "!"
	symbolUnknown = "?"
)

func getSymbol(op dotato.FileOp) string {
	switch op {
	case dotato.FileOpNone:
		return symbolNone
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
		"%s %s -> %s",
		getSymbol(p.DttOp),
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
}

func SprintPreviewImportLink(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s -> %s %s",
		getSymbol(p.DotOp),
		p.Dot.Path.Abs(),
		getSymbol(p.DttOp),
		p.Dtt.Path.Abs(),
	)
}

func SprintPreviewExportFile(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s <- %s",
		getSymbol(p.DotOp),
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
}

func SprintPreviewExportLink(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s <- %s",
		getSymbol(p.DotOp),
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
}

func SprintPreviewUnlink(p dotato.Preview) string {
	return fmt.Sprintf(
		"%s %s <-> %s",
		getSymbol(p.DotOp),
		p.Dot.Path.Abs(),
		p.Dtt.Path.Abs(),
	)
}
