package gitclone

import (
	"os"
	"testing"
)

func TestGitClone(t *testing.T) {
	defer os.RemoveAll(tempRootPath)
	info := GitClone(tempRootPath+"/image-builder", "https://github.com/backend-simulation/image-builder.git")

	_, err := os.Stat(info.Path)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("not exists: %s", err.Error())
		} else {
			t.Logf("file exists: %s", err.Error())
		}
	}

	t.Log("file exists")
	t.Logf("hash is %s", info.Tag)
}

const tempRootPath = "./tmp"
