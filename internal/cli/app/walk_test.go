package app

// import (
// 	"os"
// 	"testing"

// 	"github.com/charmbracelet/log"
// 	"github.com/msisdev/dotato/internal/lib/filesystem"
// 	"github.com/stretchr/testify/assert"
// )

// func TestWalkImportFile(t *testing.T) {
// 	logger := log.New(os.Stdout)
// 	fs := filesystem.NewOSFS()
// 	app := NewWithFS(logger, fs, false)

// 	base, notFound, err := app.E.GetConfigGroupBase("bash", "nux")
// 	assert.NoError(t, err)
// 	assert.Nil(t, notFound)
// 	assert.NotNil(t, base)

// 	err = app.WalkImportFile("bash", base, func(p Preview) error {

// 		return nil
// 	})
// 	assert.NoError(t, err)
// }