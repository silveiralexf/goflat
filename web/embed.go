package web

import (
	"embed"
	"io/fs"
	"path/filepath"
)

//go:embed dist/*
var assets embed.FS

func LoadEmbedFS(pathPrefix string) (fs.FS, error) {
	return fs.Sub(assets, filepath.Clean(pathPrefix))
}
