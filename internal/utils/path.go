package utils

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Path struct {
	Root string
	Sub  string
	Hash string
}

func GetPath() *Path {
	path := app.Window().URL().Path
	hash := app.Window().URL().Fragment
	root := ""
	sub := ""

	var paths []string = strings.Split(path, "/")
	if len(paths) > 1 {
		root = paths[1]
	}
	if len(paths) > 2 {
		sub = paths[2]
	}

	return &Path{
		Root: root,
		Sub:  sub,
		Hash: hash,
	}
}
