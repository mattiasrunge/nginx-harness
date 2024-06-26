package config

import "path/filepath"

var workDir string

func SetWorkDir(dir string) {
	absdir, _ := filepath.Abs(dir)
	workDir = absdir
}

func GetWorkDir() string {
	return workDir
}
