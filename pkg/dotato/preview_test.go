package dotato

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func TestPreviewImportFile(t *testing.T) {
	var (
		fs = memfs.New()
		d = Dotato{fs: fs, isMem: true}
		dot = gp.GardenPath{"", "home", "user", ".bashrc"}
		dtt = gp.GardenPath{"", "home", "user", "Documents", "dotato", "bash", ".bashrc"}
		content = []byte("test content")
	)
	
	// Dot: file / Dtt: not exists
	{
		// Create file
		err := fs.MkdirAll(dot.Parent().Abs(), 0755)
		assert.NoError(t, err)
		f, err := fs.Create(dot.Abs())
		assert.NoError(t, err)

		_, err = f.Write(content)
		assert.NoError(t, err)
		assert.NoError(t, f.Close())

		// Preview
		p, err := d.PreviewImportFile(dot, dtt)
		assert.NoError(t, err)
		assert.Equal(t, dot, p.Dot)
		assert.Nil(t, p.DotReal)
		assert.Equal(t, dtt, p.Dtt)
		assert.False(t, p.DttExists)
		assert.False(t, p.Equal)
	}
}