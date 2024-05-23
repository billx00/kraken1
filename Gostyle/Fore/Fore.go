package Fore

var RESET string = "\x1b[39m"
var BLACK string = "\x1b[30m"
var BLUE string = "\x1b[34m"
var CYAN string = "\x1b[36m"
var GREEN string = "\x1b[32m"
var MAGENTA string = "\x1b[35m"
var RED string = "\x1b[31m"
var WHITE string = "\x1b[37m"
var YELLOW string = "\x1b[33m"
var LIGHT_BLACK string = "\x1b[90m"
var LIGHT_BLUE string = "\x1b[94m"
var LIGHT_CYAN string = "\x1b[96m"
var LIGHT_GREEN string = "\x1b[92m"
var LIGHT_MAGENTA string = "\x1b[95m"
var LIGHT_RED string = "\x1b[91m"
var LIGHT_WHITE string = "\x1b[97m"
var LIGHT_YELLOW string = "\x1b[93m"

func Colorize(color string, text string) string {
	return color + text + RESET
}
