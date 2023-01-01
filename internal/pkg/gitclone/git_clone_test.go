package gitclone

import (
	"os"
	"testing"
)

func TestGitClone(t *testing.T) {
	defer os.RemoveAll(tempRootPath)
	info := GitClone(tempRootPath+"/image-builder", "https://github.com/backend-simulation/sample-codes.git")

	_, err := os.Stat(info.Path)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("not exists: %s", err.Error())
		} else {
			t.Logf("file already exists: %s", err.Error())
		}
	}

	t.Log("file git clone successful")
	t.Logf("hash is %s", info.Tag)
}

const tempRootPath = "./tmp"
