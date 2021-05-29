package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func safeCreateDir(dir string) {
	if err := os.Mkdir(dir, fs.ModePerm); err == nil || os.IsExist(err) {
		return
	} else {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createDefaultFiles(gitDir string) {
	safeCreateDir(filepath.Join(gitDir, "refs"))
	safeCreateDir(filepath.Join(gitDir, "refs/heads"))
	safeCreateDir(filepath.Join(gitDir, "refs/tags"))
	if err := os.Symlink("refs/heads/master", filepath.Join(gitDir, "HEAD")); os.IsExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	gitDir := ".sit"
	safeCreateDir(gitDir)
	createDefaultFiles(gitDir)
	sha1Dir := filepath.Join(gitDir, "objects")
	safeCreateDir(sha1Dir)
	for i := 0; i < 256; i++ {
		safeCreateDir(filepath.Join(sha1Dir, fmt.Sprintf("%02x", i)))
	}
	safeCreateDir(filepath.Join(gitDir, "/pack"))
}
