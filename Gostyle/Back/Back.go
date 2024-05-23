package Back

var RESET string = "\x1b[49m"
var BLACK string = "\x1b[40m"
var BLUE string = "\x1b[44m"
var CYAN string = "\x1b[46m"
var GREEN string = "\x1b[42m"
var MAGENTA string = "\x1b[45m"
var RED string = "\x1b[41m"
var WHITE string = "\x1b[47m"
var YELLOW string = "\x1b[43m"
var LIGHT_BLACK string = "\x1b[100m"
var LIGHT_BLUE string = "\x1b[104m"
var LIGHT_CYAN string = "\x1b[106m"
var LIGHT_GREEN string = "\x1b[102m"
var LIGHT_MAGENTA string = "\x1b[105m"
var LIGHT_RED string = "\x1b[101m"
var LIGHT_WHITE string = "\x1b[107m"
var LIGHT_YELLOW string = "\x1b[103m"

func Colorize(color string, text string) string {
	return color + text + RESET
}
