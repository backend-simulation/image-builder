package gitclone

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

type GitCloneInfo struct {
	Path string
	Tag  string
}

func GitClone(clonePath string, gitRepositoryUrl string) GitCloneInfo {
	r, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      gitRepositoryUrl,
		Progress: os.Stdout,
	})

	if err != nil {
		panic(err)
	}

	ref, err := r.Head()
	if err != nil {
		panic(err)
	}

	return GitCloneInfo{
		Path: clonePath,
		Tag:  ref.Hash().String()[0:10],
	}
}
