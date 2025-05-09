package filesystem

import (
	"os"
	"runtime"
)

// What is this function doing?
// 
// dotato depends on "github.com/go-git/go-billy/v5/osfs".
// osfs.New("") occurs errors when it is used with Readlink(absPath).
// So dotato needs to provide actual root path.
func GetRootDir() string {
	if runtime.GOOS == "windows" {
		if drive := os.Getenv("SystemDrive"); drive != "" {
			return drive + "\\"
		}
		panic("Oops! Dotato needs SystemDrive env var. Please set it with your current drive. E.g. 'C:'")
	}
	return "/"
}
