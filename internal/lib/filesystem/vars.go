package filesystem

import "fmt"

const (
	maxSymlinkDepth = 10 // prevent infinite loop
)

var (
	ErrTooManySymlinks = fmt.Errorf("too many levels of symbolic links")
)
