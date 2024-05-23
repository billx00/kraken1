package Help

import (
	"fmt"
	"framework/kraken/Gostyle/Log"
)

type _tmp_help struct {
	Name          string
	Functionality string
	Arguments     bool
	Doc           string
}

var _HELP_COMMANDS []_tmp_help = []_tmp_help{
	{Name: "help", Functionality: "display this message.", Doc: "", Arguments: true},
}

func DisplayAll() {
	var longest_name int = 7
	var longest_func int = 13
	for i := 0; i < len(_HELP_COMMANDS); i++ {
		if longest_name < len(_HELP_COMMANDS[i].Name) {
			longest_name = len(_HELP_COMMANDS[i].Name)
		}

		if longest_func < len(_HELP_COMMANDS[i].Functionality) {
			longest_func = len(_HELP_COMMANDS[i].Functionality)
		}
	}

	fmt.Println("Command      Functionality")
	fmt.Println("-------      " + makeStr("-", longest_func))

	for i := 0; i < len(_HELP_COMMANDS); i++ {
		fmt.Printf("%s", _HELP_COMMANDS[i].Name+makeStr(" ", longest_name-len(_HELP_COMMANDS[i].Name)))

		if _HELP_COMMANDS[i].Arguments {
			fmt.Printf("\x1b[32m%s\x1b[39m", " [+] ")
		} else {
			fmt.Printf("%s", "     ")
		}

		fmt.Printf("%s", _HELP_COMMANDS[i].Functionality+"\n")
	}

	fmt.Printf("\n%s\n", "To get more information and usage about a command, use \x1b[33m<help command_name>\x1b[39m")
}

func Help(str string) {
	for i := 0; i < len(_HELP_COMMANDS); i++ {
		if _HELP_COMMANDS[i].Name == str {
			fmt.Println(_HELP_COMMANDS[i].Doc)

			return
		}
	}

	Log.Error("Help: Command not found.")
}

func makeStr(Char string, Length int) string {
	var Str string

	for i := 0; i < Length; i++ {
		Str += Char
	}

	return Str
}

func AddCommand(_arguments bool, _name string, _func string, _doc string) {
	_HELP_COMMANDS = append(_HELP_COMMANDS, _tmp_help{Arguments: _arguments, Name: _name, Functionality: _func, Doc: _doc})
}

func GetHelp(command string) _tmp_help {
	for i := 0; i < len(_HELP_COMMANDS); i++ {
		if _HELP_COMMANDS[i].Name == command {
			return _HELP_COMMANDS[i]
		}
	}

	return _tmp_help{}
}
