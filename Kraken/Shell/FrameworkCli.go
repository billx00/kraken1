package Shell

import (
	"fmt"
	"framework/kraken/Gostyle/Fore"
	"framework/kraken/Gostyle/Log"
	"framework/kraken/Kraken/Public"
	"framework/kraken/Kraken/Session"
	"framework/kraken/Utilities"
	"os"
	"strconv"
)

func GetPrompt(isRemote bool, ExtraInfo string) string {
	var ActiveSessions int

	for i := 0; i < len(Session.Sessions); i++ {
		if Session.IsSessionActive(Session.Sessions[i]) {
			ActiveSessions++
		}
	}

	var FullEclams string
	var RemoteStr string = Fore.Colorize(Fore.BLUE, "(local)")
	var ExploitStr string
	var BotsStr string

	if isRemote {
		RemoteStr = Fore.Colorize(Fore.RED, "(remote)")
	}

	if len(Public.SelectedExploit) != 0 && !isRemote {
		ExploitStr = fmt.Sprintf("exploit:%s ", Fore.Colorize(Fore.RED, Public.SelectedExploit))
	}

	if ActiveSessions != 0 {
		BotsStr = fmt.Sprintf("bots:%s", Fore.Colorize(Fore.GREEN, strconv.Itoa(ActiveSessions)))
	}

	if len(ExploitStr) != 0 || len(BotsStr) != 0 {
		FullEclams = fmt.Sprintf("[%s%s] ", ExploitStr, BotsStr)
	}

	return fmt.Sprintf("%s kraken %s%s>> ", RemoteStr, FullEclams, ExtraInfo)
}

func SpawnShell() {
	for {
		fmt.Printf("\n%s", GetPrompt(false, Utilities.StringEmpty))

		var UserInput string = Utilities.GetInput()

		if len(UserInput) <= 0 {
			continue
		}

		var ParsedUserInput []string = Utilities.CmdParse(UserInput)

		var CommandName string = ParsedUserInput[0]

		switch CommandName {
		case ("sessions"):
			Session.SesssionsCommand()
		case ("shell"):
			ShellCommand(ParsedUserInput)
		case ("set"):
			SetCommand(ParsedUserInput)
		case ("sets"):
			SetsCommand()
		case ("unset"):
			UnsetCommand(ParsedUserInput)
		case ("export"):
			ExportCommand(ParsedUserInput)
		case ("exports"):
			ExportsCommand()
		case ("unexport"):
			UnexportCommand(ParsedUserInput)
		case ("use"):
			UseCommand(ParsedUserInput)
		case ("unuse"):
			UnuseCommand()
		case ("options"):
			OptionsCommand()
		case ("exploit"):
			ExploitCommand()
		case ("exploits"):
			ExploitsCommand()
		case ("searchploit"):
			SearchploitCommand(ParsedUserInput)
		case ("generate"):
			GenerateCommand(ParsedUserInput)
		case ("alias"):
			AliasCommand(ParsedUserInput)
		case ("reset"):
			ResetCommand()
		case ("server"):
			ServerCommand()
		case ("clear"):
			ClearCommand()
		case ("help"):
			HelpCommand(ParsedUserInput)
		case ("exit"):
			os.Exit(0)
		default:
			Log.Error("Command not found.")
		}
	}
}
