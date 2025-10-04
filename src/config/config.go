package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port    string
	Name    string
	Version string
	Dir     struct {
		Current string
		Root    string
		Tmp     string
		Data    string
		Logs    string
	}
}

func NewConfig() *Config {
	config := &Config{
		Port:    "8080",
		Name:    "AlertHub",
		Version: "0.1",
	}

	config.Dir.Current = os.Getenv("PWD")
	config.Dir.Root = "/" + config.Name
	config.Dir.Tmp = "/tmp"
	config.Dir.Data = config.Dir.Tmp + "/" + config.Name + "/data"
	config.Dir.Logs = config.Dir.Tmp + "/" + config.Name + "/logs"

	return config
}

func (c *Config) String() string {
	return fmt.Sprintf(`
Port: %s
Name: %s
Version: %s
Dir: {
	Current: %s
	Root: %s
	Tmp: %s
	Data: %s
	Logs: %s
}`, c.Port, c.Name, c.Version, c.Dir.Current, c.Dir.Root, c.Dir.Tmp, c.Dir.Data, c.Dir.Logs)
}
