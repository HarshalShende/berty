// +build !android,!darwin

package tempdir

import "os"

func tempDir() string {
	return os.TempDir()
}
