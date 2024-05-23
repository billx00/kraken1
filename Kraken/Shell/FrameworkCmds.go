package Shell

import (
	"fmt"
	"framework/kraken/Gostyle/Fore"
	"framework/kraken/Gostyle/Log"
	"framework/kraken/Help"
	"framework/kraken/Kraken/ExploitDaB"
	"framework/kraken/Kraken/PayloadGenerator"
	"framework/kraken/Kraken/Public"
	"framework/kraken/Kraken/Session"
	"framework/kraken/Utilities"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func OptionsCommand() {
	if Public.SelectedExploit == Utilities.StringEmpty {
		fmt.Printf("%s", "No exploit selected.")
		return
	}

	for i := 0; i < len(ExploitDaB.ExploitDB); i++ {
		if ExploitDaB.ExploitDB[i].Pname == Public.SelectedExploit {
			ExploitDaB.DisplayOptions(ExploitDaB.ExploitDB[i].Options)
			return
		}
	}

	Log.Error("Exploit not found.")
}

func UnexportCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) != 2 {
		fmt.Printf("%s", Help.GetHelp("unexport").Doc)
		return
	}

	delete(Public.GlobalVariables, ParsedUserInput[1])
}

func ExportCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) != 3 {
		fmt.Printf("%s", Help.GetHelp("export").Doc)
		return
	}

	Public.GlobalVariables[ParsedUserInput[1]] = ParsedUserInput[2]
	fmt.Printf("%s => %s", ParsedUserInput[1], ParsedUserInput[2])
}

func ExportsCommand() {
	for k, v := range Public.GlobalVariables {
		fmt.Printf("%s %s\n", k, v)
	}
}

func UnsetCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) != 2 {
		fmt.Printf("%s", Help.GetHelp("unset").Doc)
		return
	}

	delete(Public.PrivateVariables, ParsedUserInput[1])
}

func SetCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) != 3 {
		fmt.Printf("%s", Help.GetHelp("set").Doc)
		return
	}

	Public.PrivateVariables[ParsedUserInput[1]] = ParsedUserInput[2]
	fmt.Printf("%s => %s", ParsedUserInput[1], ParsedUserInput[2])
}

func SetsCommand() {
	for k, v := range Public.PrivateVariables {
		fmt.Printf("%s %s\n", k, v)
	}
}

func UnuseCommand() {
	Public.PrivateVariables = map[string]string{}

	Public.SelectedExploit = Utilities.StringEmpty
}

func UseCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) != 2 {
		fmt.Printf("%s", Help.GetHelp("exploit").Doc)
		return
	}

	Public.PrivateVariables = map[string]string{}

	fmt.Printf("%s", fmt.Sprintf("Exploit => %s", ParsedUserInput[1]))
	Public.SelectedExploit = ParsedUserInput[1]
}

func ExploitCommand() {
	if Public.SelectedExploit != Utilities.StringEmpty {
		ExploitDaB.RunExploit(Public.SelectedExploit)
	} else {
		fmt.Printf("%s", "No exploit selected.")
	}
}

func SearchploitCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) != 2 {
		fmt.Printf("%s", Help.GetHelp("searchploit").Doc)
		return
	}

	var exploits []string
	for i := 0; i < len(ExploitDaB.ExploitDB); i++ {
		if strings.Contains(ExploitDaB.ExploitDB[i].Pname, ParsedUserInput[1]) {
			exploits = append(exploits, ExploitDaB.ExploitDB[i].Pname)
		}
	}

	sort.Strings(exploits)

	for i := 0; i < len(exploits); i++ {
		fmt.Printf("%s\n", exploits[i])
	}
}

func ExploitsCommand() {
	var Exploits []string
	for i := 0; i < len(ExploitDaB.ExploitDB); i++ {
		Exploits = append(Exploits, ExploitDaB.ExploitDB[i].Pname)
	}

	sort.Strings(Exploits)

	for i := 0; i < len(Exploits); i++ {
		fmt.Printf("%s\n", Exploits[i])
	}
}

func ShellCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) != 2 {
		fmt.Printf("%s", Help.GetHelp("shell").Doc)
		return
	}

	var SessionID string = ParsedUserInput[1]
	var _Session Session.Session

	if Session.GetSessionIndex(SessionID) != -1 {
		if !Session.IsSessionActive(Session.Sessions[Session.GetSessionIndex(SessionID)]) {
			fmt.Printf("%s", "Session not active.")
			return
		}
		_Session = Session.Sessions[Session.GetSessionIndex(SessionID)]
	} else {
		fmt.Printf("%s", "Session was not found.")
		return
	}

	for {
		var ExtraInfo string = _Session.Os + "(" + Fore.Colorize(Fore.RED, _Session.Username+"@"+_Session.Hostname) + ") "

		fmt.Printf("%s", GetPrompt(true, ExtraInfo))

		var UserInput string = Utilities.GetInput()

		if UserInput == "exit" {
			break
		} else if len(strings.Trim(UserInput, " ")) == 0 {
			continue
		}

		if Session.GetSessionIndex(SessionID) != -1 {
			Session.Sessions[Session.GetSessionIndex(SessionID)].Job = UserInput
			Session.SessionStartWait(Session.Sessions[Session.GetSessionIndex(SessionID)].SessionId)
		}

		Session.SessionWait()
	}
}

func GenerateCommand(ParsedUserInput []string) {
	var os string
	var l string
	for i := 0; i < len(ParsedUserInput); i++ {
		var arg string = ParsedUserInput[i]
		if strings.Contains(arg, "=") {
			var split []string = strings.Split(arg, "=")
			switch split[0] {
			case ("os"):
				os = split[1]
			case ("l"):
				l = split[1]
			}
		}
	}

	if len(os) < 1 || len(l) < 1 {
		fmt.Printf("%s", Help.GetHelp("generate").Doc)
		return
	}

	var gened string = PayloadGenerator.GeneratePayload(os, l)

	if gened == "OSNx0" {
		Log.Error("Operating System is not supported.")
	} else {
		fmt.Printf("%s", gened)
	}
}

func AliasCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) != 3 {
		fmt.Printf("%s", Help.GetHelp("alias").Doc)
		return
	}

	if Session.GetSessionIndex(ParsedUserInput[1]) != -1 {
		Session.Sessions[Session.GetSessionIndex(ParsedUserInput[1])].Alias = ParsedUserInput[2]
	} else {
		fmt.Printf("Sesssion was not found.")
	}
}

func ResetCommand() {
	Log.Warning("All aliases will be removed, do you want to continue [Y/*] ")

	var uin string = Utilities.GetInput()
	if strings.ToLower(uin) == "y" {
		Session.Sessions = []Session.Session{}
		Log.Info("Sesssion list was reset.\n")
	} else {
		Log.Info("Reset was Successfuly cancled.\n")
	}
}

func ServerCommand() {
	fmt.Printf("%s\n", InfoAndArt)
}

func ClearCommand() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func HelpCommand(ParsedUserInput []string) {
	if len(ParsedUserInput) > 1 {
		Help.Help(ParsedUserInput[1])
	} else {
		Help.DisplayAll()
	}
}
