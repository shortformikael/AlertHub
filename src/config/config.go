package config

import (
	"fmt"
	"os"
)

var Port string
var Name string
var Version string
var Dir struct {
	Current string
	Root    string
	Tmp     string
	Data    string
	Logs    string
}

func Init() {
	Port = "8080"
	Name = "AlertHub"
	Version = "0.1"
	Dir.Current = os.Getenv("PWD")
	Dir.Root = "/" + Name
	Dir.Tmp = "/tmp"
	Dir.Data = Dir.Tmp + "/" + Name + "/data"
	Dir.Logs = Dir.Tmp + "/" + Name + "/logs"

	makeAllDirectories()
}

func makeAllDirectories() {
	var directories = []string{Dir.Logs, Dir.Data, Dir.Tmp}
	for _, directory := range directories {
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			os.MkdirAll(directory, 0755)
		}
	}
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
