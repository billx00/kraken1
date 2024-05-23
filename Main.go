package main

import (
	"framework/kraken/Kraken/Initialize"
	"framework/kraken/Kraken/Server"
	"framework/kraken/Kraken/Shell"
)

func main() {
	Initialize.Initialize()
	Server.StartServerAsync()
	Shell.SpawnShell()
}
