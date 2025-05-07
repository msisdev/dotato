package ignore

const (
	IsDir  = true
	IsFile = false
)

const (
	Ignored    = true
	NotIgnored = false
)

type FileEntry struct {
	path      string
	isDir     bool
	isIgnored bool
}

const (
	IgnoreFileName = "dotato.ignore"
)

type IgnoreEntry struct {
	path  string
	lines []string
}
