package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/moisespsena-go/hooks"
	post_build "github.com/moisespsena-go/hooks/post-build"
)

func main() {
	binName := os.Args[1]

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	var destNames []string
	err := filepath.Walk("dist", func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if strings.HasSuffix(path, string(filepath.Separator)+binName) {
				destNames = append(destNames, path)
			}
		}
		return err
	})

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	hooks.NewRunner(post_build.Hooks().Jobs(destNames...)).
		Run().
		Exit()
}
