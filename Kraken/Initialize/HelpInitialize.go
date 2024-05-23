package Initialize

import "framework/kraken/Help"

func SetHelpCommands() {
	Help.AddCommand(true, "shell", "Spawn an shell to an session.", "Usage -> shell <session>")

	Help.AddCommand(true, "set", "Set an Private variable.", "Usage -> set <name> <value>")
	Help.AddCommand(false, "sets", "Display Private variables.", "Usage -> sets")
	Help.AddCommand(true, "unset", "Unset a Private variables.", "Usage -> unset <name>")

	Help.AddCommand(true, "export", "Set an Global variable.", "Usage -> export <name> <value>")
	Help.AddCommand(false, "exports", "Display Global variables.", "Usage -> exports")
	Help.AddCommand(true, "unexport", "Unset an Global variable.", "Usage -> unexport <name>")

	Help.AddCommand(true, "use", "Select an exploit.", "Usage -> use <exploit>")
	Help.AddCommand(false, "unuse", "Unselect exploit.", "Usage -> unuse")

	Help.AddCommand(false, "exploit", "Run an exploit against a session.", "Usage -> exploit")
	Help.AddCommand(false, "exploits", "Display all exploits.", "Usage -> exploits")
	Help.AddCommand(true, "searchploit", "Search for exploit containing string.", "Usage -> searchploit <cont>")

	Help.AddCommand(true, "generate", "Generate a payload.", "Usage -> generate os=<os> l=<lhost>:<lport>")
	Help.AddCommand(false, "sessions", "Display all sessions.", "Usage -> sessions")
	Help.AddCommand(false, "reset", "Reset sessions.", "Usage -> reset")
	Help.AddCommand(true, "alias", "Alias Session-Id.", "Usage -> alias <session> <alias>")
	Help.AddCommand(false, "server", "Display server info.", "Usage -> server")
	Help.AddCommand(false, "clear", "clear the terminal.", "Usage -> clear")
	Help.AddCommand(false, "exit", "exit.", "Usage -> exit")
}
