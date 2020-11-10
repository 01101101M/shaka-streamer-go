package internal

import (
	"io/ioutil"
	"os"
	"syscall"
)

//Create a uniquely-named named pipe in the temp directory.
func CreatePipe(suffix string) (string, error) {
	tempDir, _ := ioutil.TempDir("", "shaka-live-")
	tempFile, err := ioutil.TempFile(tempDir, suffix)
	os.Remove(tempFile.Name())
	err = syscall.Mkfifo(tempFile.Name(), 0600)
	return tempFile.Name(), err
}
