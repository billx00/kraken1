package Initialize

import (
	"fmt"
	"framework/kraken/Gostyle/Log"
	"framework/kraken/Kraken/Arguments"
	"framework/kraken/Kraken/ExploitDaB"
	"framework/kraken/Kraken/Session"
	"framework/kraken/Kraken/Shell"
	"os"
	"os/signal"
)

func PrintInfo() {
	fmt.Printf("%s\n", Shell.InfoAndArt)
	Log.Info("HTTPEngine running on " + Arguments.ServerInfo["EngineServer"] + "\n\n")
	Log.Warning("Bull does not offer fully interactive shells.\n")
}

func StartAsyncCTRLCListener() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			_ = sig
			Session.SessionStopWait()
		}
	}()
}

func Initialize() {
	ExploitDaB.LoadExploits()
	StartAsyncCTRLCListener()
	SetHelpCommands()
	Arguments.ParseAndSet()
	PrintInfo()
}
