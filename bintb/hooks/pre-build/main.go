package main

import (
	"github.com/moisespsena-go/hooks"
	pre_build "github.com/moisespsena-go/hooks/pre-build"
)

func main() {
	hooks.NewRunner(pre_build.Hooks()).
		Run().
		Exit()
}
