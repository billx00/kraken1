package Public

import "framework/kraken/Utilities"

var ContinueShell bool = true

var SelectedExploit string

var GlobalVariables map[string]string = map[string]string{}
var PrivateVariables map[string]string = map[string]string{}

func GetGlobalOrPirvate(s string) string {
	if PrivateVariables[s] != Utilities.StringEmpty {
		return PrivateVariables[s]
	}

	return GlobalVariables[s]
}
