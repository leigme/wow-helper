package cmd

import (
	_ "embed"
	"encoding/json"
)

//go:embed defaultConfig.json
var defaultConfig []byte

const configName = ".config"

type Config struct {
	Root string `json:"root"`
}

func NewConfigBytes(rootPath ...string) ([]byte, error) {
	rp := ""
	for _, root := range rootPath {
		rp = rp + root
	}
	config := &Config{
		Root: rp,
	}
	return json.Marshal(config)
}
