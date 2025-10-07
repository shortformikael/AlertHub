package config

import (
	"fmt"
	"os"
)

var (
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
)

func Init() error {
	Port = "8080"
	Name = "AlertHub"
	Version = "0.1"
	Dir.Current = os.Getenv("PWD")
	Dir.Root = "/" + Name
	Dir.Tmp = "/tmp"
	Dir.Data = Dir.Tmp + "/" + Name + "/data"
	Dir.Logs = Dir.Tmp + "/" + Name + "/logs"

	return makeAllDirectories()
}

func makeAllDirectories() error {
	directories := []string{Dir.Logs, Dir.Data, Dir.Tmp}
	for _, directory := range directories {
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			if err := os.MkdirAll(directory, 0o755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", directory, err)
			}
		}
	}
	return nil
}

func String() string {
	return fmt.Sprintf(`
Port: %s
Name: %s
Version: %s
Dir: {
	Current: %s
	Root: %s
	Tmp:  %s
	Data: %s
	Logs: %s
}`,
		Port,
		Name,
		Version,
		Dir.Current,
		Dir.Root,
		Dir.Tmp,
		Dir.Data,
		Dir.Logs)
}
