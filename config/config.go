package config

import (
	"os"

	"github.com/sblinch/kdl-go"
)

type config struct {
	Port      int         `kdl:"port"`
	StaticDir string      `kdl:"static-dir"`
	Mongo     mongoConfig `kdl:"mongo"`
}

type mongoConfig struct {
	URI       string `kdl:"uri"`
	Database  string `kdl:"database"`
	CollUsers string `kdl:"coll-users"`
	CollNotes string `kdl:"coll-notes"`
}

var cfg config

func Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return kdl.Unmarshal(data, &cfg)
}

func Get() *config {
	return &cfg
}
